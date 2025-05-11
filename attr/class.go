package attr

import (
	"io"
	"strings"
)

var _ Renderer = ClassBuilder{}

// Builder for `class` attribute.
// Allows rendering classes conditionally using [ClassBuilder.With] method.
type ClassBuilder struct {
	classes []string
}

// Creates a [ClassBuilder] with provided classes.
func Class(classes ...string) ClassBuilder {
	return ClassBuilder{classes: classes}
}

// Adds class if cond is true.
func (cb ClassBuilder) With(class string, cond bool) ClassBuilder {
	if !cond {
		return cb
	}

	if cb.classes == nil {
		cb.classes = make([]string, 0, 1)
	}

	cb.classes = append(cb.classes, class)

	return cb
}

// Implements [Renderer].
func (c ClassBuilder) Render(w io.Writer) error {
	if c.classes == nil || len(c.classes) == 0 {
		return nil
	}

	combined := strings.Join(c.classes, " ")

	return New("class", combined).Render(w)
}

// Implements [Renderer].
func (c ClassBuilder) attrMarker() {}
