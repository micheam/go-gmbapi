package gmbapi

import (
	"time"
)

// Token ...
type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	IDToken     string `json:"id_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// Expired ...
func (t *Token) Expired(when time.Time) bool {
	if t == nil {
		return true
	}
	return when.Unix() < t.ExpiresIn
}
