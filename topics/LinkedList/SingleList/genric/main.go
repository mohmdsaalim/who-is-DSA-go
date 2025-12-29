

package main

import "fmt"

// Node of a singly linked list
type Node[T comparable] struct {
    Value T
    Next  *Node[T]
}

// LinkedList with generic type
type LinkedList[T comparable] struct {
    Head *Node[T]
}

// Prepend / Insert at front
func (ll *LinkedList[T]) Prepend(val T) {
    ll.Head = &Node[T]{Value: val, Next: ll.Head}
}

// Append / Insert at end
func (ll *LinkedList[T]) Append(val T) {
    newNode := &Node[T]{Value: val}
    if ll.Head == nil {
        ll.Head = newNode
        return
    }
    curr := ll.Head
    for curr.Next != nil {
        curr = curr.Next
    }
    curr.Next = newNode
}

// Search for first node with given value
func (ll *LinkedList[T]) Search(val T) *Node[T] {
    curr := ll.Head
    for curr != nil {
        if curr.Value == val {
            return curr
        }
        curr = curr.Next
    }
    return nil
}

// Delete first occurrence of value
func (ll *LinkedList[T]) Delete(val T) bool {
    if ll.Head == nil {
        return false
    }
    if ll.Head.Value == val {
        ll.Head = ll.Head.Next
        return true
    }
    curr := ll.Head
    for curr.Next != nil {
        if curr.Next.Value == val {
            curr.Next = curr.Next.Next
            return true
        }
        curr = curr.Next
    }
    return false
}

// DeleteAt removes a node at a given index (0-based)
func (ll *LinkedList[T]) DeleteAt(index int) bool {
    if index < 0 || ll.Head == nil {
        return false
    }
    if index == 0 {
        ll.Head = ll.Head.Next
        return true
    }
    curr := ll.Head
    for i := 0; curr != nil && i < index-1; i++ {
        curr = curr.Next
    }
    if curr == nil || curr.Next == nil {
        return false
    }
    curr.Next = curr.Next.Next
    return true
}

// Reverse the linked list
func (ll *LinkedList[T]) Reverse() {
    var prev *Node[T]
    curr := ll.Head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    ll.Head = prev
}

// Print all values
func (ll *LinkedList[T]) Print() {
    curr := ll.Head
    for curr != nil {
        fmt.Printf("%v -> ", curr.Value)
        curr = curr.Next
    }
    fmt.Println("nil")
}



func main() {
    list := LinkedList[int]{}

    list.Prepend(3)
    list.Append(5)
    list.Append(7)
    list.Prepend(1)

    list.Print()
    // Output: 1 -> 3 -> 5 -> 7 -> nil

    fmt.Println(list.Search(5).Value) // 5

    list.Delete(3)
    list.Print() // 1 -> 5 -> 7 -> nil

    list.DeleteAt(1)
    list.Print() // 1 -> 7 -> nil

    list.Reverse()
    list.Print() // 7 -> 1 -> nil
}