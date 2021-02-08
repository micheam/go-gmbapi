package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	client, err := gmbapi.New()
	if err != nil {
		log.Fatal(err)
	}
	accounts, err := client.AccountAccess().List(ctx, url.Values{})
	if err != nil {
		log.Fatal(err)
	}

	for acc := range accounts {
		acc := acc
		b, _ := json.Marshal(acc)
		fmt.Println(string(b))
	}
}
