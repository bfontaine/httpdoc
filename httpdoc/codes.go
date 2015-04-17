package httpdoc

import "errors"

// A StatusCode is an HTTP status code
type StatusCode struct {
	// Code is its actual code (e.g. "200")
	Code string

	// Text is a short text describing the status code (e.g. "OK")
	Text string

	// Desc is an human-readable description of the status code
	Desc string

	// Refs are some references for the reader
	Refs []string
}

const (
	codesSource = "codes.yml"
)

var (
	// ErrUnknownStatusCode is returned by GetStatusCode if it can’t find a
	// status code
	ErrUnknownStatusCode = errors.New("Unknown status code")
)

func (d Doc) parseStatusCodes() (m map[string]StatusCode, err error) {
	m = make(map[string]StatusCode)
	err = d.loadSource(codesSource, m)

	return m, err
}

// GetStatusCode gets a code and returns a new StatusCode struct describing it.
// If it can’t find the code an ErrUnknownStatusCode is returned.
func (d Doc) GetStatusCode(code string) (c StatusCode, err error) {
	codes, err := d.parseStatusCodes()
	if err != nil {
		return
	}

	c, ok := codes[code]
	if !ok {
		err = ErrUnknownStatusCode
	}

	c.Code = code

	return
}
