package linkedlist

// Append adds a new node to the end of the list
func (l *LinkedList) Append(value int) {
	if l.Tail == nil {
		l.Prepend(value)
		l.Tail = l.Head
		return
	}
	l.Tail.Next = &Node{Value: value}
	l.Tail = l.Tail.Next
	l.Length++
}

// Prepend adds a new node to the beginning of the list
func (l *LinkedList) Prepend(value int) {
	l.Head = &Node{Value: value, Next: l.Head}
	if l.Tail == nil {
		l.Tail = l.Head
	}
	l.Length++
}

// Pop removes the last node from the list and returns its value
func (l *LinkedList) Pop() int {
	if l.Length == 0 {
		return -1
	}
	node := l.Head
	value := l.Tail.Value
	for node != l.Tail && node.Next != l.Tail {
		node = node.Next
	}
	node.Next = nil
	l.Tail = node
	l.Length--
	if l.Length == 0 {
		l.Head = nil
		l.Tail = nil
	}
	return value
}

// PopFirst removes the first node from the list and returns its value
func (l *LinkedList) PopFirst() int {
	if l.Length == 0 {
		return -1
	}
	value := l.Head.Value
	l.Head = l.Head.Next
	l.Length--
	if l.Length == 0 {
		l.Tail = nil
	}
	return value
}

// Get returns the node at the given index
func (l *LinkedList) Get(index int) *Node {
	if index < 0 || index >= l.Length {
		return nil
	}
	node := l.Head
	for range index {
		node = node.Next
	}
	return node
}

// Set updates the value of the node at the given index
func (l *LinkedList) Set(index int, value int) {
	node := l.Get(index)
	if node == nil {
		return
	}
	node.Value = value
}

// Insert adds a new node with the given value at the given index
func (l *LinkedList) Insert(index int, value int) {
	if index == 0 {
		l.Prepend(value)
		return
	}
	if index == l.Length {
		l.Append(value)
		return
	}
	node := l.Get(index)
	if node == nil {
		return
	}
	*node = Node{Value: value, Next: &Node{Value: node.Value, Next: node.Next}}
	l.Length++
}

// Remove removes the node at the given index
func (l *LinkedList) Remove(index int) int {
	value := -1
	if index == 0 {
		value = l.PopFirst()
		return value
	}
	if index == l.Length-1 {
		value = l.Pop()
		return value
	}
	node := l.Get(index)
	if node == nil {
		return -1
	}
	value = node.Value
	*node = *node.Next
	l.Length--
	return value
}

// Reverse reverses the list
func (l *LinkedList) Reverse() {
	prev := (*Node)(nil)
	current := l.Head
	l.Tail = l.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
