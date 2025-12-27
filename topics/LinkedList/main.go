package main

import "fmt"

type Node struct{
	Data int
	Next *Node
}

type LinkedList struct{
	Head *Node
}

func (l *LinkedList) IsEmpty() bool{
	return l.Head == nil
}
// NewList
func NewList() *LinkedList{
	return &LinkedList{}
}

// Append 
func (l *LinkedList) Append(data int){
	newNode := &Node{Data: data}

	if l.IsEmpty(){
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil{
		current = current.Next
	}
	current.Next = newNode
}
//____________________________________________________________________________
//____________________________________________________________________________
//delete
func (l *LinkedList) Delete(val int){
	if l.IsEmpty(){
		fmt.Println("The List Is Empty")
		return
	}
	if l.Head.Data == val{
		l.Head = l.Head.Next
		return
	}
	prev := l.Head
	current := l.Head.Next

	for current != nil{
		if current.Data == val{
			prev.Next = current.Next
			return
		}
		prev = current
		current = current.Next
	}
 	fmt.Printf("value %d not found in the list\n", val)
}
// Display 



func (l *LinkedList) Display(){
	if l.IsEmpty(){
		fmt.Println("There is no value in List")
	}
	current := l.Head
	for current != nil{
		fmt.Printf("%d -> ",current.Data)
		current = current.Next
	}
		fmt.Print("nil")
}

func main() {
list := NewList()

list.Append(10)
list.Append(20)
list.Append(30)

list.Delete(20)
list.Display()
}