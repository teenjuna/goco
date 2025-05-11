package attr

import "io"

var _ Renderer = Lazy(nil)

// Special type of [Renderer] that constructs another [Renderer] lazily, only during rendering.
//
// It's intended to be used in combination with conditional renderers like [If] or [Switch] to
// prevent execution of resource-intensive (or whatever) work if it's not needed.
type Lazy func() Renderer

// Implements [Renderer].
func (l Lazy) Render(w io.Writer) error {
	return l().Render(w)
}

// Implements [Renderer].
func (l Lazy) attrMarker() {}
