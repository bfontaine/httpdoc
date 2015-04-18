package httpdoc

import (
	"fmt"
	"strings"
)

// Header is an HTTP header field
type Header struct {
	// Type is either "request" or "response"
	Type string

	// Text is the header’s text
	Text string

	// Desc is an human-readable description of the header
	Desc string

	// Refs are some references for the reader
	Refs []string
}

const (
	headersSource = "headers.yml"
)

func (d Doc) parseHeaders() (m map[string]Header, err error) {
	m = make(map[string]Header)
	err = d.loadSource(headersSource, m)

	return m, err
}

func (d Doc) GetHeader(name string) (h Header, err error) {
	headers, err := d.parseHeaders()
	if err != nil {
		return
	}

	name = strings.Title(strings.ToLower(name))

	h, ok := headers[name]
	if !ok {
		err = ErrUnknownHeader
	}

	h.Text = name

	return
}

func (h Header) String() string {
	return h.Text
}

func (h Header) PrettyString() string {
	return fmt.Sprintf("%s (%s field)\n\n%s\n\nReferences:\n * %s\n",
		h.Text,
		h.Type,
		h.Desc,
		strings.Join(h.Refs, "\n * "),
	)
}
