package accounts

import "github.com/urfave/cli/v2"

var Commands = &cli.Command{
	Name:    "accounts",
	Aliases: []string{"acc"},
	Subcommands: []*cli.Command{
		list, get,
	},
}
