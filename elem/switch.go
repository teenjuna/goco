package elem

import (
	"io"

	"github.com/teenjuna/goco/internal"
)

var _ Renderer = SwitchBuilder[int]{}

type SwitchBuilder[T comparable] struct {
	builder internal.SwitchBuilder[T]
}

func (s SwitchBuilder[T]) Render(w io.Writer) error {
	return s.builder.Render(w)
}

func (s SwitchBuilder[T]) elemMarker() {}

func Switch[T comparable](value T) SwitchBuilder[T] {
	return SwitchBuilder[T]{builder: internal.Switch(value)}
}

func (s SwitchBuilder[T]) Case(value T, renderer Renderer) SwitchBuilder[T] {
	return SwitchBuilder[T]{builder: s.builder.Case(value, renderer)}
}

func (s SwitchBuilder[T]) Default(renderer Renderer) Renderer {
	return SwitchBuilder[T]{builder: s.builder.Default(renderer)}
}
