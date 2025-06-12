package stack

// Stack is an interface that defines the methods for a stack
type Stack interface {
	Push(value int)
	Pop() int
	Peek() int
	Height() int
	Print()
}

type node struct {
	value int
	next  *node
}

type stack struct {
	top    *node
	height int
}

// NewStack creates a new stack with the given values
// using a linked list stack implementation is more efficient if we don't know the size of the stack
// and we need to resize the underlying array and copy the values which is expensive
func NewStack(values []int) Stack {
	stack := &stack{}
	for _, value := range values {
		stack.Push(value)
	}
	return stack
}

type arrayStack struct {
	values []int
	length int
}

// NewArrayStack creates a new array stack with the given values and size
// using an array stack implementation is more efficient if we know the size of the stack
// and we don't need to resize the underlying array
func NewArrayStack(values []int, size int) Stack {
	stack := &arrayStack{
		values: make([]int, 0, size),
		length: 0,
	}
	for _, value := range values {
		stack.Push(value)
	}
	return stack
}
