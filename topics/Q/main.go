package main

import "fmt"

type Q struct{
	Data []int
}
// 
func (q *Q) add(val int){
	q.Data = append(q.Data, val)
}

func (q *Q) del() int{
	toRemove := q.Data[0]
	q.Data = q.Data[1:]
	return toRemove
}

func main() {
	Queue := Q{}
	Queue.add(10)
	Queue.add(30)
	Queue.add(120)
	Queue.add(1034)
		 fmt.Println(Queue)
	Queue.del()
	fmt.Println(Queue)
}