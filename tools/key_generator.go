// A tool for output private key and address with a given or random mnemonic and path, example:
//
// 1. go run tools/key_generator.go -mnemonic="xxx" -path="xxx" -wallet="btc-bip44"   // for a certain mnemonic and path for btc
// 2. go run tools/key_generator.go -wallet="eth"																	// for random mnemonic and path for eth

package main

import (
	"flag"
	"strings"

	"github.com/shawn-cx-li/wallet-core/pkg/coins/btc"
	"github.com/shawn-cx-li/wallet-core/pkg/coins/eth"
	"github.com/shawn-cx-li/wallet-core/pkg/crypto"
	"github.com/shawn-cx-li/wallet-core/pkg/interfaces"

	log "github.com/sirupsen/logrus"
)

const (
	ripplePath       = "m/44'/144'/0'/0/0"
	ethPath          = "m/44'/60'/0'/0/0"
	btcBIP44Path     = "m/44'/0'/0'/0/0"
	btcBIP49Path     = "m/49'/0'/0'/0/0"
	btcBIP44TESTPath = "m/44'/1'/0'/0/0"
	btcBIP49TESTPath = "m/49'/1'/0'/0/0"
	dashPath         = "m/44'/5'/0'/0/0"
)

var defaultPath = "m/44'/60'/0'/0/0"

func getParams() (mnemonic, path, wallet string) {
	defaultMnemonic, _ := crypto.GetNewMnemonic()
	flag.StringVar(&mnemonic, "mnemonic", defaultMnemonic, "Mnemonic of the HD Wallet")
	flag.StringVar(&path, "path", defaultPath, "Path of the private key")
	flag.StringVar(&wallet, "wallet", defaultPath, "Wallet to create")

	flag.Parse()
	if path == defaultPath {
		switch wallet {
		case "btc-bip44":
			path = btcBIP44Path
		case "eth":
			path = ethPath
		}
	}

	return
}

func getKey(mnemonic, path, wallet string) interfaces.Key {
	wallet = strings.ToLower(wallet)
	var key interfaces.Key

	switch wallet {
	case "btc-bip44":
		key, _ = btc.NewKey(mnemonic, path, btc.NewOpts(btc.BITCOIN, btc.BIP44))
	case "eth":
		key, _ = eth.NewKey(mnemonic, path, eth.Opts{})
	}

	return key
}

func main() {
	mnemonic, path, wallet := getParams()

	key := getKey(mnemonic, path, wallet)

	log.Infof("mnemonic: %s", mnemonic)
	log.Infof("path: %s", path)

	privKey := key.PrivateKeyString()
	log.Infof("privKey: %s", privKey)

	addr, _ := key.Address()
	log.Infof("addr: %s", addr)
}
