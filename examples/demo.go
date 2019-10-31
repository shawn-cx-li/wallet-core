package main

import (
	"github.com/shawn-cx-li/wallet-core/pkg/base58"
	"github.com/shawn-cx-li/wallet-core/pkg/coins/btc"
	"github.com/shawn-cx-li/wallet-core/pkg/coins/eth"
	"github.com/shawn-cx-li/wallet-core/pkg/coins/ripple"
	log "github.com/sirupsen/logrus"
)

const (
	mnemonic         = "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"
	ripplePath       = "m/44'/144'/0'/0/0"
	ethPath          = "m/44'/60'/0'/0/0"
	btcBIP44Path     = "m/44'/0'/0'/0/0"
	btcBIP49Path     = "m/49'/0'/0'/0/0"
	btcBIP44TESTPath = "m/44'/1'/0'/0/0"
	btcBIP49TESTPath = "m/49'/1'/0'/0/0"
	dashPath         = "m/44'/5'/0'/0/0"
)

func demo_ripple() {
	key, _ := ripple.NewKey(mnemonic, ripplePath, ripple.Opts{})
	addr, _ := key.Address()
	priv, err := key.PrivateKeyString()
	if err != nil {
		log.Error(err)
	}
	log.Info("ripple address: ", addr)
	log.Info("ripple privatekey: ", priv)
}

func demo_btc() {
	key, _ := btc.NewKey(mnemonic, btcBIP44Path, btc.NewOpts(btc.BITCOIN, btc.BIP44))
	addr, _ := key.Address()
	pubKey, _ := key.PublicKeyString()
	log.Info("btc address: ", addr)
	log.Info("btc public key: ", pubKey)
	d, _ := base58.Decode(addr, btc.ALPHABET)
	log.Info("after: ", d)
}

func demo_btc_testnet() {
	key, _ := btc.NewKey(mnemonic, btcBIP49TESTPath, btc.NewOpts(btc.BITCOIN_TESTNET, btc.BIP49))
	addr, _ := key.Address()
	log.Info("btc testnet address: ", addr)
}

func demo_dash() {
	key, _ := btc.NewKey(mnemonic, dashPath, btc.NewOpts(btc.DASH, btc.BIP44))
	addr, _ := key.Address()
	log.Info("dash address: ", addr)
}

func demo_bch() {
	key, _ := btc.NewKey(mnemonic, "m/44'/145'/0'/0/0", btc.NewOpts(btc.BITCOIN, btc.BIP44))
	addr, _ := key.Address()
	pubKey, _ := key.PublicKeyString()
	log.Info("bch address: ", addr)
	log.Info("bch public key: ", pubKey)
}

func demo_eth() {
	addr, _ := eth.GenerateAddress(mnemonic, ethPath, "")
	log.Info("address: ", addr)
}
