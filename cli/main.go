package main

import (
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "wallet cli"
	app.Usage = "A wallet cli for creating keys and signing transaction"
	app.Author = "shawn-cx-li"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Create wallet with given mnemonic and path for family",
			Action:  create,
		},
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
