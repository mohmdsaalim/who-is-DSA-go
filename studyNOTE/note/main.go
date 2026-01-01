package main

type Node struct {
	Data int
	next *Node
}
type linkedlist struct {
	head *Node
}
// append
func (l *linkedlist) Append(val int){
	newnode := &Node{Data: val}
	if l.head == nil{
		l.head =newnode
		return
	}
	current := l.head
	for current.next != nil{
		current = current.next
	}
	current.next = newnode
}
// Delete 
func (l *linkedlist) Delete(val int){
	if l.head.Data == val{
		l.head = l.head.next
		return
	}
	prev := l.head
	current :=l.head.next
	for current != nil{
		if current.Data == val{
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
	}
}
// Prepend
func (l *linkedlist) Prepend(val int) {
	newnode := &Node{Data: val}

	newnode.next = l.head
	l.head = newnode
}

func main() {

}