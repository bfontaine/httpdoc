package httpdoc

// A Resource is a documented resource
type Resource interface {
	String() string

	// PrettyString is intended as a more verbose `String()` for user display
	PrettyString() string
}

// NilResource is an empty resource
type NilResource struct{}

func (n NilResource) String() string       { return "" }
func (n NilResource) PrettyString() string { return "" }

// InvalidResource is returned by some functions where a Resource is needed but
// no one can be found
var InvalidResource NilResource
