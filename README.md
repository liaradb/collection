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

# Constructors

## Iterator

```go
func New[T any](i iter.Seq[T]) Collection[T]
```

## Slice

```go
func NewFromSlice[T any](s []T) Collection[T]
```

## List

```go
func NewFromList[T any](l *list.List) Collection[T]
```

## Channel

```go
func NewFromChan[T any](ctx context.Context, ch <-chan T) Collection[T]
```

# Transformers

Creates a new `Collection` and does not consume the iterator.

## Filter

```go
func (c Collection[T]) Filter(p func(T) bool) Collection[T]
```

## Map

```go
func (c Collection[T]) Map[U any](p func(T) U) Collection[U]
```

# Sinks

Transforms the iterator into a new type. Consumes the iterator.

## Iterator

```go
func (c Collection[T]) Iter() iter.Seq[T]
```

## Slice

```go
func (c Collection[T]) Slice() []T
```

## Reduce

```go
func (c Collection[T]) Reduce[U any](p func(T, U) U) U
```
