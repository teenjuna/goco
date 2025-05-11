package elem_test

import (
	"slices"
	"testing"

	"github.com/teenjuna/goco/elem"
)

func TestSeq(t *testing.T) {
	assert(t,
		`<foo /><bar /><baz />`,
		elem.Seq(slices.Values(elem.Slice{
			elem.New("foo", true, nil),
			elem.New("bar", true, nil),
			elem.New("baz", true, nil),
		})),
	)
}
