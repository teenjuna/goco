// Everything related to HTML attributes.
package attr

import (
	"github.com/teenjuna/goco/internal"
)

// Something that renders as one or more HTML attributes.
type Renderer interface {
	internal.Renderer

	// Just a marker method needed to prevent this interface from being implemented in external
	// packages.
	attrMarker()
}
