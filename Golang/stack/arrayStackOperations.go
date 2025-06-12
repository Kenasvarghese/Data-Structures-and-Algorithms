package stack

import "fmt"

// Push adds a new value to the top of the stack
func (s *arrayStack) Push(value int) {
	s.values = append(s.values, value)
	s.length++
}

// Pop removes the top value from the stack
func (s *arrayStack) Pop() int {
	if s.length == 0 {
		return -1
	}
	val := s.values[s.length-1]
	s.values = s.values[:s.length-1]
	s.length--
	return val
}

// Peek returns the top value from the stack without removing it
func (s *arrayStack) Peek() int {
	if s.length == 0 {
		return -1
	}
	return s.values[s.length-1]
}

// Height returns the height of the stack
func (s *arrayStack) Height() int {
	return s.length
}

// Print prints the stack
func (s *arrayStack) Print() {
	fmt.Println("Stack:")
	for i := s.length - 1; i >= 0; i-- {
		fmt.Println(s.values[i])
	}
	fmt.Println("End of Stack")
}
