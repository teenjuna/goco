package attr_test

import (
	"testing"

	"github.com/teenjuna/goco/attr"
)

func TestClass(t *testing.T) {
	assert(t,
		` class="foo bar qux"`,
		attr.Class("foo", "bar").With("baz", false).With("qux", true),
	)
}
