package attr_test

import (
	"io"
	"testing"

	"github.com/teenjuna/goco/attr"
)

func TestLazy(t *testing.T) {
	var flag bool

	a := attr.Lazy(func() attr.Renderer {
		flag = true
		return attr.ID("foo")
	})

	if flag {
		t.Error("flag is set prematurely")
	}

	_ = a.Render(io.Discard)

	if !flag {
		t.Error("flag is not set")
	}
}
