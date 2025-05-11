package elem_test

import (
	"testing"

	"github.com/teenjuna/goco/attr"
	"github.com/teenjuna/goco/elem"
)

func TestNew(t *testing.T) {
	assert(t,
		``,
		elem.Node{},
	)

	assert(t,
		`<foo />`,
		elem.New("foo", true, nil),
	)
	assert(t,
		`<foo bar="baz" />`,
		elem.New("foo", true, attr.New("bar", "baz")),
	)
	assert(t,
		`<foo bar="baz" />`,
		elem.New("foo", true, attr.New("bar", "baz")),
	)
	assert(t,
		`<foo bar="baz" qux />`,
		elem.New("foo", true, attr.Slice{attr.New("bar", "baz"), attr.New("qux", true)}),
	)
	assert(t,
		`<foo></foo>`,
		elem.New("foo", false, nil),
	)
	assert(t,
		`<foo bar="baz"></foo>`,
		elem.New("foo", false, attr.New("bar", "baz")),
	)
	assert(t,
		`<foo bar="baz"></foo>`,
		elem.New("foo", false, attr.New("bar", "baz")),
	)
	assert(t,
		`<foo bar="baz" qux></foo>`,
		elem.New("foo", false, attr.Slice{attr.New("bar", "baz"), attr.New("qux", true)}),
	)
	assert(t,
		`<foo>qux</foo>`,
		elem.New("foo", false, nil, elem.Text("qux")),
	)
	assert(t,
		`<foo bar="baz">qux</foo>`,
		elem.New("foo", false, attr.New("bar", "baz"), elem.Text("qux")),
	)
}
