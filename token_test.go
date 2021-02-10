package gmbapi

import (
	"testing"
	"time"
)

func TestExpired(t *testing.T) {
	now := time.Now()
	sut := &Token{ExpiresIn: now.Unix()}

	t.Run("true if token is nil", func(t *testing.T) {
		want := true
		got := (*Token)(nil).Expired(now.Add(5 * time.Second))
		if want != got {
			t.Errorf("want %t, but got %t", want, got)
		}
	})
	t.Run("true if before token.ExpiresIn", func(t *testing.T) {
		want := true
		got := sut.Expired(now.Add(-1 * time.Second))
		if want != got {
			t.Errorf("want %t, but got %t", want, got)
		}
	})
	t.Run("false if equal token.ExpiresIn", func(t *testing.T) {
		want := false
		got := sut.Expired(now)
		if want != got {
			t.Errorf("want %t, but got %t", want, got)
		}
	})
}
