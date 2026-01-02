package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

type Stack struct{
	top *Node
	size int
}
// empty chk func
func (s Stack) IsEmpty() bool {
	return s.top == nil	
}
// creating func
func NewStack() *Stack {
	return &Stack{top: nil, size: 0}
}
// Push
func (s *Stack) Push(val int){
	newnode := &Node{
		Value: val,
		Next: s.top,
	}
	s.top = newnode
	s.size ++
}
// Display
func (s Stack) Display(){
	if s.IsEmpty() {
		fmt.Println("No value in Stack")
	}
	current :=s.top
	fmt.Println("Top")
	for current != nil{
		fmt.Println(current.Value)
		current = current.Next
	}
	fmt.Println("bottom")
}
// pop
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0 , false
	}
	val := s.top.Value
	s.top = s.top.Next
	s.size --
	return val, true
}
// peek
func (s Stack) peek()(int, bool){
	if s.IsEmpty(){
		return 0, false
	}
	return s.top.Value, true
}
func main() {
	s := NewStack()
	// Push
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	s.Display()
	if v, ok := s.Pop(); ok{
		fmt.Println("Popped", v)
	}
	if v, ok :=s.peek();ok{
		fmt.Println("peek", v)
	}
	// displaying
	s.Display()

	//Pop
}