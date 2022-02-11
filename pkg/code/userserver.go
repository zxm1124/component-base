package code

import "net/http"

// userserver code begin 11 00 01
const (
	// ErrUserNotFoundCode - 404: User not found
	ErrUserNotFoundCode int = iota + 110001

	// ErrUserAlreadyExistCode - 400: User already exist.
	ErrUserAlreadyExistCode
)

var (
	ErrUserNotFound = NewCode(ErrUserNotFoundCode, http.StatusNotFound, "User not found")

	ErrUserAlreadyExist = NewCode(ErrUserAlreadyExistCode, http.StatusBadRequest, "User already exist")
)
