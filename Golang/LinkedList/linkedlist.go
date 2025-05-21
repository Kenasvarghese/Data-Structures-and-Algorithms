package linkedlist

type Node struct {
	Value int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Length int
}

func NewLinkedList(values ...int) *LinkedList {
	ll := &LinkedList{
		Head: nil,
		Tail: nil,
		Length: 0,
	}
	for i, value := range values {
		ll.Add(value)
		if i == 0 {
			ll.Tail = ll.Head
		}
	}
	return ll
}

func (l *LinkedList) Add(value int) {
	l.Head = &Node{Value: value, Next: l.Head}
}