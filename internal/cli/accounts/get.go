package accounts

import (
	"errors"
	"fmt"
	"os"
	"strings"

	gmbapi "github.com/micheam/go-gmbapi"
	"github.com/urfave/cli/v2"
)

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
		name, err := parseArg(c.Args().First())
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

var ErrInvalidAccountName = errors.New("invalid account name")

// TODO(micheam): move under package gmbapi
func parseArg(s string) (accountName string, err error) {
	ss := strings.Split(s, "/")
	if len(ss) > 0 {
		if ss[0] != "accounts" {
			return "", ErrInvalidAccountName
		}
		return s, nil
	}
	return fmt.Sprintf("accounts/%s", s), nil
}
