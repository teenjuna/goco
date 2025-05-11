package attr

import (
	"fmt"
	"io"
)

var _ Renderer = Node{}

// General [Renderer] for an attribute. It is used to create spec attributes (like [Class]) and
// should be used to create custom attributes.
type Node struct {
	quote  string
	render func(quote string) func(w io.Writer) error
}

// Creates new [Node].
//
// Value can be either string or bool. Bool attributes have special treatment:
// https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes#boolean_attributes.
func New[V string | bool](name string, value V) Node {
	return Node{
		quote: `"`,
		render: func(quote string) func(w io.Writer) error {
			return func(w io.Writer) error {
				switch value := any(value).(type) {
				// Boolean attributes render only the key if value is true.
				// If value is false, nothing is rendered.
				case bool:
					if value {
						// Render only the key.
						if _, err := w.Write([]byte(" ")); err != nil {
							return err
						}
						if _, err := w.Write([]byte(name)); err != nil {
							return err
						}
					} else {
						// Do nothing.
					}
				// String attributes are rendered as is.
				case string:
					// Render key, equals sign and quoted value.
					if _, err := w.Write([]byte(" ")); err != nil {
						return err
					}
					if _, err := w.Write([]byte(name)); err != nil {
						return err
					}
					if _, err := w.Write([]byte("=")); err != nil {
						return err
					}
					if _, err := w.Write([]byte(quote)); err != nil {
						return err
					}
					if _, err := w.Write([]byte(value)); err != nil {
						return err
					}
					if _, err := w.Write([]byte(quote)); err != nil {
						return err
					}
				default:
					// Sanity check. Should never happen due to generic constraints.
					panic(fmt.Sprintf("invalid type of value: %T", value))
				}

				return nil
			}
		},
	}
}

// Makes [Node] use single quotes for it's value (e.g. foo='bar').
func (n Node) SingleQuoted() Node {
	n.quote = `'`
	return n
}

// Implements [Renderer].
func (n Node) Render(w io.Writer) error {
	if n.render == nil {
		return nil
	}
	return n.render(n.quote)(w)
}

// Implements [Renderer].
func (n Node) attrMarker() {}
