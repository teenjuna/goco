package internal

import "io"

type SwitchBuilder[T comparable] struct {
	value  T
	result Renderer
}

func (s SwitchBuilder[T]) Render(w io.Writer) error {
	if s.result == nil {
		return nil
	}
	return s.result.Render(w)
}

func (s SwitchBuilder[T]) attrMarker() {}

func Switch[T comparable](value T) SwitchBuilder[T] {
	return SwitchBuilder[T]{value: value}
}

func (s SwitchBuilder[T]) Case(value T, content Renderer) SwitchBuilder[T] {
	if value != s.value || s.result != nil {
		return s
	}
	return SwitchBuilder[T]{value: s.value, result: content}
}

func (s SwitchBuilder[T]) Default(content Renderer) SwitchBuilder[T] {
	if s.result != nil {
		return s
	}
	return SwitchBuilder[T]{value: s.value, result: content}
}
