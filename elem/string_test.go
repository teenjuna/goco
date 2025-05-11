package elem_test

import (
	"testing"

	"github.com/teenjuna/goco/elem"
)

func TestString(t *testing.T) {
	a := elem.New("foo", true, nil)
	s := elem.MustString(a)
	if s != `<foo />` {
		t.Error("string rendered incorrectly")
	}
}
