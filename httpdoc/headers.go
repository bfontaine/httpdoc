package httpdoc

import (
	"fmt"
	"strings"
)

// Header is an HTTP header field
type Header struct {
	// Type is either "request", "response", or "both"
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

// GetHeader returns a Header
func (d Doc) GetHeader(name string) (h Header, err error) {
	headers, err := d.parseHeaders()
	if err != nil {
		return
	}

	// normalize the name
	if len(name) > 2 {
		name = strings.Title(strings.ToLower(name))
	} else {
		name = strings.ToUpper(name)
	}

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

// PrettyString returns an human-readable description of the header
func (h Header) PrettyString() string {
	t := h.Type
	if t == "both" {
		t = "request and response"
	}

	return fmt.Sprintf("%s (%s field)\n\n%s\n\nReferences:\n * %s\n",
		h.Text,
		t,
		h.Desc,
		strings.Join(h.Refs, "\n * "),
	)
}
