package gmbapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

func ExampleAccountAccess_List() {
	ctx := context.Background()
	client, _ := New()
	accounts, _ := client.AccountAccess().List(ctx, url.Values{})
	for account := range accounts {
		b, _ := json.Marshal(account)
		fmt.Println(string(b))
	}
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
	var (
		ctx            = context.Background()
		c   Credential = &myCred{
			clientID:     "my client id",
			clientSecret: "my client secret",
			refreshToken: "my refresh token",
		}
	)
	client := &Client{Cred: c}
	accounts, _ := client.AccountAccess().List(ctx, url.Values{})
	for account := range accounts {
		b, _ := json.Marshal(account)
		fmt.Println(string(b))
	}
	// will print all your accounts
}

func ExampleLocationAccess_List() {
	ctx := context.Background()
	client, _ := New()
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
