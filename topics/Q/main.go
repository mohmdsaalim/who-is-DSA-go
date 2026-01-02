package main

import "fmt"

// Node structure
type Node struct {
	Value int
	Next  *Node
}

// Queue structure
type Queue struct {
	front *Node
	rear  *Node
	size  int
}

// Create new queue
func NewQueue() *Queue {
	return &Queue{
		front: nil,
		rear:  nil,
		size:  0,
	}
}

// EnQueue (O(1))
func (q *Queue) EnQueue(val int) {
	newnode := &Node{Value: val}

	if q.front == nil {
		q.front = newnode
		q.rear = newnode
	} else {
		q.rear.Next = newnode
		q.rear = newnode
	}
	q.size++
}

// DeQueue (O(1))
func (q *Queue) DeQueue() (int, bool) {
	if q.front == nil {
		return 0, false
	}

	val := q.front.Value
	q.front = q.front.Next
	q.size--

	if q.front == nil {
		q.rear = nil
	}

	return val, true
}

// Peek (O(1))
func (q *Queue) Peek() (int, bool) {
	if q.front == nil {
		return 0, false
	}
	return q.front.Value, true
}

// IsEmpty
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Size
func (q *Queue) Size() int {
	return q.size
}

// Display
func (q *Queue) Display() {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return
	}

	fmt.Print("Front -> ")
	current := q.front
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println("<- Rear")
}

// Main function
func main() {
	q := NewQueue()

	q.EnQueue(10)
	q.EnQueue(20)
	q.EnQueue(30)

	q.Display()

	val, _ := q.DeQueue()
	fmt.Println("Dequeued:", val)

	q.Display()

	front, _ := q.Peek()
	fmt.Println("Front element:", front)

	fmt.Println("Size:", q.Size())
}