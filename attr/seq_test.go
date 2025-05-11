package attr_test

import (
	"slices"
	"testing"

	"github.com/teenjuna/goco/attr"
)

func TestSeq(t *testing.T) {
	assert(t,
		` foo="bar" bar="baz" baz`,
		attr.Seq(slices.Values(attr.Slice{
			attr.New("foo", "bar"),
			attr.New("bar", "baz"),
			attr.New("baz", true),
			attr.New("qux", false),
		})),
	)
}
