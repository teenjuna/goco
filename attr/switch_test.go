package attr_test

import (
	"testing"

	"github.com/teenjuna/goco/attr"
)

func TestSwitch(t *testing.T) {
	assert(t,
		``,
		attr.Switch(1).
			Case(2, attr.New("foo", true)).
			Case(3, attr.New("bar", true)),
	)
	assert(t,
		` foo`,
		attr.Switch(1).
			Case(1, attr.New("foo", true)).
			Case(3, attr.New("bar", true)).
			Default(attr.New("baz", true)),
	)
	assert(t,
		` bar`,
		attr.Switch(1).
			Case(2, attr.New("foo", true)).
			Case(1, attr.New("bar", true)).
			Default(attr.New("baz", true)),
	)
	assert(t,
		` baz`,
		attr.Switch(1).
			Case(2, attr.New("foo", true)).
			Case(3, attr.New("bar", true)).
			Default(attr.New("baz", true)),
	)
}
