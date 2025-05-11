// Everything related to HTML elements.
package elem

import (
	"github.com/teenjuna/goco/internal"
)

// Something that renders as one or more HTML elements.
type Renderer interface {
	internal.Renderer

	// Just a marker method needed to prevent this interface from being implemented in external
	// packages.
	elemMarker()
}
