package linkedlist

import "errors"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

type List struct {
	head *Node // The list first node
	tail *Node // The list last node
}

func NewList(args ...interface{}) *List {
	newList := new(List)

	for _, arg := range args {
		newList.Push(arg)
	}

	return newList
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	newNode := &Node{Value: v, prev: nil, next: nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode.next = l.head
	l.head.prev = newNode
	l.head = newNode
}

func (l *List) Push(v interface{}) {
	newNode := &Node{Value: v, prev: l.tail, next: nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	l.tail = newNode
}

func (l *List) Shift() (interface{}, error) {
	if l.head == nil || l.tail == nil {
		return nil, errors.New("Empty list!")
	}

	oldHead := l.head
	newHead := oldHead.Next()
	if newHead != nil {
		newHead.prev = nil
		l.head = newHead
	} else {
		l.head = nil
		l.tail = nil
	}

	return oldHead.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.tail == nil || l.head == nil {
		return nil, errors.New("Empty list!")
	}

	oldTail := l.tail
	newTail := oldTail.Prev()
	if newTail != nil {
		newTail.next = nil
		l.tail = newTail
	} else {
		l.head = nil
		l.tail = nil
	}

	return oldTail.Value, nil
}

func (l *List) Reverse() {
	if l.head == nil {
		return
	}

	newList := new(List)
	node := l.tail

	for {
		newList.Push(node.Value)

		node = node.prev
		if node == nil {
			break
		}
	}

	*l = *newList
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
