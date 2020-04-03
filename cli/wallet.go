package main

import (
	"strings"

	"github.com/shawn-cx-li/wallet-core/pkg/coins/btc"
	"github.com/shawn-cx-li/wallet-core/pkg/coins/eth"
	"github.com/shawn-cx-li/wallet-core/pkg/coins/ripple"
	"github.com/shawn-cx-li/wallet-core/pkg/interfaces"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type Wallet struct {
	mnemonic string
	path     string
}

func getKey(mnemonic, path, family string) (key interfaces.Key, err error) {
	family = strings.ToLower(family)

	switch family {
	case "btc-bip44":
		key, err = btc.NewKey(mnemonic, path, btc.NewOpts(btc.BITCOIN, btc.BIP44))
	case "btc-bip44-test":
		key, err = btc.NewKey(mnemonic, path, btc.NewOpts(btc.BITCOIN_TESTNET, btc.BIP44))
	case "btc-bip49":
		key, err = btc.NewKey(mnemonic, path, btc.NewOpts(btc.BITCOIN, btc.BIP49))
	case "btc-bip49-test":
		key, err = btc.NewKey(mnemonic, path, btc.NewOpts(btc.BITCOIN_TESTNET, btc.BIP49))
	case "eth":
		key, err = eth.NewKey(mnemonic, path, eth.Opts{})
	case "ripple":
		key, err = ripple.NewKey(mnemonic, path, ripple.Opts{})
	}
	return
}

func NewWallet(mnemonic, path string) (w *Wallet) {
	return &Wallet{mnemonic, path}
}

func create(*cli.Context) {
	conf, err := getConfig()
	if err != nil {
		panic(err)
	}

	wallet := NewWallet(conf.Wallet.Mnemonic, conf.Wallet.Path)
	wallet.printKeys(conf.Family)
}

func (w *Wallet) printKeys(family string) {
	key, err := getKey(w.mnemonic, w.path, family)
	if err != nil {
		log.Panic(err)
	}

	log.Infof("mnemonic: %s", w.mnemonic)
	log.Infof("path: %s", w.path)

	privKey := key.PrivateKeyString()
	log.Infof("privKey: %s", privKey)

	addr, _ := key.Address()
	log.Infof("addr: %s", addr)
}
