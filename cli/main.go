package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func info() {
	app.Name = "wallet-core cli from shawn-cx-li"
	app.Usage = "A wallet cli for creating keys and signing transaction"
	app.Version = "0.0.1"
}

func commands() {
	app.Commands = []*cli.Command{
		createWalletCmd(),
		generateMnemonicCmd(),
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
