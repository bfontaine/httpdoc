package httpdoc

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"

	"github.com/bfontaine/httpdoc/Godeps/_workspace/src/gopkg.in/yaml.v2"
)

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

var (
	statusCodeRe *regexp.Regexp
)

func init() {
	var err error

	statusCodeRe, err = regexp.Compile(`^[12345]\d{2}$`)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}

// GetResourceFor gets a name and tries to find a corresponding resource. It'll
// return ErrUnknownResource if it can’t find it.
func (d Doc) GetResourceFor(name string) (Resource, error) {
	if statusCodeRe.MatchString(name) {
		return d.GetStatusCode(name)
	}

	return InvalidResource, ErrUnknownResource
}
