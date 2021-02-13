package gmbapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

func ExampleAccountAccess_Streaming() {
	ctx := context.Background()
	client, _ := New()
	accounts := client.AccountAccess().Streaming(ctx, url.Values{})

	for chunk := range accounts {
		if chunk.Err != nil {
			fmt.Println(chunk.Err)
			return
		}
		for i := range chunk.Accounts {
			a := chunk.Accounts[i]
			b, _ := json.Marshal(a)
			fmt.Println(string(b))
		}
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

func ExampleAccountAccess_Streaming_withcredential() {
	var (
		ctx            = context.Background()
		c   Credential = &myCred{
			clientID:     "my client id",
			clientSecret: "my client secret",
			refreshToken: "my refresh token",
		}
	)
	client := &Client{Cred: c}
	accounts := client.AccountAccess().Streaming(ctx, url.Values{})
	for chunk := range accounts {
		for i := range chunk.Accounts {
			a := chunk.Accounts[i]
			b, _ := json.Marshal(a)
			fmt.Println(string(b))
		}
	}
	// will print all your accounts
}
