package gmbapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/cenkalti/backoff"
)

// constants for call apis
var (
	BaseEndpoint   string = "https://mybusiness.googleapis.com/v4"
	Oauth2Endpoint string = "https://www.googleapis.com/oauth2/v4/token"
)

// Token ...
type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	IDToken     string `json:"id_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// Client ...
type Client struct {
	*Token
}

// New ...
func New(clientID, clientSecret, refreshToken string) (*Client, error) {
	var err error
	var b []byte
	if b, err = doRequest(
		http.MethodPost, Oauth2Endpoint,
		url.Values{
			"client_id":     []string{clientID},
			"client_secret": []string{clientSecret},
			"grant_type":    []string{"refresh_token"},
			"refresh_token": []string{refreshToken},
		}); err != nil {
		return nil, err
	}
	var token = new(Token)
	if err = json.Unmarshal(b, token); err != nil {
		return nil, err
	}
	return &Client{Token: token}, nil
}

func doRequest(method, url string, values url.Values) ([]byte, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(values.Encode()))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal google api response: %w", err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// TODO(micheam): May be useful to be able to specify externally.
const maxRetry uint64 = 4

func (c *Client) doRequest(method, url string, body io.Reader, param url.Values) ([]byte, error) {
	var result []byte
	op := func() error {
		var err error
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			return fmt.Errorf("failed to create http request: %w", err)
		}
		req.Header.Set("Authorization", "Bearer "+c.AccessToken)
		req.URL.RawQuery = param.Encode()
		resp, err := new(http.Client).Do(req)
		if err != nil {
			return fmt.Errorf("failed to do http-request: %w", err)
		}
		defer resp.Body.Close()

		result, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		if resp.StatusCode >= http.StatusMultipleChoices {
			return fmt.Errorf(string(result))
		}
		log.Println(resp.StatusCode)
		return nil
	}
	boff := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), maxRetry)
	if err := backoff.Retry(op, boff); err != nil {
		return nil, err
	}
	return result, nil
}
