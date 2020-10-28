package config

import "github.com/kelseyhightower/envconfig"

// Credential ...
type Credential struct {
	ClientID     string `envconfig:"client_id"`
	ClientSecret string `envconfig:"client_secret"`
	RefreshToken string `envconfig:"refresh_token"`
}

// GetClientID return ClientID string
func (c *Credential) GetClientID() string {
	return c.ClientID
}

// GetClientSecret return ClientSecret string
func (c *Credential) GetClientSecret() string {
	return c.ClientSecret
}

// GetRefreshToken return RefreshToken string
func (c *Credential) GetRefreshToken() string {
	return c.RefreshToken
}

// Load ...
func Load() (*Credential, error) {
	var c = new(Credential)
	err := envconfig.Process("GMB", c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
