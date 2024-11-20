package linkedlist

import (
	"testing"

	"github.com/fullgukbap/data-structure/linkedlist/doublelinkedlist"
	"github.com/fullgukbap/data-structure/linkedlist/singlelinkedlist"
)

func BenchmarkSingleLinkedList(b *testing.B) {
	l := &singlelinkedlist.LinkedList[int]{}
	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}

	l.Reverse()
}

func BenchmarkDoubleLinkedList(b *testing.B) {
	l := &doublelinkedlist.LinkedList[int]{}
	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}

	l.Reverse()
}
