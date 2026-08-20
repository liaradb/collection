package collection

import (
	"fmt"
	"slices"
	"testing"
	"testing/synctest"
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

func TestNewFromChan(t *testing.T) {
	t.Parallel()
	synctest.Test(t, func(t *testing.T) {
		want := []string{"0", "1", "2"}

		ch := make(chan int)
		go func() {
			for i := range 3 {
				ch <- i
			}
			close(ch)
		}()

		result := NewFromChan(t.Context(), ch).
			Map(func(i int) string { return fmt.Sprintf("%v", i) }).
			Slice()

		if !slices.Equal(result, want) {
			t.Errorf("incorrec result: %v, expected: %v", result, want)
		}
	})
}
