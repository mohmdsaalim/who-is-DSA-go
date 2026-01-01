package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Linkedlist struct {
	Head *Node
}
// node
func NewList()  *Linkedlist{
	return &Linkedlist{}
}
// append
func (l *Linkedlist) Append(val int) {
	newnode := &Node{Data: val}

	if l.Head == nil{
		l.Head = newnode
		return
	}

	current := l.Head
	for current.Next != nil{
		current = current.Next
	}
	current.Next = newnode

}
// Delete 
func (l *Linkedlist) Delete(val int){
		if l.Head == nil{
			fmt.Println("no value")
		}
		// head
		if l.Head.Data == val{
			l.Head = l.Head.Next
		}
		// looping
		prev := l.Head
		current := l.Head.Next
		for current!= nil{
			if current.Data == val{
				prev.Next = current.Next
			}
			prev = current
			current = current.Next
		}
}
//Display
func (l Linkedlist) Display() {

if l.Head == nil{
	fmt.Println("There is no value to display")
}	
current := l.Head
	for current != nil{
		fmt.Printf("%d ->",current.Data)
		current = current.Next
	}
}
func main() {
	ll := NewList()
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	ll.Append(40)
	ll.Delete(10)
	ll.Display()
}