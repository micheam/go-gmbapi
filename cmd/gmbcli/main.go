package main

import (
	"os"

	"github.com/micheam/go-gmbapi/internal/cli/accounts"
	"github.com/urfave/cli/v2"
)

var Version string = "0.1.0"

func main() {
	newApp().Run(os.Args)
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
	app.Commands = commands
	return app
}

var commands = []*cli.Command{
	accounts.Commands,
}
