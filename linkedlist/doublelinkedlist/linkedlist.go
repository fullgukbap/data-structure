package doublelinkedlist

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]

	Value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) PushBack(val T) {
	newNode := &Node[T]{
		Value: val,
	}

	if l.root == nil {
		l.root = newNode
		l.tail = newNode
		l.count = 1
		return
	}

	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode
	l.count++
}

func (l *LinkedList[T]) PushFront(val T) {
	newNode := &Node[T]{
		Value: val,
	}

	if l.root == nil {
		l.root = newNode
		l.tail = newNode
		l.count = 1
		return
	}

	newNode.next = l.root
	l.root.prev = newNode
	l.root = newNode
	l.count++

}

func (l *LinkedList[T]) InsertAfter(node *Node[T], val T) {
	if !l.isInclude(node) {
		return
	}

	newNode := &Node[T]{
		Value: val,
	}

	if l.Back() == node {
		l.PushBack(val)
		return
	}

	newNode.prev = node
	newNode.next = node.next
	node.next.prev = newNode
	node.next = newNode
	l.count++
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], val T) {
	if !l.isInclude(node) {
		return
	}

	if l.Front() == node {
		l.PushFront(val)
		return
	}

	newNode := &Node[T]{
		Value: val,
	}

	newNode.next = node
	newNode.prev = node.prev
	node.prev.next = newNode
	node.prev = newNode
	l.count++
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	i := 0
	for n := l.root; n != nil; n = n.next {
		if i == idx {
			return n
		}
		i++
	}

	return nil
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) isInclude(target *Node[T]) bool {
	for n := l.root; n != nil; n = n.next {
		if target == n {
			return true
		}
	}

	return false
}
func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.root == nil {
		return nil
	}
	n := l.root
	l.root = n.next
	if l.root != nil {
		l.root.prev = nil
	} else {
		l.tail = nil
	}
	n.next = nil
	l.count--
	return n
}

func (l *LinkedList[T]) PopBack() *Node[T] {
	if l.tail == nil {
		return nil
	}
	n := l.tail
	l.tail = n.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.root = nil
	}
	n.prev = nil
	l.count--
	return n
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if node == l.root {
		l.PopFront()
		return
	} else if node == l.tail {
		l.PopBack()
		return
	}

	if !l.isInclude(node) {
		return
	}
	prev := node.prev
	next := node.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}

	node.prev = nil
	node.next = nil
	l.count--
}

func (l *LinkedList[T]) Reverse() {
	if l.root == nil {
		return
	}
	node := l.root
	var next *Node[T]

	for node != nil {
		next = node.next

		node.prev, node.next = node.next, node.prev
		node = next
	}

	l.root, l.tail = l.tail, l.root
}
