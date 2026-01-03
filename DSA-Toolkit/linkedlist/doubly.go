package linkedlist

type DNode struct {
	Data int
	Prev *DNode
	Next *DNode
}

type DoublyList struct {
	Head *DNode
}

func (l *DoublyList) InsertFront(val int) {
	n := &DNode{Data: val}
	if l.Head != nil {
		l.Head.Prev = n
		n.Next = l.Head
	}
	l.Head = n
}

func (l *DoublyList) Delete(val int) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == val {
		l.Head = l.Head.Next
		if l.Head != nil {
			l.Head.Prev = nil
		}
		return
	}
	c := l.Head
	for c != nil {
		if c.Data == val {
			if c.Prev != nil {
				c.Prev.Next = c.Next
			}
			if c.Next != nil {
				c.Next.Prev = c.Prev
			}
			return
		}
		c = c.Next
	}
}