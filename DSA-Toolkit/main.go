package main

import (
	"fmt"

	"dsa-toolkit/arrays"
	"dsa-toolkit/linkedlist"
	"dsa-toolkit/stack"
	"dsa-toolkit/queue"
	"dsa-toolkit/recursion"
)

func main() {
	// Arrays
	arr := []int{1, 2, 3}
	arr = arrays.Insert(arr, 1, 10)
	fmt.Println("Array:", arr)

	// Singly Linked List
	list := linkedlist.SinglyList{}
	list.InsertEnd(10)
	list.InsertEnd(20)
	list.InsertEnd(30)
	list.Delete(20)
	fmt.Println("LinkedList:", list.ToSlice())

	// Stack
	s := stack.Stack{}
	s.Push(5)
	s.Push(10)
	v, _ := s.Pop()
	fmt.Println("Stack Pop:", v)

	// Queue
	q := queue.Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	d, _ := q.Dequeue()
	fmt.Println("Queue Dequeue:", d)

	// Recursion
	fmt.Println("Factorial:", recursion.Factorial(5))
}