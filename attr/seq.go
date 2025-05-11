package attr

import (
	"io"
	"iter"
)

var _ Renderer = Seq(nil)

// Sequence of [Renderer] items.
type Seq iter.Seq[Renderer]

// Implements [Renderer].
func (s Seq) Render(w io.Writer) error {
	for node := range s {
		if node == nil {
			continue
		}
		if err := node.Render(w); err != nil {
			return err
		}
	}
	return nil
}

// Implements [Renderer].
func (s Seq) attrMarker() {}
