package main

import (
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/davecgh/go-spew/spew"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
)

const (
	// maxChannelSize is the larget channel that the faucet will create to
	// another peer.
	maxChannelSize int64 = 4000000

	// minChannelSize is the smallest channel that the faucet will extend
	// to a peer.
	minChannelSize int64 = 10000
)

// chanCreationError is an enum which describes the exact nature of an error
// encountered when a user attempts to create a channel with the faucet. This
// enum is used within the templates to determine at which input item the error
// occurred  and also to generate an error string unique to the error.
type chanCreationError uint8

const (
	// NoError is the default error which indicates either the form hasn't
	// yet been submitted or no errors have arisen.
	NoError chanCreationError = iota

	// InvalidAddress indicates that the passed node address is invalid.
	InvalidAddress

	// NotConnected indicates that the target peer isn't connected to the
	// faucet.
	NotConnected

	// ChanAmountNotNumber indicates that the amount specified for the
	// amount to fund the channel with isn't actually a number.
	ChanAmountNotNumber

	// ChannelTooLarge indicates that the amounts specified to fund the
	// channel with is greater than maxChannelSize.
	ChannelTooLarge

	// ChannelTooSmall indicates that the channel size required is below
	// minChannelSize.
	ChannelTooSmall

	// PushIncorrect indicates that the amount specified to push to the
	// other end of the channel is greater-than-or-equal-to the local
	// funding amount.
	PushIncorrect

	// ChannelOpenFail indicates some error occurred when attempting to
	// open a channel with the target peer.
	ChannelOpenFail

	// HaveChannel indicates that the faucet already has a channel open
	// with the target node.
	HaveChannel
)

// String returns a human readable string describing the chanCreationError.
// This string is used in the templates in order to display the error to the
// user.
func (c chanCreationError) String() string {
	switch c {
	case NoError:
		return ""
	case InvalidAddress:
		return "node address is invalid"
	case NotConnected:
		return "not connected to node"
	case ChanAmountNotNumber:
		return "amount must be int"
	case ChannelTooLarge:
		return "channel value is too large"
	case ChannelTooSmall:
		return "channel value is too small (min is 10k sat)"
	case PushIncorrect:
		return "push amount is incorrect"
	case ChannelOpenFail:
		return "unable to open channel"
	case HaveChannel:
		return "node has channel with faucet already"
	default:
		return fmt.Sprintf("%v", uint8(c))
	}
}

// lightningFaucet is a Bitcoin Channel Faucet. The faucet itself is a web app
// that is capable of programmatically opening channels with users with the
// size of the channel parametrized by the user. The faucet required a
// connection to a local lnd node in order to operate properly. The faucet
// implements the constrains on the channel size, and also will only open a
// single channel to a particular node. Finally, the faucet will periodically
// close channels based on their age as the faucet will only open up 100
// channels total at any given time.
type lightningFaucet struct {
	lnd lnrpc.LightningClient

	templates *template.Template

	openChanMtx  sync.RWMutex
	openChannels map[wire.OutPoint]time.Time
}

// newLightningFaucet creates a new channel faucet that's bound to a cluster of
// lnd nodes, and uses the passed templates to render the web page.
func newLightningFaucet(lndHost string,
	templates *template.Template) (*lightningFaucet, error) {

	// First attempt to establish a connection to lnd's RPC sever.
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(*lndNodes, opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to dial to lnd's gRPC server: ",
			err)
	}

	// If we're able to connect out to the lnd node, then we can start up
	// the faucet safely.
	lnd := lnrpc.NewLightningClient(conn)

	return &lightningFaucet{
		lnd:       lnd,
		templates: templates,
	}, nil
}

// homePageContext defines the initial context required for rendering home
// page. The home page displays some basic statistics, errors in the case of an
// invalid channel submission, and finally a splash page upon successful
// creation of a channel.
type homePageContext struct {
	// NumCoins is the number of coins in BTC that the faucet has available
	// for channel creation.
	NumCoins float64

	// NumChannels is the number of active channels that the faucet has
	// open at a given time.
	NumChannels uint32

	// NodeAddr is the full <pubkey>@host:port where the faucet can be
	// connect to.
	NodeAddr string

	// SubmissionError is a enum that stores if any error took place during
	// the creation of a channel.
	SubmissionError chanCreationError

	// ChannelTxid is the txid of the created funding channel. If this
	// field is an empty string, then that indicates the channel hasn't yet
	// been created.
	ChannelTxid string

	// NumConfs is the number of confirmations required for the channel to
	// open up.
	NumConfs uint32
}

