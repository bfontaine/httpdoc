package httpdoc

import "errors"

var (
	// ErrUnknownStatusCode is returned by GetStatusCode if it can’t find a
	// status code
	ErrUnknownStatusCode = errors.New("Unknown status code")

	ErrUnknownResource = errors.New("Unknown resource")
)
