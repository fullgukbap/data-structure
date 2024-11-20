package stack

type SliceStack[T any] struct {
	slice []T
}

func NewSliceStack[T any]() *SliceStack[T] {
	return &SliceStack[T]{
		slice: make([]T, 10, 20),
	}
}

func (s *SliceStack[T]) Push(val T) {
	s.slice = append(s.slice, val)
}

func (s *SliceStack[T]) Pop() T {
	var last T
	if len(s.slice) == 0 {
		return last
	}

	last = s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return last
}
