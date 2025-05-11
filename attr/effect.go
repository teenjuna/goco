package attr

import "io"

var _ Renderer = Effect(nil)

// Special type of [Renderer] that executes a failable function during the rendering.
type Effect func() error

// Implements [Renderer].
func (e Effect) Render(_ io.Writer) error {
	return e()
}

// Implements [Renderer].
func (e Effect) attrMarker() {}
