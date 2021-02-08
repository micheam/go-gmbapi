package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/micheam/go-gmbapi"
)

func init() {
	// Overwrite with your credentials for GoogleApi.
	// os.Setenv("GMB_CLIENT_ID", "<YOUR_CLIENT_ID>")
	// os.Setenv("GMB_CLIENT_SECRET", "<YOUR_CLIENT_SECRET>")
	// os.Setenv("GMB_REFRESH_TOKEN", "<YOUR_REFRESH_TOKEN>")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s [flags] [account_name]:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "\nFLAGS:\n\n")
		flag.PrintDefaults()
	}
}

var all = flag.Bool("all", false, "list all locations available under your credential.")

func main() {
	flag.Parse()
	if !*all && len(flag.Args()) == 0 {
		fmt.Println("account_name required")
		os.Exit(1)
	}

	ctx := context.Background()
	client, err := gmbapi.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	accounts := make(chan *gmbapi.Account, 1)
	if *all {
		_accounts, err := client.AccountAccess().List(ctx, url.Values{})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		go func() {
			defer close(accounts)
			for acc := range _accounts {
				a := acc
				accounts <- a
			}
		}()
	} else {
		accounts = make(chan *gmbapi.Account, 1)
		go func() {
			defer close(accounts)
			a := flag.Arg(0)
			accounts <- &gmbapi.Account{Name: &a}
		}()
	}

	for a := range accounts {
		a := a
		locs, err := client.LocationAccess(a).List(ctx, url.Values{})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for loc := range locs {
			b, _ := json.Marshal(loc)
			fmt.Println(string(b))
		}
	}
}
