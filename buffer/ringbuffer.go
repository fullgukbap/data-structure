package buffer

type RingBuffer[T any] struct {
	buf []T

	readPoint  int
	writePoint int
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buf: make([]T, size),
	}
}

func (r *RingBuffer[T]) Write(data []T) {
	copy(r.buf[r.writePoint:], data)
}

func (r *RingBuffer[T]) Readable() {
	
}