package singlelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushBack(1)
	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)

	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushBack(2)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 2, l.Back().Value)

	l.PushBack(4)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 4, l.Back().Value)
}

func TestRoundCount(t *testing.T) {
	var l LinkedList[string]

	l.PushBack("1")
	l.PushBack("2")
	l.PushBack("3")
	l.PushBack("4")
	l.PushBack("5")

	assert.Equal(t, l.RoundCount(), 5)
	assert.Equal(t, l.Count(), 5)
}

func TestGetAt(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	assert.Nil(t, l.GetAt(-1))
	assert.Nil(t, l.GetAt(999999999))

	assert.Equal(t, 1, l.GetAt(0).Value)
	assert.Equal(t, 2, l.GetAt(1).Value)
	assert.Equal(t, 3, l.GetAt(2).Value)

}

func TestInsertAfter(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	lastNode := l.Back()
	l.InsertAfter(lastNode, 100)
	assert.Equal(t, l.Back().Value, 100)

	// before 1, 2, 3, 100
	// after 1, 2, 3, 5, 100

	gotNode := l.GetAt(2)
	l.InsertAfter(gotNode, 5)
	assert.Equal(t, l.GetAt(3).Value, 5)
}

func TestIsInclude(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	firstNode := l.GetAt(0)
	secondNode := l.GetAt(1)
	thirdNode := l.GetAt(2)

	assert.True(t, l.isInclude(firstNode))
	assert.True(t, l.isInclude(secondNode))
	assert.True(t, l.isInclude(thirdNode))

}

func TestFindPrevNode(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(0)
	l.PushBack(1)
	l.PushBack(2)

	assert.Equal(t, l.GetAt(0).Value, l.findPrevNode(l.GetAt(1)).Value)
	assert.Equal(t, l.GetAt(1).Value, l.findPrevNode(l.GetAt(2)).Value)
}

func TestInsertBefore(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(0) // 0
	l.PushBack(1) // 1
	l.PushBack(2) // 2
	l.PushBack(3) // 3
	l.PushBack(4) // 4

	// 0, 1, 2, 3, 4
	l.InsertBefore(l.GetAt(1), 6)
	// 0, 6, 1, 2, 3, 4

	assert.Equal(t, l.GetAt(1).Value, 6)
}

func TestRemoveNode(t *testing.T) {

}

func TestReverse(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.Reverse()

	assert.Equal(t, 4, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 3, l.GetAt(1).Value)
}

func TestSwapReverse(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.SwapReverse()

	assert.Equal(t, 4, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 3, l.GetAt(1).Value)
}
