package stack

import "fmt"

// Push adds a new value to the top of the stack
func (s *stack) Push(value int) {
	newNode := &node{
		Value: value,
		Next:  s.Top,
	}
	s.Top = newNode
	s.Height++
}

// Pop removes the top value from the stack
func (s *stack) Pop() int {
	if s.Top == nil {
		return -1
	}
	val := s.Top.Value
	s.Top = s.Top.Next
	s.Height--
	return val
}

// Peek returns the top value from the stack without removing it
func (s *stack) Peek() int {
	if s.Top == nil {
		return -1
	}
	return s.Top.Value
}

// Print prints the stack
func (s *stack) Print() {
	temp := s.Top
	fmt.Println("Stack:")
	for temp != nil {
		fmt.Println(temp.Value)
		temp = temp.Next
	}
	fmt.Println("End of Stack")
}
