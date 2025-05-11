package attr_test

import (
	"testing"

	"github.com/teenjuna/goco/attr"
)

func TestString(t *testing.T) {
	a := attr.New("foo", "bar")
	s := attr.MustString(a)
	if s != ` foo="bar"` {
		t.Error("string rendered incorrectly")
	}
}
