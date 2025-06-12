package stack

import "fmt"

// Push adds a new value to the top of the stack
func (s *arrayStack) Push(value int) {
	s.Values = append(s.Values, value)
	s.Length++
}

// Pop removes the top value from the stack
func (s *arrayStack) Pop() int {
	if s.Length == 0 {
		return -1
	}
	val := s.Values[s.Length-1]
	s.Values = s.Values[:s.Length-1]
	s.Length--
	return val
}

// Peek returns the top value from the stack without removing it
func (s *arrayStack) Peek() int {
	if s.Length == 0 {
		return -1
	}
	return s.Values[s.Length-1]
}

// Print prints the stack
func (s *arrayStack) Print() {
	fmt.Println("Stack:")
	for _, val := range s.Values {
		fmt.Println(val)
	}
	fmt.Println("End of Stack")
}
