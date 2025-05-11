package attr_test

import (
	"testing"

	"github.com/teenjuna/goco/attr"
)

func TestIf(t *testing.T) {
	assert(t,
		``,
		attr.
			If(false, attr.New("foo", true)).
			ElseIf(false, attr.New("bar", true)),
	)
	assert(t,
		` foo`,
		attr.
			If(true, attr.New("foo", true)).
			ElseIf(true, attr.New("bar", true)).
			Else(attr.New("baz", true)),
	)
	assert(t,
		` bar`,
		attr.
			If(false, attr.New("foo", true)).
			ElseIf(true, attr.New("bar", true)).
			Else(attr.New("baz", true)),
	)
	assert(t,
		` baz`,
		attr.
			If(false, attr.New("foo", true)).
			ElseIf(false, attr.New("bar", true)).
			Else(attr.New("baz", true)),
	)
}
