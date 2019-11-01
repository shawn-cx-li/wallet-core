package main

import (
	"github.com/shawn-cx-li/wallet-core/pkg/coins/eth"
	"github.com/shawn-cx-li/wallet-core/pkg/crypto"

	log "github.com/sirupsen/logrus"
)

var path = "m/44'/60'/0'/0/0"

func main() {
	mnemonic, _ := crypto.GetNewMnemonic()

	log.Infof("mnemonic: %s", mnemonic)

	key, _ := eth.NewKey(mnemonic, path, eth.Opts{})

	privKey, _ := key.PrivateKeyString()
	log.Infof("privKey: %s", privKey)

	addr, _ := key.Address()
	log.Infof("addr: %s", addr)
}
