package httpdoc

import "errors"

var (
	// ErrUnknownResource is returned by GetResourceFor if it can’t find a
	// resource
	ErrUnknownResource = errors.New("Unknown resource")

	// ErrUnknownStatusCode is returned by GetStatusCode if it can’t find a
	// status code
	ErrUnknownStatusCode = errors.New("Unknown status code")
)
