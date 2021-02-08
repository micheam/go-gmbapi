package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	gmbapi "github.com/micheam/go-gmbapi"
)

func init() {
	// Overwrite with your credentials for GoogleApi.
	// os.Setenv("GMB_CLIENT_ID", "<YOUR_CLIENT_ID>")
	// os.Setenv("GMB_CLIENT_SECRET", "<YOUR_CLIENT_SECRET>")
	// os.Setenv("GMB_REFRESH_TOKEN", "<YOUR_REFRESH_TOKEN>")
}

func main() {
	ctx := context.Background()
	client, _ := gmbapi.New()
	accounts, _ := client.AccountAccess().List(ctx, url.Values{})

	for acc := range accounts {
		acc := acc

		locs, _ := client.LocationAccess(acc).List(ctx, url.Values{})
		for loc := range locs {
			b, _ := json.Marshal(loc)
			fmt.Println(string(b))
		}
	}
}
