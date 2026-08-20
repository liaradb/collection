package collection

import (
	"fmt"
	"slices"
	"testing"
)

func TestCollection(t *testing.T) {
	t.Parallel()

	want := []string{"0", "1", "2"}
	result := New(func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	}).
		Map(func(i int) string { return fmt.Sprintf("%v", i) }).
		Slice()

	if !slices.Equal(result, want) {
		t.Errorf("incorrec result: %v, expected: %v", result, want)
	}
}
