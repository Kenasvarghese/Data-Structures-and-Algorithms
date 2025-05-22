package linkedlist

import (
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

// NewLinkedList creates a new linked list with the given values
func NewLinkedList(values ...int) *LinkedList {
	ll := &LinkedList{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
	for _, value := range values {
		ll.Append(value)
		if ll.Tail == nil {
			ll.Tail = ll.Head
		}
	}
	return ll
}

// String returns a string representation of the linked list
// thus the Stringer interface is implemented
func (l *LinkedList) String() string {
	values := []string{}
	for node := l.Head; node != nil; node = node.Next {
		values = append(values, strconv.Itoa(node.Value))
	}
	return strings.Join(values, " -> ") + "\nLength: " + strconv.Itoa(l.Length)
}
