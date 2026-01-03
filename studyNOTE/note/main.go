package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

type Queue struct {
	front *Node
	rear *Node
	size int
}

func NewQueue()*Queue{
	return &Queue{
		front: nil,
		rear: nil,
		size: 0,
	}
}
// EnQueue
func (s *Queue) EnQueue(val int){
	newnode := &Node{Value: val}
	if s.front == nil{
		s.front = newnode
		s.rear = newnode
	}else{
		s.rear.Next = newnode
		s.rear = newnode
	}
	s.size++
}
// Dequeue
func (s *Queue) DeQueue()(int, bool) {
	if s.front == nil{
		return 0, false
	}
	val := s.front.Value
	s.front = s.front.Next
	s.size --
	return val,true
}
func (s *Queue) Display(){
	if s.front == nil{
		fmt.Println("No value in here")
	}
	current := s.front
	for current != nil{
		fmt.Println(current.Value)
		current = current.Next
	}
}
func main() {
	que := NewQueue()
	que.EnQueue(10)
		que.EnQueue(20)
							que.EnQueue(30)
							
							que.Display()
							que.DeQueue()
}