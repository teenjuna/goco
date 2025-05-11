package attr_test

import (
	"testing"

	"github.com/teenjuna/goco/attr"
)

func TestEffect(t *testing.T) {
	var (
		cnt = 0
		inc = attr.Effect(func() error {
			cnt += 1
			return nil
		})
	)

	assert(t, ``, inc, inc, inc)

	if cnt != 3 {
		t.Fatalf("cnt is %d, but should be 3", cnt)
	}
}
