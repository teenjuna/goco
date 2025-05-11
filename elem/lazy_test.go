package elem_test

import (
	"io"
	"testing"

	"github.com/teenjuna/goco/elem"
)

func TestLazy(t *testing.T) {
	var flag bool

	a := elem.Lazy(func() elem.Renderer {
		flag = true
		return elem.Div(nil)
	})

	if flag {
		t.Error("flag is set prematurely")
	}

	_ = a.Render(io.Discard)

	if !flag {
		t.Error("flag is not set")
	}
}
