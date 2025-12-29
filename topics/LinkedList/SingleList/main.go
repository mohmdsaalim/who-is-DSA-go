package main

import "fmt"

// Node represents a single node in the linked list
type Node struct {
	Value int
	Next  *Node
}

// SingleList represents the singly linked list
type SingleList struct {
	Head *Node
}

// NewSingleList creates and returns a new empty linked list
func NewSingleList() *SingleList {
	return &SingleList{}
}

// IsEmpty checks if the list is empty
func (s *SingleList) IsEmpty() bool {
	return s.Head == nil
}

// Append inserts a node at the end of the list
func (s *SingleList) Append(val int) {
	newNode := &Node{Value: val}

	if s.IsEmpty() {
		s.Head = newNode
		return
	}

	current := s.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Display prints all elements in the list
func (s *SingleList) Display() {
	if s.IsEmpty() {
		fmt.Println("empty list")
		return
	}

	current := s.Head
	fmt.Print("Head -> ")
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

// Delete removes the first node that matches the given value
func (s *SingleList) Delete(val int) {
	if s.IsEmpty() {
		fmt.Println("empty list")
		return
	}

	// Case: delete the head
	if s.Head.Value == val {
		s.Head = s.Head.Next
		return
	}

	prev := s.Head
	current := s.Head.Next

	for current != nil {
		if current.Value == val {
			prev.Next = current.Next
			return
		}
		prev = current
		current = current.Next
	}

	fmt.Printf("value %d not found in the list\n", val)
}

// Search checks whether a value exists in the list
func (s *SingleList) Search(val int) bool {
	current := s.Head
	for current != nil {
		if current.Value == val {
			return true
		}
		current = current.Next
	}
	return false
}

// Length returns the total number of nodes in the list
func (s *SingleList) Length() int {
	count := 0
	current := s.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func main() {
	list := NewSingleList()

	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Display() // Head -> 10 -> 20 -> 30 -> nil

	fmt.Println("Length:", list.Length())

	list.Delete(20)
	list.Display() // Head -> 10 -> 30 -> nil

	fmt.Println("Search 10:", list.Search(10))
	fmt.Println("Search 50:", list.Search(50))

	list.Delete(70) // not found
	list.Display()
}




// GENRIC VERSION 

// package main

// import "fmt"

// // Node of a singly linked list
// type Node[T comparable] struct {
//     Value T
//     Next  *Node[T]
// }

// // LinkedList with generic type
// type LinkedList[T comparable] struct {
//     Head *Node[T]
// }

// // Prepend / Insert at front
// func (ll *LinkedList[T]) Prepend(val T) {
//     ll.Head = &Node[T]{Value: val, Next: ll.Head}
// }

// // Append / Insert at end
// func (ll *LinkedList[T]) Append(val T) {
//     newNode := &Node[T]{Value: val}
//     if ll.Head == nil {
//         ll.Head = newNode
//         return
//     }
//     curr := ll.Head
//     for curr.Next != nil {
//         curr = curr.Next
//     }
//     curr.Next = newNode
// }

// // Search for first node with given value
// func (ll *LinkedList[T]) Search(val T) *Node[T] {
//     curr := ll.Head
//     for curr != nil {
//         if curr.Value == val {
//             return curr
//         }
//         curr = curr.Next
//     }
//     return nil
// }

// // Delete first occurrence of value
// func (ll *LinkedList[T]) Delete(val T) bool {
//     if ll.Head == nil {
//         return false
//     }
//     if ll.Head.Value == val {
//         ll.Head = ll.Head.Next
//         return true
//     }
//     curr := ll.Head
//     for curr.Next != nil {
//         if curr.Next.Value == val {
//             curr.Next = curr.Next.Next
//             return true
//         }
//         curr = curr.Next
//     }
//     return false
// }

// // DeleteAt removes a node at a given index (0-based)
// func (ll *LinkedList[T]) DeleteAt(index int) bool {
//     if index < 0 || ll.Head == nil {
//         return false
//     }
//     if index == 0 {
//         ll.Head = ll.Head.Next
//         return true
//     }
//     curr := ll.Head
//     for i := 0; curr != nil && i < index-1; i++ {
//         curr = curr.Next
//     }
//     if curr == nil || curr.Next == nil {
//         return false
//     }
//     curr.Next = curr.Next.Next
//     return true
// }

// // Reverse the linked list
// func (ll *LinkedList[T]) Reverse() {
//     var prev *Node[T]
//     curr := ll.Head
//     for curr != nil {
//         next := curr.Next
//         curr.Next = prev
//         prev = curr
//         curr = next
//     }
//     ll.Head = prev
// }

// // Print all values
// func (ll *LinkedList[T]) Print() {
//     curr := ll.Head
//     for curr != nil {
//         fmt.Printf("%v -> ", curr.Value)
//         curr = curr.Next
//     }
//     fmt.Println("nil")
// }



// func main() {
//     list := LinkedList[int]{}

//     list.Prepend(3)
//     list.Append(5)
//     list.Append(7)
//     list.Prepend(1)

//     list.Print()
//     // Output: 1 -> 3 -> 5 -> 7 -> nil

//     fmt.Println(list.Search(5).Value) // 5

//     list.Delete(3)
//     list.Print() // 1 -> 5 -> 7 -> nil

//     list.DeleteAt(1)
//     list.Print() // 1 -> 7 -> nil

//     list.Reverse()
//     list.Print() // 7 -> 1 -> nil
// }