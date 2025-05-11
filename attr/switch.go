package attr

import (
	"io"

	"github.com/teenjuna/goco/internal"
)

var _ Renderer = SwitchBuilder[int]{}

// Builder for conditional rendering (switch, case, default).
type SwitchBuilder[T comparable] struct {
	builder internal.SwitchBuilder[T]
}

// Creates a [SwitchBuilder] for provided value.
func Switch[T comparable](value T) SwitchBuilder[T] {
	return SwitchBuilder[T]{builder: internal.Switch(value)}
}

// Makes [SwitchBuilder] render provided content if provided value matches [SwitchBuilder]'s value
// and there were no previous matches.
func (s SwitchBuilder[T]) Case(value T, content Renderer) SwitchBuilder[T] {
	return SwitchBuilder[T]{builder: s.builder.Case(value, content)}
}

// Makes [SwitchBuilder] render provided content if there were no previous matches.
func (s SwitchBuilder[T]) Default(content Renderer) Renderer {
	return SwitchBuilder[T]{builder: s.builder.Default(content)}
}

// Implements [Renderer].
func (s SwitchBuilder[T]) Render(w io.Writer) error {
	return s.builder.Render(w)
}

// Implements [Renderer].
func (s SwitchBuilder[T]) attrMarker() {}
