// package main

// import "fmt"
// type Node struct{
// 	Data int
// 	Prev *Node
// 	Next *Node
// }

// type LinkedList struct{
// 	head *Node
// }
// func SetList() *LinkedList {
// 	return &LinkedList{}
// }
// // empty chk
// func (l *LinkedList) IsEmpty() bool{
// 	return l.head == nil
// }
// // append
// func (l *LinkedList) Append(val int) {
// 	newnode := &Node{Data: val}

// 	if l.IsEmpty() {
// 		l.head = newnode
// 		return
// 	}
// 	current := l.head
// 	for current.Next != nil{
// 		current = current.Next
// 	}
// 	current.Next = newnode
// 	newnode.Prev = current
// }

// // prepend
// func (l *LinkedList) Prepend(val int){
// 	newnode := &Node{Data: val}

// 	if l.head != nil{
// 		newnode.Next = l.head
// 		l.head.Prev = newnode
// 	}
// 	l.head = newnode
// }

// // delete
// func (l *LinkedList) Delete(val int){
// 	// checking is empty
// 	if l.IsEmpty(){
// 		fmt.Println("There is no value")
// 		return
// 	}

// 	current := l.head
// 	for current != nil{
// 		if current.Data == val{
// 			if current.Prev != nil{
// 				current.Prev.Next = current.Next
// 			}else{
// 				l.head = current.Next
// 			}

// 			if current.Next != nil{
// 				current.Next.Prev = current.Prev
// 			}
// 			fmt.Println("deleted")
// 			return
// 		}
// 		current = current.Next
// 	}
// 	fmt.Println("no value found")
// }

// // Display
// func (l *LinkedList) Display(){
// 	if l.IsEmpty(){
// 		fmt.Println("NO value")
// 		return
// 	}
// 	current := l.head
// 	fmt.Print("Head --> ")
// 	for current != nil{
// 		fmt.Printf("%d <--> ", current.Data)
// 		current = current.Next
// 	}
// fmt.Println("Nill")
// }
// func main() {
// 	list := SetList()
// 	list.Append(10)
// 	list.Append(20)
// 	list.Append(30)
// 	list.Append(60)
// 	list.Append(40)
// 	list.Append(50)

// 	list.Delete(30)
// 	list.Display()

// }
package main

import "fmt"

type Node struct{
	Data int
	Next *Node
}
type LinkedList struct{
	head *Node
}
// cr func
func List() *LinkedList{
	return &LinkedList{}
}
// empty check
func (l LinkedList)IsEmpty()bool{
	return l.head == nil
}
// append
func (l *LinkedList) Append(val int){
	newNode := &Node{Data: val}

	if l.IsEmpty(){
		l.head = newNode
		return
	}
	current := l.head
	for current.Next != nil{
		current = current.Next
	}
	current.Next = newNode
}
// Display
func (l *LinkedList) Display(){
	if l.IsEmpty(){
		fmt.Print("No Value")
	}
	current := l.head
	fmt.Print("Head -> ")
	for current != nil{
		fmt.Printf(" %d -> ", current.Data)
		current = current.Next
	}
	fmt.Print("Nil")
}
// search
func (l *LinkedList) srch(val int){
	current := l.head

	for current != nil{
		if current.Data == val {
			fmt.Printf("The value %d is available ",current.Data)
			return
		}
		current = current.Next
	}
	fmt.Println("The value is not available")
}
// delete
func (l *LinkedList) Delete(val int) {
	if l.IsEmpty() {
		fmt.Println("there is no value to delete")
		return
	}

	// Case 1: delete head
	if l.head.Data == val {
		l.head = l.head.Next
		fmt.Println("value deleted")
		return
	}

	prev := l.head
	current := l.head.Next

	for current != nil {
		if current.Data == val {
			prev.Next = current.Next
			fmt.Println("value deleted")
			return
		}
		prev = current
		current = current.Next
	}

	fmt.Println("value not found")
}
func main() {
	ll := List()
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	ll.Append(40)

	ll.Delete(10)
	//display
	ll.Display()
	ll.srch(20)
	
}