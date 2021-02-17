package accounts

import (
	"fmt"
	"net/url"
	"os"

	gmbapi "github.com/micheam/go-gmbapi"
	"github.com/urfave/cli/v2"
)

var Commands = &cli.Command{
	Name:    "accounts",
	Aliases: []string{"acc"},
	Subcommands: []*cli.Command{
		list, get,
	},
}

var list = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "listing all available accounts",
	Action: func(c *cli.Context) error {
		client, err := gmbapi.New()
		if err != nil {
			return fmt.Errorf("failed to  create gmbapi client: %w", err)
		}
		accounts, err := client.AccountAccess().List(c.Context, url.Values{})
		if err != nil {
			return fmt.Errorf("failed to list accounts: %w", err)
		}
		p, out := GetPresenter(c), os.Stdout
		for i := range accounts {
			if err := p.Handle(out, accounts[i]); err != nil {
				return fmt.Errorf("account(number=%s): %w",
					accounts[i].AccountNumber, err)
			}
		}
		return nil
	},
}

var get = &cli.Command{
	Name:      "get",
	Usage:     "get account by id or name",
	ArgsUsage: "[ACCOUNT_NAME or ACCOUNT_ID]",
	Action: func(c *cli.Context) error {
		client, err := gmbapi.New()
		if err != nil {
			return fmt.Errorf("failed to  create gmbapi client: %w", err)
		}
		if !c.Args().Present() {
			return fmt.Errorf("account Name or ID required")
		}
		name, err := gmbapi.ParseAccountName(c.Args().First())
		if err != nil {
			return err
		}
		account, err := client.AccountAccess().Get(c.Context, name)
		if err != nil {
			return fmt.Errorf("failed to gete account: %w", err)
		}
		p, out := GetPresenter(c), os.Stdout
		return p.Handle(out, account)
	},
}
