package attr_test

import (
	"testing"

	"github.com/teenjuna/goco/attr"
)

func TestZero(t *testing.T) {
	assert(t, ``, attr.Node{})
}

func TestNew(t *testing.T) {
	assert(t, ` foo="bar"`, attr.New("foo", "bar"))
	assert(t, ` foo='bar'`, attr.New("foo", "bar").SingleQuoted())
	assert(t, ` foo`, attr.New("foo", true))
	assert(t, ``, attr.New("foo", false))
}
