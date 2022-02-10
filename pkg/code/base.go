package code

import "net/http"

// Common: basic code
const (
	// ErrSuccessCode - 200: OK.
	ErrSuccessCode int = iota + 100001

	// ErrUnknownCode - 500: Internal server error.
	ErrUnknownCode

	// ErrBindCode - 400: Error occurred while binding the request body to the struct.
	ErrBindCode

	// ErrValidationCode - 400: Validation failed.
	ErrValidationCode

	// ErrTokenInvalidCode - 401: Token invalid.
	ErrTokenInvalidCode

	// ErrPageNotFoundCode - 404: Page not found.
	ErrPageNotFoundCode
)

// common: database
const (
	// ErrDatabaseCode database error
	ErrDatabaseCode int = iota + 100101
)

const ()

var (
	// ErrSuccess - 200: OK.
	ErrSuccess = NewCode(ErrSuccessCode, http.StatusOK, "OK")

	// ErrUnknown - 500: Internal server error.
	ErrUnknown = NewCode(ErrUnknownCode, http.StatusInternalServerError, "Internal server error")

	// ErrBind - 400: Error occurred while binding the request body to the struct.
	ErrBind = NewCode(ErrBindCode, http.StatusBadRequest, "Error occurred while binding the request body to the struct")

	// ErrValidation - 400: Validation failed.
	ErrValidation = NewCode(ErrValidationCode, http.StatusBadRequest, "Validation failed")

	// ErrTokenInvalid - 401: Token invalid.
	ErrTokenInvalid = NewCode(ErrTokenInvalidCode, http.StatusUnauthorized, "Token invalid")

	// ErrPageNotFound - 404: Page not found.
	ErrPageNotFound = NewCode(ErrPageNotFoundCode, http.StatusNotFound, "Page not found")
)

var (
	// ErrDatabase - 500: Database error
	ErrDatabase = NewCode(ErrDatabaseCode, http.StatusInternalServerError, "Database error")
)
