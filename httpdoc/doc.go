package httpdoc

import (
	"io/ioutil"
	"path/filepath"

	"github.com/bfontaine/httpdoc/Godeps/_workspace/src/gopkg.in/yaml.v2"
)

// A Resource is a documented resource
type Resource interface {
	String() string

	// PrettyString is intended as a more verbose `String()` for user display
	PrettyString() string
}

// Doc represents a documentation source
type Doc struct {
	// the root dir in which are the documentation files
	RootDir string
}

func (d Doc) loadSource(filename string, target interface{}) (err error) {
	var data []byte

	path := filepath.Join(d.RootDir, filename)

	data, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	return yaml.Unmarshal(data, target)
}
