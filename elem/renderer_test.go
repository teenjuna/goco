package elem_test

import (
	"testing"

	"github.com/teenjuna/goco/elem"
	"github.com/teenjuna/goco/internal"
)

func assert(t *testing.T, expected string, nodes ...elem.Renderer) {
	t.Helper()

	actual := internal.MustString(nodes...)

	if actual != expected {
		t.Fatalf(
			"Element rendered unexpectedly.\n  Actual: %s\nExpected: %s\n",
			actual,
			expected,
		)
	}
}
