package main

import "fmt"

type node struct{
	Data int
	Next *node
}
type LinkedList struct{
	head *node
	length int
}
func (l *LinkedList) prepend(n *node){
	second := l.head
	l.head = n
	l.head.Next = second
	l.length++
}

func (l LinkedList) printData(){
	toprint := l.head
	for l.length != 0{
		fmt.Printf("%d ", toprint.Data)
		toprint = toprint.Next
		l.length--
	}
	fmt.Printf("\n")
}

func main() {
	myList := LinkedList{}
	node1 := &node{Data: 10}
	node2 := &node{Data: 12}
	node3 := &node{Data: 13}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)

myList.printData()
}