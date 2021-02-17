package accounts

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/micheam/go-gmbapi"
	"github.com/urfave/cli/v2"
)

type Presenter interface {
	Handle(out io.Writer, a gmbapi.Account) error
}

func GetPresenter(c *cli.Context) Presenter {
	style := strings.ToLower(c.String("output"))
	switch style {
	case "text":
		return new(TextPresenter)
	case "json":
		return new(JsonPresenter)
	}
	log.Printf("unknown format style: %s", style)
	return new(TextPresenter)
}

type TextPresenter struct{}

func (*TextPresenter) Handle(out io.Writer, a gmbapi.Account) (err error) {
	_, err = fmt.Fprintf(out, "%+20s\t%s\n", a.ID(), a.AccountName)
	return
}

type JsonPresenter struct{}

func (*JsonPresenter) Handle(out io.Writer, a gmbapi.Account) error {
	b, err := json.Marshal(a)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(out, "%s\n", string(b))
	return err
}
