package collection

import "iter"

type Collection2[T any, U any] struct {
	i iter.Seq2[T, U]
}

func New2[T any, U any](i iter.Seq2[T, U]) Collection2[T, U] {
	return Collection2[T, U]{i}
}

func (c Collection2[T, U]) Iter() iter.Seq2[T, U] { return c.i }

func (c Collection2[T, U]) Filter(p func(T, U) bool) Collection2[T, U] {
	return Collection2[T, U]{
		i: func(yield func(T, U) bool) {
			for t, u := range c.i {
				if p(t, u) && !yield(t, u) {
					return
				}
			}
		},
	}
}

func (c Collection2[T, U]) Map1[V any](p func(T, U) V) Collection[V] {
	return Collection[V]{
		i: func(yield func(V) bool) {
			for t, u := range c.i {
				if !yield(p(t, u)) {
					return
				}
			}
		},
	}
}

func (c Collection2[T, U]) Map[V any, W any](p func(T, U) (V, W)) Collection2[V, W] {
	return Collection2[V, W]{
		i: func(yield func(V, W) bool) {
			for t, u := range c.i {
				if !yield(p(t, u)) {
					return
				}
			}
		},
	}
}

func (c Collection2[T, U]) MapFirst() Collection[T] {
	return Collection[T]{
		i: func(yield func(T) bool) {
			for t := range c.i {
				if !yield(t) {
					return
				}
			}
		},
	}
}

func (c Collection2[T, U]) MapSecond() Collection[U] {
	return Collection[U]{
		i: func(yield func(U) bool) {
			for _, u := range c.i {
				if !yield(u) {
					return
				}
			}
		},
	}
}

func (c Collection2[T, U]) Reduce[V any](p func(T, U, V) V) V {
	var v V
	for t, u := range c.i {
		v = p(t, u, v)
	}
	return v
}
