package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}
// Insert value at beginning
func Insert(head **Node, data int) {
	newNode := &Node{Data: data}
	newNode.Next = *head
	*head = newNode
}
// Print linked list
func PrintList(head *Node) {
	for head != nil {
		fmt.Print(head.Data, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}
func main() {
	var head *Node = nil // empty list
	
fmt.Println(&head)

	Insert(&head, 10)
	PrintList(head)

	Insert(&head, 20)
	PrintList(head)

	Insert(&head, 30)
	PrintList(head)
}