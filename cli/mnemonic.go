package main

import (
	"github.com/shawn-cx-li/wallet-core/pkg/crypto"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var size int

func generateMnemonicCmd() *cli.Command {
	return &cli.Command{
		Name:    "generate-mnemonic",
		Aliases: []string{"gm"},
		Usage:   "Generate a random mnemonic with a given size",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "size",
				Usage:       "size of mnemonic to generate",
				Value:       256,
				Destination: &size,
			},
		},
		Action: newMnemonic,
	}
}

func newMnemonic(*cli.Context) (err error) {
	mnemonic, err := crypto.GetNewMnemonic(size)
	if err != nil {
		return err
	}

	log.Infof("mnemonic: %s", mnemonic)
	return nil
}
