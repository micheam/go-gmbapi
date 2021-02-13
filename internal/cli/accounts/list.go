package accounts

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	gmbapi "github.com/micheam/go-gmbapi"
	"github.com/urfave/cli/v2"
)

var list = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "list available accounts",
	Action: func(c *cli.Context) error {
		client, err := gmbapi.New()
		if err != nil {
			return fmt.Errorf("can't create gmbapi client: %w", err)
		}
		accounts, err := client.AccountAccess().List(c.Context, url.Values{})
		if err != nil {
			return fmt.Errorf("failed to list accounts: %w", err)
		}
		for a := range accounts {
			acc := a
			b, err := json.Marshal(acc)
			if err != nil {
				fmt.Fprintf(os.Stderr,
					"cant marshal account(number=%s): %s", acc.AccountNumber, err.Error())
				continue
			}
			fmt.Println(string(b))
		}
		return nil
	},
}
