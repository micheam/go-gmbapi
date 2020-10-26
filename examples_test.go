package gmbapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func ExampleAccountAccess_List() {
	var (
		clientID     = "<GOOGLEAPI_CLIENT_ID>"
		clientSecret = "<GOOGLEAPI_CLIENT_SECRET>"
		refreshToken = "<GOOGLEAPI_REFRESH_TOKEN>"
	)

	client, _ := New(clientID, clientSecret, refreshToken)
	accounts, _ := client.AccountAccess().List(url.Values{})

	// Print all Accounts
	b, _ := json.Marshal(accounts)
	fmt.Println(string(b))
}

func ExampleLocationAccess_List() {
	var (
		clientID     = "<GOOGLEAPI_CLIENT_ID>"
		clientSecret = "<GOOGLEAPI_CLIENT_SECRET>"
		refreshToken = "<GOOGLEAPI_REFRESH_TOKEN>"
	)

	client, _ := New(clientID, clientSecret, refreshToken)
	accounts, _ := client.AccountAccess().List(url.Values{})

	// Print all Locations under account.
	var locations []*Location
	for _, acc := range accounts.Accounts {
		acc := acc
		locs, _ := client.LocationAccess(acc).List(url.Values{})
		locations = append(locations, locs.Locations...)
	}
	b, _ := json.Marshal(locations)
	fmt.Println(string(b))
}
