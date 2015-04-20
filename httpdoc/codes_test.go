package httpdoc

import (
	"testing"

	"github.com/bfontaine/httpdoc/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestValidCode(t *testing.T) {
	s, err := DefaultDoc.GetStatusCode("200")

	assert.Nil(t, err)
	assert.Equal(t, s.Code, "200")
	assert.Equal(t, s.Text, "OK")
}

func TestInValidCode(t *testing.T) {
	_, err := DefaultDoc.GetStatusCode("123")

	assert.Equal(t, err, ErrUnknownStatusCode)
}
