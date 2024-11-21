package buffer

import (
	"golang.org/x/exp/constraints"
)

type SliceBuffer[T any] struct {
	buf []T
}

func NewSliceBuffer[T any]() *SliceBuffer[T] {
	return &SliceBuffer[T]{}
}

func (s *SliceBuffer[T]) Write(data []T) {
	s.buf = append(s.buf, data...)
}

func (s *SliceBuffer[T]) Readable() int {
	return len(s.buf)
}

func min[T constraints.Ordered](a, b T) T {
	if a > b {
		return b
	}

	return a
}

func (s *SliceBuffer[T]) Read(count int) []T {
	readCnt := min(count, s.Readable())
	dest := make([]T, readCnt)

	copy(dest, s.buf[:readCnt])
	s.buf = s.buf[readCnt:]
	return dest
}