// fetchHomeState is helper functions that populates the homePageContext with
// the latest state from the local lnd node.
func (l *lightningFaucet) fetchHomeState() (*homePageContext, error) {
	// First query for the general information from the lnd node, this'll
	// be used to populate the number of active channel as well as the
	// identity of the node.
	infoReq := &lnrpc.GetInfoRequest{}
	nodeInfo, err := l.lnd.GetInfo(ctxb, infoReq)
	if err != nil {
		return nil, err
	}

	// Next obtain the wallet's available balance which indicates how much
	// we can allocate towards channels.
	balReq := &lnrpc.WalletBalanceRequest{}
	walletBalance, err := l.lnd.WalletBalance(ctxb, balReq)
	if err != nil {
		return nil, err
	}

	nodeAddr := fmt.Sprintf("%v@%v:10011", nodeInfo.IdentityPubkey, *lndIP)
	return &homePageContext{
		NumCoins:    walletBalance.Balance,
		NumChannels: nodeInfo.NumActiveChannels,
		NodeAddr:    nodeAddr,
		NumConfs:    1,
	}, nil
}

// faucetHome renders the main home page for the faucet. This includes the form
// to create channels, the network statistics, and the splash page upon channel
// success.
//
// NOTE: This method implements the http.Handler interface.
func (l *lightningFaucet) faucetHome(w http.ResponseWriter, r *http.Request) {
	// First obtain the home template from our cache of pre-compiled
	// templates.
	homeTemplate := l.templates.Lookup("index.html")

	// In order to render the home template we'll need the necessary
	// context, so we'll grab that from the lnd daemon now in order to get
	// the most up to date state.
	homeInfo, err := l.fetchHomeState()
	if err != nil {
		log.Println("unable to fetch home state: ", err)
		homeTemplate.Execute(w, nil)
		return
	}

	// If the method is GET, then we'll render the home page with the form
	// itself.
	switch {
	case r.Method == http.MethodGet:
		if err := homeTemplate.Execute(w, homeInfo); err != nil {
			log.Println("unable to render home page: %v", err)
		}

	// Otherwise, if the method is POST, then the user is submitting the
	// form to open a channel, so we'll pass that off to the openChannel
	// handler.
	case r.Method == http.MethodPost:
		l.openChannel(homeTemplate, homeInfo, w, r)

	// If the method isn't either of those, then this is an error as we
	// only support the two methods above.
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}

	return
}

// channelExistsWithNode return true if the faucet already has a channel open
// with the target node, and false otherwise.
func (l *lightningFaucet) channelExistsWithNode(nodePub string) bool {
	listChanReq := &lnrpc.ListChannelsRequest{}
	resp, err := l.lnd.ListChannels(ctxb, listChanReq)
	if err != nil {
		return false
	}

	for _, channel := range resp.Channels {
		if channel.RemotePubkey == nodePub {
			return true
		}
	}

	return false
}

// connectedToNode returns true if the faucet is connected to the node, and
// false otherwise.
func (l *lightningFaucet) connectedToNode(nodePub string) bool {
	peersReq := &lnrpc.ListPeersRequest{}
	resp, err := l.lnd.ListPeers(ctxb, peersReq)
	if err != nil {
		return false
	}

	for _, peer := range resp.Peers {
		if peer.PubKey == nodePub {
			return true
		}
	}

	return false
}

