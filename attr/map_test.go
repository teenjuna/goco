package attr_test

import (
	"slices"
	"testing"

	"github.com/teenjuna/goco/attr"
)

func TestMap(t *testing.T) {
	assert(t,
		` foo bar`,
		attr.Map(
			slices.Values([]string{"foo", "bar"}),
			func(item string) attr.Renderer {
				return attr.New(item, true)
			},
		),
	)
}
