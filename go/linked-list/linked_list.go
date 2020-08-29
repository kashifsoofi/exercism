// Package linkedlist is a simple library
// that implements a list using doubly linked list.
package linkedlist

import "errors"

// Node to hold data and pointer to next and prev nodes
type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

// Next returns pointer to next node
func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.next
}

// Prev returns pointer to previous node
func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}
	return n.prev
}

// List implementation using linked list
type List struct {
	head *Node
	tail *Node
}

// ErrEmptyList list is empty
var ErrEmptyList = errors.New("empty list")

// NewList creates a new list
func NewList(args ...interface{}) *List {
	l := &List{
		head: nil,
		tail: nil,
	}

	for _, v := range args {
		l.PushBack(v)
	}

	return l
}

// First pointer to the first node (head).
func (l *List) First() *Node {
	return l.head
}

// Last pointer to the last node (tail)
func (l *List) Last() *Node {
	return l.tail
}

// PushBack inserts value at back of list.
func (l *List) PushBack(v interface{}) {
	n := &Node{
		Val:  v,
		prev: l.tail,
	}

	if l.tail != nil {
		l.tail.next = n
	} else {
		l.head = n
	}
	l.tail = n
}

// PopBack removes node from back of list and returns value if successful
func (l *List) PopBack() (interface{}, error) {
	if l.tail == nil {
		return nil, ErrEmptyList
	}

	n := l.tail
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}

	return n.Val, nil
}

// PushFront inserts value at head of list.
func (l *List) PushFront(v interface{}) {
	n := &Node{
		Val:  v,
		next: l.head,
	}
	if l.head != nil {
		l.head.prev = n
	} else {
		l.tail = n
	}
	l.head = n
}

// PopFront removes node from front of list and returns value if successful
func (l *List) PopFront() (interface{}, error) {
	if l.head == nil {
		return nil, ErrEmptyList
	}

	n := l.head
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}

	return n.Val, nil
}

// Reverse reverses list
func (l *List) Reverse() {
	h, t := l.head, l.tail
	for n := l.head; n != nil; n = n.prev {
		n.prev, n.next = n.next, n.prev
	}
	l.head, l.tail = t, h
}
