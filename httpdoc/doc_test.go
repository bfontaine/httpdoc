package httpdoc

import (
	"testing"

	"github.com/bfontaine/httpdoc/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/bfontaine/httpdoc/Godeps/_workspace/src/github.com/stretchr/testify/require"
)

func TestDefaultDocRootDir(t *testing.T) {
	assert.NotEqual(t, DefaultDoc.RootDir, "")
}

func TestDocRecognizeStatus(t *testing.T) {
	s, err := DefaultDoc.GetResourceFor("304")

	require.Nil(t, err)

	switch s.(type) {
	case StatusCode:
		// success
	default:
		t.Fatal("GetResourceFor(\"304\") should return a StatusCode")
	}

	assert.Equal(t, "304", s.(StatusCode).Code)
}

func TestDocRecognizeMethod(t *testing.T) {
	s, err := DefaultDoc.GetResourceFor("options")

	require.Nil(t, err)

	switch s.(type) {
	case Method:
		// success
	default:
		t.Fatal("GetResourceFor(\"options\") should return a Method")
	}

	assert.Equal(t, "OPTIONS", s.(Method).Verb)
}

func TestDocRecognizeHeader(t *testing.T) {
	s, err := DefaultDoc.GetResourceFor("te")

	require.Nil(t, err)

	switch s.(type) {
	case Header:
		// success
	default:
		t.Fatal("GetResourceFor(\"te\") should return a Header")
	}

	assert.Equal(t, "TE", s.(Header).Text)
}
