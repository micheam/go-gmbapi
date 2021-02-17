package locations

import (
	"fmt"
	"net/url"
	"os"

	"github.com/micheam/go-gmbapi"
	"github.com/urfave/cli/v2"
)

var Commands = &cli.Command{
	Name:    "locations",
	Aliases: []string{"loc"},
	Subcommands: []*cli.Command{
		list, get,
	},
}

var list = &cli.Command{
	Name:      "list",
	Aliases:   []string{"ls"},
	Usage:     "listing all locations under account",
	ArgsUsage: "[ACCOUNT_NAME or ACCOUNT_ID]",
	Action: func(c *cli.Context) error {
		client, err := gmbapi.New()
		if err != nil {
			return fmt.Errorf("failed to  create gmbapi client: %w", err)
		}
		parentName, err := gmbapi.ParseAccountName(c.Args().First())
		if err != nil {
			return err
		}
		ctx := c.Context
		parent, err := client.AccountAccess().Get(ctx, parentName)
		if err != nil {
			return fmt.Errorf("failed to get parent account info: %w", err)
		}
		locations, err := client.LocationAccess(parent).List(ctx, url.Values{})
		if err != nil {
			return fmt.Errorf("failed to list locations: %w", err)
		}
		p, out := GetPresenter(c), os.Stdout
		for i := range locations {
			if err := p.Handle(out, locations[i]); err != nil {
				return fmt.Errorf("locs(id=%s): %w",
					locations[i].StoreCode, err)
			}
		}
		return nil
	},
}

var get = &cli.Command{
	Name:      "get",
	Usage:     "get location by name",
	ArgsUsage: "[LOCATION_NAME]",
	Action: func(c *cli.Context) error {
		client, err := gmbapi.New()
		if err != nil {
			return fmt.Errorf("failed to create gmbapi client: %w", err)
		}
		if !c.Args().Present() {
			return fmt.Errorf("account Name or ID required")
		}
		parentName, err := gmbapi.ParseAccountName(c.Args().First())
		if err != nil {
			return err
		}
		ctx := c.Context
		parent, err := client.AccountAccess().Get(ctx, parentName)
		if err != nil {
			return fmt.Errorf("failed to get account: %w", err)
		}
		loc, err := client.LocationAccess(parent).Get(ctx, c.Args().First())
		if err != nil {
			return fmt.Errorf("failed to get location: %w", err)
		}
		return GetPresenter(c).Handle(os.Stdout, loc)
	},
}
