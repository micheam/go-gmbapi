package main

import (
	"fmt"
	"os"

	"github.com/micheam/go-gmbapi/internal/cli/accounts"
	"github.com/micheam/go-gmbapi/internal/cli/locations"
	"github.com/urfave/cli/v2"
)

var Version string = "0.1.0"

func main() {
	err := newApp().Run(os.Args)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "gmbcli"
	app.Usage = "cli client for Google My Business API"
	app.Version = Version
	app.Authors = []*cli.Author{
		{
			Name:  "Michito Maeda",
			Email: "michito.maeda@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Value:   "text",
			Usage:   "formatting `style` for output. (text|json)",
		},
	}
	app.Commands = commands
	return app
}

var commands = []*cli.Command{
	accounts.Commands,
	locations.Commands,
}
