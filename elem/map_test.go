package elem_test

import (
	"slices"
	"testing"

	"github.com/teenjuna/goco/elem"
)

func TestMap(t *testing.T) {
	assert(t,
		`<foo /><bar />`,
		elem.Map(
			slices.Values([]string{"foo", "bar"}),
			func(item string) elem.Renderer {
				return elem.New(item, true, nil)
			},
		),
	)
}
