package stack

import "fmt"

// Push adds a new value to the top of the stack
func (s *stack) Push(value int) {
	newNode := &node{
		value: value,
		next:  s.top,
	}
	s.top = newNode
	s.height++
}

// Pop removes the top value from the stack
func (s *stack) Pop() int {
	if s.top == nil {
		return -1
	}
	val := s.top.value
	s.top = s.top.next
	s.height--
	return val
}

// Peek returns the top value from the stack without removing it
func (s *stack) Peek() int {
	if s.top == nil {
		return -1
	}
	return s.top.value
}

// Height returns the height of the stack
func (s *stack) Height() int {
	return s.height
}

// Print prints the stack
func (s *stack) Print() {
	temp := s.top
	fmt.Println("Stack:")
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
	fmt.Println("End of Stack")
}
