package gmbapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/micheam/go-gmbapi/internal/config"
)

// constants for call apis
var (
	BaseEndpoint   string = "https://mybusiness.googleapis.com/v4"
	Oauth2Endpoint string = "https://www.googleapis.com/oauth2/v4/token"
)

// Credential define credential info for GMB Client
type Credential interface {
	GetClientID() string
	GetClientSecret() string
	GetRefreshToken() string
}

// Client ...
type Client struct {
	Cred Credential
	lock sync.RWMutex
	*Token
}

// New ...
func New() (*Client, error) {
	c, err := config.Load()
	if err != nil {
		return nil, err
	}
	return &Client{Cred: c}, nil
}

func doRequest(ctx context.Context, method, url string, values url.Values) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, strings.NewReader(values.Encode()))
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
const maxRetry uint64 = 3

func (c *Client) doRequest(ctx context.Context, basetime time.Time, method, _url string, body io.ReadSeeker, param url.Values) ([]byte, error) {
	if c.tokenExpired(basetime) {
		if err := c.tokenReflesh(ctx); err != nil {
			return nil, fmt.Errorf("failed to reflesh token: %w", err)
		}
	}

	var result []byte
	op := func() error {
		var err error
		if body != nil {
			if _, err := body.Seek(0, 0); err != nil {
				err = fmt.Errorf("failed to seek to head of body: %w", err)
				return backoff.Permanent(err)
			}
		}
		req, err := http.NewRequestWithContext(ctx, method, _url, body)
		if err != nil {
			err = fmt.Errorf("failed to create http request: %w", err)
			return backoff.Permanent(err)
		}
		req.Header.Set("Authorization", "Bearer "+c.AccessToken)
		req.URL.RawQuery = param.Encode()
		resp, err := new(http.Client).Do(req)
		if err != nil {
			err = fmt.Errorf("failed to do http-request: %w", err)
			return backoff.Permanent(err)
		}
		defer resp.Body.Close()

		result, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			err = fmt.Errorf("failed to read response body: %w", err)
			return backoff.Permanent(err)
		}

		if 300 <= resp.StatusCode && resp.StatusCode < 500 {
			return backoff.Permanent(errors.New(string(result)))
		} else if 500 <= resp.StatusCode {
			return errors.New(string(result)) // !! Retry !!
		}
		return nil
	}
	bf := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), maxRetry)
	if err := backoff.Retry(op, bf); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) tokenExpired(on time.Time) bool {
	return c.Token.Expired(on)
}

func (c *Client) tokenReflesh(ctx context.Context) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	b, err := doRequest(ctx, http.MethodPost, Oauth2Endpoint, url.Values{
		"client_id":     []string{c.Cred.GetClientID()},
		"client_secret": []string{c.Cred.GetClientSecret()},
		"grant_type":    []string{"refresh_token"},
		"refresh_token": []string{c.Cred.GetRefreshToken()}})

	if err != nil {
		return fmt.Errorf("failed to do request: %w", err)
	}

	c.Token = new(Token)
	if err = json.Unmarshal(b, c.Token); err != nil {
		return fmt.Errorf("failed to ummarshal response: %w", err)
	}
	return nil
}
