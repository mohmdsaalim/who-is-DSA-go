package main

import "fmt"

//nodes
type Node struct{
	Data int
	next *Node
}


//list
type LinkedList struct{
	head *Node
}


// Newnode
func SingleList() *LinkedList{
	return &LinkedList{}
}


// empty check
func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}



//append
func (l *LinkedList)Append(val int){
newNode := &Node{Data: val}
	if l.IsEmpty(){
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil{
		 current = current.next 
	}
	current.next = newNode
}


// Display 
func (l LinkedList) Dipaly()  {
	if l.IsEmpty(){
		fmt.Println("No Value in LinkedList")
		return
	}
	current := l.head
	fmt.Print("Head --> ")
	for current != nil{
		fmt.Printf("%d -->", current.Data)
		current = current.next
	}
	fmt.Print("nill \n")
}


//delete
func (l *LinkedList) Delete(val int)  {
	if l.IsEmpty(){
		fmt.Println("There no value in LinkedList")
	}
	if l.head.Data == val{
		l.head = l.head.next
		return
	}
	prev := l.head
	current := l.head.next

	for current != nil{
		if current.Data == val{
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
	}
	fmt.Printf("value not fount %d\n",val)
}


// SEARCH 
func (l *LinkedList) Search(val int) bool{
	if l.IsEmpty(){
		fmt.Println("There no Value")
	}
	current := l.head
	for current != nil{
		if current.Data == val{
			fmt.Println("value fount", val)
			return true
		}
		current = current.next
	}
	fmt.Println("No value in list", val)
	return false
}


// MAIN
func main() {
	list := SingleList()
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Delete(1)
	list.Dipaly()
	list.Search(10)
}