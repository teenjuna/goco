package attr

import (
	"io"

	"github.com/teenjuna/goco/internal"
)

var _ Renderer = IfBuilder{}

// Builder for conditional rendering (if, else if, else).
type IfBuilder struct {
	builder internal.IfBuilder
}

// Creates an [IfBuilder] and that will render provided content if cond is true.
func If(cond bool, content Renderer) IfBuilder {
	return IfBuilder{builder: internal.If(cond, content)}
}

// Makes [IfBuilder] render provided content if cond is true and previous [IfBuilder] conditions
// are false.
func (i IfBuilder) ElseIf(cond bool, content Renderer) IfBuilder {
	return IfBuilder{builder: i.builder.ElseIf(cond, content)}
}

// Makes [IfBuilder] render provided content previous [IfBuilder] conditions are false.
func (i IfBuilder) Else(content Renderer) Renderer {
	return IfBuilder{builder: i.builder.Else(content)}
}

// Implements [Renderer].
func (i IfBuilder) Render(w io.Writer) error {
	return i.builder.Render(w)
}

// Implements [Renderer].
func (i IfBuilder) attrMarker() {}
