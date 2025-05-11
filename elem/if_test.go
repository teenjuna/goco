package elem_test

import (
	"testing"

	"github.com/teenjuna/goco/elem"
)

func TestIf(t *testing.T) {
	assert(t,
		``,
		elem.
			If(false, elem.New("foo", true, nil)).
			ElseIf(false, elem.New("bar", true, nil)),
	)
	assert(t,
		`<foo />`,
		elem.
			If(true, elem.New("foo", true, nil)).
			ElseIf(true, elem.New("bar", true, nil)).
			Else(elem.New("baz", true, nil)),
	)
	assert(t,
		`<bar />`,
		elem.
			If(false, elem.New("foo", true, nil)).
			ElseIf(true, elem.New("bar", true, nil)).
			Else(elem.New("baz", true, nil)),
	)
	assert(t,
		`<baz />`,
		elem.
			If(false, elem.New("foo", true, nil)).
			ElseIf(false, elem.New("bar", true, nil)).
			Else(elem.New("baz", true, nil)),
	)
}
