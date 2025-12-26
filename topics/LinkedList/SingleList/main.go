package main

import (
	"fmt"
)

type Node struct{
	Data int
	Next *Node
}

type SingleList struct{
	Head *Node
}
// NewSingleList create and returns a new empty linked List
func NewSingleList() *SingleList{
	return &SingleList{}
}
// IsEmpty checks if the list is empty
func (s *SingleList) IsEmplty() bool {
	return s.Head == nil
}


// Append inserts a node at the end of the list
func (s *SingleList) Append(data int) {
	newnode := &Node{Data: 10}

	if s.IsEmplty(){
		s.Head = newnode
		return
	}
	current := s.Head
	for current.Next != nil{
		current = current.Next
	}
	current.Next = newnode
}


// Display prints all elements in the list
func (s *SingleList) Display(){
	if s.IsEmplty(){
		fmt.Println("empty list")
		return
	}

	current := s.Head
	fmt.Print("Head -> ")
	for current != nil{
		fmt.Printf("%d ->", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}


func main() {
list := NewSingleList()

list.Append(10)
list.Append(20)
list.Append(30)

list.Display()
}