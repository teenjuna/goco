package attr_test

import (
	"testing"

	"github.com/teenjuna/goco/attr"
)

func TestSlice(t *testing.T) {
	assert(t,
		` foo="bar" bar="baz" baz`,
		attr.Slice{
			attr.New("foo", "bar"),
			attr.New("bar", "baz"),
			attr.New("baz", true),
			attr.New("qux", false),
		},
	)
}
