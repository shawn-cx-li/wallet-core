package main

import (
	"strings"

	"github.com/shawn-cx-li/wallet-core/pkg/coins/btc"
	"github.com/shawn-cx-li/wallet-core/pkg/coins/eth"
	"github.com/shawn-cx-li/wallet-core/pkg/coins/ripple"
	"github.com/shawn-cx-li/wallet-core/pkg/interfaces"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

type Wallet struct {
	Mnemonic string
	Path     string
	Family   string
}

const (
	defaultConfigFile = "./config/config.yaml"
)

var wallet Wallet

func (w *Wallet) createWallet(*cli.Context) error {
	w.printKeys()
	return nil
}

func createWalletCmd() *cli.Command {
	wallet = Wallet{}

	var flags = []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "mnemonic",
			Usage:       "mnemonic of the HD Wallet",
			Destination: &wallet.Mnemonic,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "path",
			Usage:       "path to derive private key",
			Destination: &wallet.Path,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "family",
			Usage:       "wallet family to create",
			Value:       "btc-bip44",
			Destination: &wallet.Family,
		}),
		&cli.StringFlag{
			Name:  "config",
			Value: defaultConfigFile,
		},
	}

	return &cli.Command{
		Name:    "create",
		Aliases: []string{"c"},
		Usage:   "Create wallet with given mnemonic and path for family",
		Before:  altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config")),
		Flags:   flags,
		Action:  wallet.createWallet,
	}
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

func (w *Wallet) printKeys() {
	key, err := getKey(w.Mnemonic, w.Path, w.Family)
	if err != nil {
		log.Panic(err)
	}

	log.Infof("mnemonic: %s", w.Mnemonic)
	log.Infof("path: %s", w.Path)

	privKey := key.PrivateKeyString()
	log.Infof("privKey: %s", privKey)

	addr, _ := key.Address()
	log.Infof("addr: %s", addr)
}
