package elem

import (
	"io"

	"github.com/teenjuna/goco/attr"
)

var _ Renderer = Node{}

// General [Renderer] for an element. It is used to create spec elements (like [Div]) and should be
// used to create custom elements.
type Node struct {
	render func(w io.Writer) error
}

// Creates new [Node].
//
// Panics if children are provided to a void element.
func New(name string, void bool, attrs attr.Renderer, children ...Renderer) Node {
	if void && len(children) > 0 {
		panic("void element can't have children")
	}

	return Node{render: func(w io.Writer) error {
		if _, err := w.Write([]byte("<")); err != nil {
			return err
		}

		if _, err := w.Write([]byte(name)); err != nil {
			return err
		}

		if attrs != nil {
			if err := attrs.Render(w); err != nil {
				return err
			}
		}

		if void {
			if _, err := w.Write([]byte(" />")); err != nil {
				return err
			}
		} else {
			if _, err := w.Write([]byte(">")); err != nil {
				return err
			}

			for _, child := range children {
				if child == nil {
					continue
				}
				if err := child.Render(w); err != nil {
					return err
				}
			}

			if _, err := w.Write([]byte("</")); err != nil {
				return err
			}
			if _, err := w.Write([]byte(name)); err != nil {
				return err
			}
			if _, err := w.Write([]byte(">")); err != nil {
				return err
			}

		}

		return nil
	}}
}

// Implements [Renderer].
func (n Node) Render(w io.Writer) error {
	if n.render == nil {
		return nil
	}
	return n.render(w)
}

// Implements [Renderer].
func (Node) elemMarker() {}
