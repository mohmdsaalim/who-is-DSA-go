// package main

// import "fmt"

// type Stack struct{
// 	Data []int
// }
// // push

// func (s *Stack) push(val int){
// 	s.Data = append(s.Data, val)
// }

// // pop
// func (s *Stack) pop() int {
// l := len(s.Data)-1
// Tr := s.Data[l]
// s.Data = s.Data[:l]
// return Tr
// }


// func main() {
// 	stack := Stack{}
// 	stack.push(10)
// 	stack.push(20)
// 	stack.push(30)
// 	stack.push(40)
// 	fmt.Println(stack)
// 	stack.pop()
// 	fmt.Println(stack)

// }


package main

import (
    "fmt"
    "strings"
)

func WordFrequency(s string) map[string]int {
    freq := make(map[string]int)
    words := strings.Fields(s)

    for _, word := range words {
        freq[word]++
    }
    return freq
}

func main() {
    text := "go is fast go is simple"
    fmt.Println(WordFrequency(text))
}