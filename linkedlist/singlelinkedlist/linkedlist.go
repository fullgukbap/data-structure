package singlelinkedlist

type Node[T any] struct {
	next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root *Node[T]
	tail *Node[T]

	count int
}

func (l *LinkedList[T]) PushBack(value T) {
	// 카운트 증가
	l.count++

	// 새로운 노드 생성
	newNode := &Node[T]{
		next:  nil,
		Value: value,
	}

	// 비어있는 링크드 리스트
	if l.root == nil {
		l.root = newNode
		l.tail = newNode
		return
	}

	// 기존에 있는 링크드 리스트
	l.tail.next = newNode
	l.tail = newNode
}

func (l *LinkedList[T]) PushFront(value T) {
	// 카운트 증가
	l.count++

	// 새로운 노드 생성
	newNode := &Node[T]{
		next:  nil,
		Value: value,
	}

	// 만약 빈 링크드 리스트
	if l.root == nil {
		l.root = newNode
		l.tail = newNode
		return
	}

	// 기존에 있는 링크드 리스트 처리
	newNode.next = l.root
	l.root = newNode
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) RoundCount() int {
	cnt := 0

	for n := l.root; n != nil; n = n.next {
		cnt++
	}

	return cnt
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if l.Count() < idx || idx < 0 {
		return nil
	}

	var gotNode *Node[T]
	i := 0
	for n := l.root; n != nil; n = n.next {
		if i == idx {
			gotNode = n
		}
		i++
	}

	return gotNode
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {

	// 전달된 노드가 l에 존재하는 노드인가?
	if !l.isInclude(node) {
		return
	}

	newNode := &Node[T]{
		Value: value,
	}

	// 만약 tail에 삽입한다면?
	if l.tail == node {
		l.PushBack(value)
		return
	}

	newNode.next, node.next = node.next, newNode
	l.count++
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {
	if !l.isInclude(node) {
		return
	}

	newNode := &Node[T]{
		Value: value,
	}

	// 만약 root 노드에 추가하고 싶다면?
	if l.root == node {
		l.PushFront(value)
		return
	}

	prevNode := l.findPrevNode(node)

	newNode.next = prevNode.next
	prevNode.next = newNode
	l.count++
}

func (l LinkedList[T]) isInclude(node *Node[T]) bool {
	for n := l.root; n != nil; n = n.next {
		if n == node {
			return true
		}
	}

	return false
}

func (l LinkedList[T]) findPrevNode(currentNode *Node[T]) *Node[T] {
	for n := l.root; n != nil; n = n.next {
		if n.next == currentNode {
			return n
		}
	}

	return nil
}

func (l *LinkedList[T]) PopFront() {
	if l.root == nil {
		return
	}

	l.root.next, l.root = nil, l.root.next
	if l.root == nil {
		l.tail = nil
	}
	l.count--
}

func (l *LinkedList[T]) RemoveNode(target *Node[T]) {
	if target == l.root {
		l.PopFront()
		return
	}

	if !l.isInclude(target) {
		return
	}

	prevNode := l.findPrevNode(target)
	prevNode.next = target.next
	target.next = nil
	l.count--
}
