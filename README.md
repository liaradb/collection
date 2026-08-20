# collection

This is a simple wrapper around iterators to provide method chaining.

```go
c := New(func(yield func(int) bool) {
	for i := range 10 {
		if !yield(i) {
			return
		}
	}
})

result := c.
	Filter(func(i int) bool { return i%2 == 0 }).
	Map(func(i int) string { return fmt.Sprintf("%v", i) }).
	Slice()

```
