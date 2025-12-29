package main

import "fmt"

type DNode struct {
    Value int
    Prev  *DNode
    Next  *DNode
}

type DoublyList struct {
    Head *DNode
}
// appending
func (d *DoublyList) Append(val int) {
    newNode := &DNode{Value: val}

    if d.Head == nil {
        d.Head = newNode
        return
    }

    current := d.Head
    for current.Next != nil {
        current = current.Next
    }

    current.Next = newNode
    newNode.Prev = current
}

//Prepend
func (d *DoublyList) Prepend(val int) {
    newNode := &DNode{Value: val}

    if d.Head != nil {
        newNode.Next = d.Head
        d.Head.Prev = newNode
    }

    d.Head = newNode
}
// delte by value
func (d *DoublyList) Delete(val int) {
    if d.Head == nil {
        return
    }

    current := d.Head

    // Head deletion
    if current.Value == val {
        d.Head = current.Next
        if d.Head != nil {
            d.Head.Prev = nil
        }
        return
    }

    for current != nil {
        if current.Value == val {
            if current.Next != nil {
                current.Next.Prev = current.Prev
            }
            current.Prev.Next = current.Next
            return
        }
        current = current.Next
    }
}
// display
func (d *DoublyList) Display() {
    current := d.Head
    fmt.Print("nil <-> ")
    for current != nil {
        fmt.Printf("%d <-> ", current.Value)
        current = current.Next
    }
    fmt.Println("nil")
}


func main() {
    dl := DoublyList{}

    dl.Append(10)
    dl.Append(20)
    dl.Append(30)

    dl.Display()

    dl.Prepend(5)
    dl.Display()

    dl.Delete(20)
    dl.Display()
}