package gmbapi

import "errors"

var (
	// ErrNotFound will occurs on no entity found.
	ErrNotFound = errors.New("Not Found")
)
