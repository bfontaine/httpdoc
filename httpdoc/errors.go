package httpdoc

import "errors"

var (
	// ErrUnknownResource is returned by GetResourceFor if it can’t find a
	// resource
	ErrUnknownResource = errors.New("Unknown resource")

	// ErrUnknownStatusCode is returned by GetStatusCode if it can’t find a
	// status code
	ErrUnknownStatusCode = errors.New("Unknown status code")

	// ErrUnknownHeader is returned by GetHeader if it can’t find an header
	// field
	ErrUnknownHeader = errors.New("Unknown header field")

	// ErrUnknownMethod is returned by GetMethod if it can’t find a method
	ErrUnknownMethod = errors.New("Unknown method")
)
