// package main
// // Using *Type allows Go to work with addresses instead of values, enabling efficient memory usage, shared access, and in-place modification of data.
// import "fmt"
// //& adress of
// // pointing to
// func main() {
// 	a := 10
// ad(a)// dont use it
// Add(&a)// proffesinla style
// }

// // immutabity doesnt on change in here not in memory
// func ad(p int){
// p *= p
// fmt.Println(p)
// }

// // effincency
// func Add(p *int){
// *p *= *p
// fmt.Println(p, *p)
// }

package main

import "fmt"
type node struct{
	data int
	next *node
}

func main() {
	user1 := &node{10,nil}
	user2 := &node{20,nil}
	user1.next = user2
	fmt.Println(user1,user2)
}






















