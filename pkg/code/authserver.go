package code

import "net/http"

const (
	ErrNotAuthorizedCode int = iota + 100201
)

var (
	// ErrNotAuthorized User not authorized
	ErrNotAuthorized = NewCode(ErrNotAuthorizedCode, http.StatusUnauthorized, "User not authorized")
)
