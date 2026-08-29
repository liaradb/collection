package collection

import (
	"fmt"
	"slices"
	"testing"
)

func TestCollection2(t *testing.T) {
	t.Parallel()

	want := []string{
		"(0, string - 0)",
		"(1, string - 1)",
		"(2, string - 2)"}
	result := New2(func(yield func(int, string) bool) {
		for i := range 3 {
			if !yield(i, fmt.Sprintf("%v", i)) {
				return
			}
		}
	}).
		Map(func(i int, s string) (string, string) {
			return fmt.Sprintf("%v", i), fmt.Sprintf("string - %v", s)
		}).
		Map1(func(s1, s2 string) string {
			return fmt.Sprintf("(%v, %v)", s1, s2)
		}).
		Slice()

	if !slices.Equal(result, want) {
		t.Errorf("incorrec result: %v, expected: %v", result, want)
	}
}
