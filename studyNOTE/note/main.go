package main



type Node struct {
	Value int
	Next *Node
}

type Queue struct {
	front *Node
	rear *Node
	size int
}

func NewQue() *Queue{
	return &Queue{
		front: nil,
		rear: nil,
		size: 0,
	}
}
// Emque 
func (s * Queue) Enque(val int){
	newnode := &Node{Value: val}

	if s.front == nil{
		s.front = newnode
		s.rear = newnode
	}else{
		s.front.Next = newnode
		s.rear = newnode
	}
	s.size++
}
// deque 
func (s *Queue) Deque()(int, bool){
	if s.front == nil{
		return 0 , false
	}
	val := s.front.Value
	s.front = s.front.Next
	s.size--
	if s.front == nil{
		s.rear = nil
	}
	return val, true
}
func main() {
	
}