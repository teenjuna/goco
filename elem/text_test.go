package elem_test

import (
	"testing"

	"github.com/teenjuna/goco/elem"
)

func TestText(t *testing.T) {
	assert(t, ``, elem.Text(""))
	assert(t, `foo`, elem.Text("foo"))
	assert(t, `&lt;bar/&gt;`, elem.Text("<bar/>"))
	assert(t, ``, elem.Textf(""))
	assert(t, `foo`, elem.Textf("%s", "foo"))
	assert(t, `&lt;bar/&gt;`, elem.Textf("<%s/>", "bar"))
}

func TestRaw(t *testing.T) {
	assert(t, ``, elem.RawText(""))
	assert(t, `foo`, elem.RawText("foo"))
	assert(t, `<bar />`, elem.RawText("<bar />"))
	assert(t, ``, elem.RawTextf(""))
	assert(t, `foo`, elem.RawTextf("%s", "foo"))
	assert(t, `<bar/>`, elem.RawTextf("<%s/>", "bar"))
}
