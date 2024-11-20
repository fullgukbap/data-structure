package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	s := New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 1, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 3, s.Pop())
}

func TestSlicePush(t *testing.T) {
	s := NewSliceQueue[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 1, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 3, s.Pop())
}

func BenchmarkLinkedListQueue(b *testing.B) {
	linkedListQueue := New[int]()

	for i := 0; i < b.N; i++ {
		linkedListQueue.Push(i)
	}

	for i := 0; i < b.N; i++ {
		linkedListQueue.Pop()
	}
}

func BenchmarkSliceQueue(b *testing.B) {
	sliceQueue := NewSliceQueue[int]()

	for i := 0; i < b.N; i++ {
		sliceQueue.Push(i)
	}

	for i := 0; i < b.N; i++ {
		sliceQueue.Pop()
	}
}
