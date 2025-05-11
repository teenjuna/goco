package attr_test

import (
	"testing"

	"github.com/teenjuna/goco/attr"
	"github.com/teenjuna/goco/internal"
)

func assert(t *testing.T, expected string, nodes ...attr.Renderer) {
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
