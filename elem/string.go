package elem

import "github.com/teenjuna/goco/internal"

// Renders element nodes into a string.
func String(nodes ...Renderer) (string, error) {
	return internal.String(nodes...)
}

// Same as [String], but panics on error.
func MustString(nodes ...Renderer) string {
	return internal.MustString(nodes...)
}
