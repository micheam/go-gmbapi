package accounts

import (
	"encoding/json"
	"fmt"
	"net/url"

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
		accounts := client.AccountAccess().Streaming(c.Context, url.Values{})
		if err != nil {
			return fmt.Errorf("failed to list accounts: %w", err)
		}

		for chunk := range accounts {
			if chunk.Err != nil {
				return chunk.Err
			}
			for i := range chunk.Accounts {
				acc := chunk.Accounts[i]
				b, err := json.Marshal(acc)
				if err != nil {
					return fmt.Errorf("account(number=%s): %w", acc.AccountNumber, err)
				}
				fmt.Println(string(b))
			}
		}
		return nil
	},
}
