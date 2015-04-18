package httpdoc

import (
	"go/build"
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

var DefaultDoc Doc

func (d Doc) loadSource(filename string, target interface{}) (err error) {
	var data []byte

	path := filepath.Join(d.RootDir, filename)

	data, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	return yaml.Unmarshal(data, target)
}

func defaultDocDir() string {
	gopath := build.Default.GOPATH

	return filepath.Join(gopath, "src",
		"github.com", "bfontaine", "httpdoc", "_docs")
}

var (
	statusCodeRe *regexp.Regexp
)

func init() {
	var err error

	// set the default doc
	DefaultDoc = Doc{RootDir: defaultDocDir()}

	// compile regular expressions
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
