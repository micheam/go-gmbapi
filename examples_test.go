package gmbapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func ExampleAccountAccess_List() {
	client, _ := New()
	accounts, _ := client.AccountAccess().List(url.Values{})
	b, _ := json.Marshal(accounts)
	fmt.Println(string(b))
	// will print all your accounts
}

type myCred struct {
	clientID     string
	clientSecret string
	refreshToken string
}

func (c *myCred) GetClientID() string {
	return c.clientID
}
func (c *myCred) GetClientSecret() string {
	return c.clientSecret
}
func (c *myCred) GetRefreshToken() string {
	return c.refreshToken
}

func ExampleAccountAccess_List_withcredential() {
	var c Credential = &myCred{
		clientID:     "my client id",
		clientSecret: "my client secret",
		refreshToken: "my refresh token",
	}

	client := &Client{Cred: c}
	accounts, _ := client.AccountAccess().List(url.Values{})
	b, _ := json.Marshal(accounts)
	fmt.Println(string(b))
	// will print all your accounts
}

func ExampleLocationAccess_List() {
	client, _ := New()
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
