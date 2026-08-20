package collection

import (
	"container/list"
	"context"
	"iter"
)

type Collection[T any] struct {
	i iter.Seq[T]
}

func New[T any](i iter.Seq[T]) Collection[T] {
	return Collection[T]{i}
}

func NewFromSlice[T any](s []T) Collection[T] {
	return Collection[T]{
		i: func(yield func(T) bool) {
			for _, t := range s {
				if !yield(t) {
					return
				}
			}
		},
	}
}

func NewFromList[T any](l *list.List) Collection[T] {
	return Collection[T]{
		i: func(yield func(T) bool) {
			e := l.Front()
			for e != nil {
				// Should panic if invalid
				t := e.Value.(T)
				if !yield(t) {
					return
				}
				e = e.Next()
			}
		},
	}
}

func NewFromChan[T any](ctx context.Context, ch <-chan T) Collection[T] {
	return Collection[T]{
		i: func(yield func(T) bool) {
			for {
				select {
				case t, ok := <-ch:
					if !ok || !yield(t) {
						return
					}
				case <-ctx.Done():
					return
				}
			}
		},
	}
}

func (c Collection[T]) Iter() iter.Seq[T] { return c.i }

func (c Collection[T]) Slice() []T {
	s := make([]T, 0)
	for t := range c.i {
		s = append(s, t)
	}
	return s
}

func (c Collection[T]) Filter(p func(T) bool) Collection[T] {
	return Collection[T]{
		i: func(yield func(T) bool) {
			for t := range c.i {
				if p(t) && !yield(t) {
					return
				}
			}
		},
	}
}

func (c Collection[T]) Map[U any](p func(T) U) Collection[U] {
	return Collection[U]{
		i: func(yield func(U) bool) {
			for t := range c.i {
				if !yield(p(t)) {
					return
				}
			}
		},
	}
}

func (c Collection[T]) Reduce[U any](p func(T, U) U) U {
	var u U
	for t := range c.i {
		u = p(t, u)
	}
	return u
}
