package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	gmbapi "github.com/micheam/google-my-business-go"
)

func init() {
	// Overwrite with your credentials for GoogleApi.
	// os.Setenv("GMB_CLIENT_ID", "<YOUR_CLIENT_ID>")
	// os.Setenv("GMB_CLIENT_SECRET", "<YOUR_CLIENT_SECRET>")
	// os.Setenv("GMB_REFRESH_TOKEN", "<YOUR_REFRESH_TOKEN>")
}

func main() {
	ctx := context.Background()
	client, err := gmbapi.New()
	if err != nil {
		log.Fatal(err)
	}
	accounts, err := client.AccountAccess().List(ctx, url.Values{})
	if err != nil {
		log.Fatal(err)
	}

	var locations []*gmbapi.Location
	for acc := range accounts {
		acc := acc
		locs, err := client.LocationAccess(acc).List(ctx, url.Values{})
		if err != nil {
			log.Fatal(err)
		}
		locations = append(locations, locs.Locations...)
	}
	b, _ := json.Marshal(locations)
	fmt.Println(string(b))
}
