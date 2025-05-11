package attr

import (
	"iter"

	"github.com/teenjuna/goco/internal"
)

// Maps a generic sequence of items to a sequence of [Renderer] items.
func Map[T any](items iter.Seq[T], transform func(item T) Renderer) Seq {
	return Seq(internal.Map(items, transform))
}
