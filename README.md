## Lightning Network Faucet

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/lightninglabs/lightning-faucet/blob/master/LICENSE) 
&nbsp;&nbsp;&nbsp;&nbsp;
[![Irc](https://img.shields.io/badge/chat-on%20freenode-brightgreen.svg)](https://webchat.freenode.net/?channels=lnd) 

The Lightning Network Faucet is a faucet that is currently deployed on the
Bitcoin and Litecoin testnets.

- Bitcoin: https://faucet.lightning.community
- Litecoin: https://ltc.faucet.lightning.community

The Testnet Lightning Faucet (TLF) is similar to other existing Bitcoin
faucets.  However, rather then sending bitcoin directly on-chain to a user of
the faucet, the TLF will instead open a payment channel with the target user.
The user can then either use their new link to the Lightning Network to
facilitate payments, or immediately close the channel (which immediately
credits them on-chain like regular faucets).

Currently the TLF is only compatible with `lnd`, but in the near future as the
other Lightning implementations ([eclair](https://github.com/ACINQ/eclair/),
[c-lightning](https://github.com/ElementsProject/lightning),
[lit](https://github.com/ElementsProject/lightning), and more) become
compatible, the faucet will also be usable between all the active
implementations.

## Installation
  In order to build from source, the following build dependencies are 
  required:
  
  * **Go:** Installation instructions can be found [here](http://golang.org/doc/install). 
  
    It is recommended to add `$GOPATH/bin` to your `PATH` at this point.
    **Note:** If you are building with `Go 1.5`, then you'll need to 
    enable the vendor experiment by setting the `GO15VENDOREXPERIMENT` 
    environment variable to `1`. If you're using `Go 1.6` or later, then
    it is safe to skip this step.

  * **Glide:** This project uses `Glide` to manage dependencies as well 
    as to provide *reproducible builds*. To install `Glide`, execute the
    following command (assumes you already have Go properly installed):
    ```
    $ go get -u github.com/Masterminds/glide
    ```

With the preliminary steps completed, to install the Testnet Lightning Faucet
```
$ git clone https://github.com/lightninglabs/lightning-faucet $GOPATH/src/github.com/lightninglabs/lightning-faucet
$ cd $GOPATH/src/github.com/lightninglabs/lightning-faucet
$ glide install
$ go install -v
```

## Deploying The Faucet

Once you have the faucet installed, you'll need to ensure you have a local
[`lnd`](https://github.com/lightningnetwork/lnd) active and fully synced.

Once the node is synced, execute the following command (from this directory) to
deploy the faucet:
```
lightning-faucet --lnd_ip=X.X.X.X --lnd_port=9735
```

Where `X.X.X.X` is the public, reachable IP address for your active `lnd` node.
The configuration for the faucet includes a TLS certificate provided by [Let's
Encrypt](https://letsencrypt.org) so your faucet will be reachable over `https`
by default.


### Updating
To update your version of the `TLF` to the latest version run the following
commands:
```
$ cd $GOPATH/src/github.com/lightninglabs/lightning-faucet
$ git pull && glide install
$ go install -v
```
