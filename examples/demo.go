package main

import (
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

func demo_btc() {
	key, _ := btc.NewKey(mnemonic, btcBIP44Path, btc.NewOpts(btc.BITCOIN, btc.BIP44))
	privKey := key.PrivateKeyString()
	pubKey := key.PublicKeyString()
	addr, _ := key.Address()
	log.Info("btc private key: ", privKey)
	log.Info("btc public key: ", pubKey)
	log.Info("btc address: ", addr)
}

func demo_btc_testnet() {
	key, _ := btc.NewKey(mnemonic, btcBIP49TESTPath, btc.NewOpts(btc.BITCOIN_TESTNET, btc.BIP49))
	privKey := key.PrivateKeyString()
	pubKey := key.PublicKeyString()
	addr, _ := key.Address()
	log.Info("btc testnet private key: ", privKey)
	log.Info("btc testnet public key: ", pubKey)
	log.Info("btc testnet address: ", addr)
}

func demo_dash() {
	key, _ := btc.NewKey(mnemonic, dashPath, btc.NewOpts(btc.DASH, btc.BIP44))
	privKey := key.PrivateKeyString()
	pubKey := key.PublicKeyString()
	addr, _ := key.Address()
	log.Info("dash private key: ", privKey)
	log.Info("dash public key: ", pubKey)
	log.Info("dash address: ", addr)
}

func demo_eth() {
	key, _ := eth.NewKey(mnemonic, ethPath, eth.Opts{})
	privKey := key.PrivateKeyString()
	pubKey := key.PublicKeyString()
	addr, _ := key.Address()
	log.Info("eth private key: ", privKey)
	log.Info("eth public key: ", pubKey)
	log.Info("eth address: ", addr)
}

func demo_ripple() {
	key, _ := ripple.NewKey(mnemonic, ripplePath, ripple.Opts{})
	priv := key.PrivateKeyString()
	addr, _ := key.Address()

	log.Info("ripple privatekey: ", priv)
	log.Info("ripple address: ", addr)
}
