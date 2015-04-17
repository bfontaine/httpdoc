package httpdoc

import "errors"

type Code struct {
	Refs             []string
	Code, Text, Desc string
}

const (
	codesSource = "codes.yml"
)

var (
	ErrUnknownHTTPCode = errors.New("Unknown HTTP code")
)

func (d Doc) parseCodes() (m map[string]Code, err error) {
	m = make(map[string]Code)
	err = d.loadSource(codesSource, m)

	return m, err
}

func (d Doc) GetCode(code string) (c Code, err error) {
	codes, err := d.parseCodes()
	if err != nil {
		return
	}

	c, ok := codes[code]
	if !ok {
		err = ErrUnknownHTTPCode
	}

	c.Code = code

	return
}
