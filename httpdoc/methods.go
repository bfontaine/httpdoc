package httpdoc

import (
	"fmt"
	"strings"
)

// Method is an HTTP method
type Method struct {
	// Verb is the method’s verb
	Verb string

	// Desc is an human-readable description of the method
	Desc string

	// Refs are some references for the reader
	Refs []string
}

const (
	methodsSource = "methods.yml"
)

func (d Doc) parseMethods() (m map[string]Method, err error) {
	m = make(map[string]Method)
	err = d.loadSource(methodsSource, m)

	return m, err
}

// GetMethod returns a Method from its verb
func (d Doc) GetMethod(name string) (m Method, err error) {
	methods, err := d.parseMethods()
	if err != nil {
		return
	}

	name = strings.ToUpper(name)

	m, ok := methods[name]
	if !ok {
		err = ErrUnknownMethod
	}

	m.Verb = name

	return
}

func (m Method) String() string {
	return m.Verb
}

// PrettyString returns an human-readable description of this method
func (m Method) PrettyString() string {
	return fmt.Sprintf("%s\n\n%s\n\nReferences:\n * %s\n",
		m.Verb,
		m.Desc,
		strings.Join(m.Refs, "\n * "),
	)
}