// openChannel is a hybrid http.Handler that handles: the validation of the
// channel creation form, rendering errors to the form, and finally creating
// channels if all the parameters check out.
func (l *lightningFaucet) openChannel(homeTemplate *template.Template,
	homeState *homePageContext, w http.ResponseWriter, r *http.Request) {

	// Before we can obtain the values the user entered in the form, we
	// need to parse all parameters.  First attempt to establish a
	// connection with the
	if err := r.ParseForm(); err != nil {
		http.Error(w, "unable to parse form", 500)
		return
	}

	nodePubStr := r.FormValue("node")

	// With the forms details parsed, extract out the public key of the
	// target peer.
	nodePub, err := hex.DecodeString(nodePubStr)
	if err != nil {
		homeState.SubmissionError = InvalidAddress
		homeTemplate.Execute(w, homeState)
		return
	}

	// If we already have a channel with this peer, then we'll fail the
	// request as we have a policy of only one channel per node.
	if l.channelExistsWithNode(nodePubStr) {
		homeState.SubmissionError = HaveChannel
		homeTemplate.Execute(w, homeState)
	}

	// If we're not connected to the node, then we won't be able to extend
	// a channel to them. So we'll exit early with an error here.
	if !l.connectedToNode(nodePubStr) {
		homeState.SubmissionError = NotConnected
		homeTemplate.Execute(w, homeState)
		return
	}

	// With the connection established (or already present) with the target
	// peer, we'll now parse out the rest of the fields, performing
	// validation and exiting early if any field is invalid.
	chanSize, err := strconv.ParseInt(r.FormValue("amt"), 10, 64)
	if err != nil {
		homeState.SubmissionError = ChanAmountNotNumber
		homeTemplate.Execute(w, homeState)
		return
	}
	pushAmt, err := strconv.ParseInt(r.FormValue("bal"), 10, 64)
	if err != nil {
		homeState.SubmissionError = PushIncorrect
		homeTemplate.Execute(w, homeState)
		return
	}

	// With the initial validation complete, we'll now ensure the channel
	// size and push amt meet our constraints.
	switch {
	// The target channel can't be below the constant min channel size.
	case chanSize < minChannelSize:
		homeState.SubmissionError = ChannelTooSmall
		homeTemplate.Execute(w, homeState)
		return

	// The target channel can't be above the max channel size.
	case chanSize > maxChannelSize:
		homeState.SubmissionError = ChannelTooLarge
		homeTemplate.Execute(w, homeState)
		return

	// The amount pushed to the other side as part of the channel creation
	// MUST be less than the size of the channel itself.
	case pushAmt >= chanSize:
		homeState.SubmissionError = PushIncorrect
		homeTemplate.Execute(w, homeState)
		return
	}

	// If we were able to connect to the peer successfully, and all the
	// parameters check out, then we'll parse out the remaining channel
	// parameters and initiate the funding workflow.
	openChanReq := &lnrpc.OpenChannelRequest{
		NodePubkey:         nodePub,
		LocalFundingAmount: chanSize,
		PushSat:            pushAmt,
		NumConfs:           1,
	}
	log.Printf("attempting to create channel with params: %v",
		spew.Sdump(openChanReq))

	openChanStream, err := l.lnd.OpenChannel(ctxb, openChanReq)
	if err != nil {
		http.Error(w, "unable to open chnnel", 500)
		return
	}

	// Consume the first update from the open channel stream which
	// indicates that the channel has been broadcast to the network.
	chanUpdate, err := openChanStream.Recv()
	if err != nil {
		http.Error(w, "unable to open chnnel", 500)
		return
	}

	pendingUpdate := chanUpdate.Update.(*lnrpc.OpenStatusUpdate_ChanPending).ChanPending
	fundingTXID, _ := chainhash.NewHash(pendingUpdate.Txid)

	log.Printf("channel created with txid: %v", fundingTXID)

	homeState.ChannelTxid = fundingTXID.String()
	if err := homeTemplate.Execute(w, homeState); err != nil {
		log.Printf("unable to render home page: %v", err)
	}
}

// activeChannels returns a page describing all the currently active channels
// the faucet is maintaining.
func (l *lightningFaucet) activeChannels(w http.ResponseWriter, r *http.Request) {

	openChanReq := &lnrpc.ListChannelsRequest{}
	openChannels, err := l.lnd.ListChannels(ctxb, openChanReq)
	if err != nil {
		http.Error(w, "unable to fetch channels", 500)
		return
	}

	l.templates.Lookup("active.html").Execute(w, openChannels)
}

func (l *lightningFaucet) pendingChannels(w http.ResponseWriter, r *http.Request) {
}
