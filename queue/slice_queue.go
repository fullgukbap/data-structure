package queue

type SliceQueue[T any] struct {
	slice []T
}

func NewSliceQueue[T any]() *SliceQueue[T] {
	return &SliceQueue[T]{}
}

func (q *SliceQueue[T]) Push(val T) {
	q.slice = append(q.slice, val)
}

func (q *SliceQueue[T]) Pop() T {
	var front T
	if len(q.slice) == 0 {
		return front
	}

	front = q.slice[0]
	q.slice = q.slice[1:]
	return front
}
