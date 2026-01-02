package main

import "fmt"

// Node structure
type Node struct {
	Value int
	Next  *Node
}

// Stack structure
type Stack struct {
	top  *Node
	size int
}

// Create new stack
func NewStack() *Stack {
	return &Stack{
		top:  nil,
		size: 0,
	}
}

// Push (O(1))
func (s *Stack) Push(val int) {
	newnode := &Node{
		Value: val,
		Next:  s.top,
	}
	s.top = newnode
	s.size++
}

// Pop (O(1))
func (s *Stack) Pop() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	val := s.top.Value
	s.top = s.top.Next
	s.size--
	return val, true
}

// Peek (O(1))
func (s *Stack) Peek() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	return s.top.Value, true
}

// IsEmpty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Size
func (s *Stack) Size() int {
	return s.size
}

// Display
func (s *Stack) Display() {
	if s.top == nil {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Print("Top -> ")
	current := s.top
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}

// Main function
func main() {
	stack := NewStack()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	stack.Display()

	val, _ := stack.Pop()
	fmt.Println("Popped:", val)

	stack.Display()

	top, _ := stack.Peek()
	fmt.Println("Top element:", top)

	fmt.Println("Size:", stack.Size())
}