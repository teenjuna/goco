package elem

import (
	"fmt"
	"html"
	"io"
)

// NOTE: maybe it should be in goco/text:
// - text.Escaped(...)
// - text.Escapedf(...)
// - text.Unescaped(...)
// - text.Unescapedf(...)
// But will it be able to be a child of an element?

func Text(str string) Node {
	return Node{render: func(w io.Writer) error {
		escaped := html.EscapeString(str)
		_, err := w.Write([]byte(escaped))
		return err
	}}
}

func Textf(format string, args ...any) Node {
	return Node{render: func(w io.Writer) error {
		formatted := fmt.Sprintf(format, args...)
		escaped := html.EscapeString(formatted)
		_, err := w.Write([]byte(escaped))
		return err
	}}
}

func RawText(str string) Node {
	return Node{render: func(w io.Writer) error {
		_, err := w.Write([]byte(str))
		return err
	}}
}

func RawTextf(format string, args ...any) Node {
	return Node{render: func(w io.Writer) error {
		_, err := w.Write(fmt.Appendf(nil, format, args...))
		return err
	}}
}
