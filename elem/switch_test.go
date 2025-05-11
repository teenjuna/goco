package elem_test

import (
	"testing"

	"github.com/teenjuna/goco/elem"
)

func TestSwitch(t *testing.T) {
	assert(t,
		``,
		elem.Switch(1).
			Case(2, elem.New("foo", true, nil)).
			Case(3, elem.New("bar", true, nil)),
	)
	assert(t,
		`<foo />`,
		elem.Switch(1).
			Case(1, elem.New("foo", true, nil)).
			Case(3, elem.New("bar", true, nil)).
			Default(elem.New("baz", true, nil)),
	)
	assert(t,
		`<bar />`,
		elem.Switch(1).
			Case(2, elem.New("foo", true, nil)).
			Case(1, elem.New("bar", true, nil)).
			Default(elem.New("baz", true, nil)),
	)
	assert(t,
		`<baz />`,
		elem.Switch(1).
			Case(2, elem.New("foo", true, nil)).
			Case(3, elem.New("bar", true, nil)).
			Default(elem.New("baz", true, nil)),
	)
}
