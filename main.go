package main

import (
	"context"
	"crypto/tls"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"time"

	"golang.org/x/crypto/acme/autocert"

	"github.com/gorilla/mux"
)

var (
	// netParams is the Bitcoin network that the faucet is operating on.
	netParams = flag.String("net", "testnet", "bitcoin network to operate on")

	// lndNodes is a list of lnd nodes that the faucet should connect out
	// to.
	//
	// TODO(roasbeef): channels should be balanced in a round-robin manner
	// between all available lnd nodes.
	lndNodes = flag.String("nodes", "localhost:10009", "comma separated "+
		"list of host:port")

	// lndIP is the IP address that should be advertised on the home page.
	// Users can connect out to this address in order to sync up their
	// graph state.
	lndIP = flag.String("lnd_ip", "10.0.0.9", "the public IP address of "+
		"the faucet's node")

	// port is the port that the http server should listen on.
	port = flag.String("port", "8080", "port to list for http")
)

// equal reports whether the first argument is equal to any of the remaining
// arguments. This function is used as a custom function within templates to do
// richer equality tests.
func equal(x, y interface{}) bool {
	if reflect.DeepEqual(x, y) {
		return true
	}

	return false
}

var (
	// templateGlobPattern is the pattern than matches all the HTML
	// templates in the static directory
	templateGlobPattern = filepath.Join(staticDirName, "*.html")

	// customFuncs is a registry of custom functions we use from within the
	// templates.
	customFuncs = template.FuncMap{
		"equal": equal,
	}

	// ctxb is a global context with no timeouts that's used within the
	// gRPC requests to lnd.
	ctxb = context.Background()
)

const (
	staticDirName = "static"
)

func main() {
	flag.Parse()

	// Pre-compile the list of templates so we'll catch any error sin the
	// templates as soon as the binary is run.
	faucetTemplates := template.Must(template.New("faucet").
		Funcs(customFuncs).
		ParseGlob(templateGlobPattern))

	// With the templates
	faucet, err := newLightningFaucet(*lndNodes, faucetTemplates)
	if err != nil {
		log.Fatalf("unable to create faucet: %v", err)
		return
	}

	// Create a new mux in order to route a request based on its path to a
	// dedicated http.Handler.
	r := mux.NewRouter()
	r.HandleFunc("/", faucet.faucetHome).Methods("POST", "GET")
	r.HandleFunc("/channels/active", faucet.activeChannels)
	r.HandleFunc("/channels/pending", faucet.activeChannels)

	// Next create a static file server which will dispatch our static
	// files. We rap the file sever http.Handler is a handler that strips
	// out the absolute file path since it'll dispatch based on solely the
	// file name.
	staticFileServer := http.FileServer(http.Dir(staticDirName))
	staticHandler := http.StripPrefix("/static/", staticFileServer)
	r.PathPrefix("/static/").Handler(staticHandler)

	// With all of our paths registered we'll register our mux as part of
	// the global http handler.
	http.Handle("/", r)

	// Create a directory cache so the certs we get from Let's Encrypt are
	// cached locally. This avoids running into their rate-limiting by
	// requesting too many certs.
	certCache := autocert.DirCache("certs")

	// Create the auto-cert manager which will automatically obtain a
	// certificate provided by Let's Encrypt.
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      certCache,
		HostPolicy: autocert.HostWhitelist("faucet.lightning.community"),
	}

	// As we'd like all requests to default to https, redirect all regular
	// http requests to the https version of the faucet.
	go http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		targetURL := "https://" + r.Host + r.URL.String()
		if len(r.URL.RawQuery) > 0 {
			targetURL += "?" + r.URL.RawQuery
		}

		http.Redirect(w, r, targetURL, http.StatusPermanentRedirect)
	}))

	// Finally, create the http server, passing in our TLS configuration.
	httpServer := &http.Server{
		Handler:      r,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		Addr:         ":https",
		TLSConfig:    &tls.Config{GetCertificate: m.GetCertificate},
	}
	if err := httpServer.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
