// A tool for output ETH private key and address with a given or random mnemonic and path, example:
//
// 1. go run tools/eth.go -mnemonic="xxx" -path="xxx"   // for a certain mnemonic and path
// 2. go run tools/eth.go 								// for random mnemonic and path

package main

import (
	"flag"

	"github.com/shawn-cx-li/wallet-core/pkg/coins/eth"
	"github.com/shawn-cx-li/wallet-core/pkg/crypto"

	log "github.com/sirupsen/logrus"
)

var defaultPath = "m/44'/60'/0'/0/0"

func getParams() (mnemonic, path string) {
	defaultMnemonic, _ := crypto.GetNewMnemonic()
	flag.StringVar(&mnemonic, "mnemonic", defaultMnemonic, "Mnemonic of the HD Wallet")
	flag.StringVar(&path, "path", defaultPath, "Path of the private key")

	flag.Parse()

	return
}

func main() {
	mnemonic, path := getParams()

	key, _ := eth.NewKey(mnemonic, path, eth.Opts{})

	privKey := key.PrivateKeyString()
	log.Infof("privKey: %s", privKey)

	addr, _ := key.Address()
	log.Infof("addr: %s", addr)
}
