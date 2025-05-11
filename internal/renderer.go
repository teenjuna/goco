package internal

import (
	"io"
)

// Something that renders to one or more HTML nodes.
type Renderer interface {
	Render(w io.Writer) error
}
