package linkedlist

// Node represents a single node
type Node struct {
	Data int
	Next *Node
}

// SinglyList represents the linked list
type SinglyList struct {
	Head *Node
}

// ---------- Utility ----------

// Check if list is empty
func (l *SinglyList) IsEmpty() bool {
	return l.Head == nil
}

// ---------- Insert Operations ----------

// Insert at front (O(1))
func (l *SinglyList) InsertFront(val int) {
	l.Head = &Node{Data: val, Next: l.Head}
}

// Insert at end (O(n))
func (l *SinglyList) InsertEnd(val int) {
	newNode := &Node{Data: val}
	if l.IsEmpty() {
		l.Head = newNode
		return
	}
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

// Insert at position (O(n))
func (l *SinglyList) InsertAt(pos int, val int) {
	if pos <= 0 {
		l.InsertFront(val)
		return
	}
	curr := l.Head
	for i := 0; curr != nil && i < pos-1; i++ {
		curr = curr.Next
	}
	if curr == nil {
		return
	}
	curr.Next = &Node{Data: val, Next: curr.Next}
}

// ---------- Delete Operations ----------

// Delete from front (O(1))
func (l *SinglyList) DeleteFront() {
	if l.IsEmpty() {
		return
	}
	l.Head = l.Head.Next
}

// Delete from end (O(n))
func (l *SinglyList) DeleteEnd() {
	if l.IsEmpty() || l.Head.Next == nil {
		l.Head = nil
		return
	}
	curr := l.Head
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil
}

// Delete by value (O(n))
func (l *SinglyList) Delete(val int) {
	if l.IsEmpty() {
		return
	}
	if l.Head.Data == val {
		l.Head = l.Head.Next
		return
	}
	curr := l.Head
	for curr.Next != nil {
		if curr.Next.Data == val {
			curr.Next = curr.Next.Next
			return
		}
		curr = curr.Next
	}
}

// ---------- Search & Access ----------

// Search value (O(n))
func (l *SinglyList) Search(val int) bool {
	curr := l.Head
	for curr != nil {
		if curr.Data == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

// Get length (O(n))
func (l *SinglyList) Length() int {
	count := 0
	curr := l.Head
	for curr != nil {
		count++
		curr = curr.Next
	}
	return count
}

// ---------- Advanced Operations ----------

// Reverse list (O(n))
func (l *SinglyList) Reverse() {
	var prev *Node
	curr := l.Head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}

// Convert list to slice (O(n))
func (l *SinglyList) ToSlice() []int {
	var result []int
	curr := l.Head
	for curr != nil {
		result = append(result, curr.Data)
		curr = curr.Next
	}
	return result
}