package attr

import "io"

var _ Renderer = Slice{}

// Slice of [Renderer] items.
type Slice []Renderer

// Implements [Renderer].
func (s Slice) Render(w io.Writer) error {
	for _, node := range s {
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
func (s Slice) attrMarker() {}
