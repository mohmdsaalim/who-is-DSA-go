package main

import "fmt"
//& adress of 
// pointing to
func main() {
	i, j := 10, 10

	pi := &i
	pj := &j
	*pi = 300
	newpi := *pi
	fmt.Print(pi, pj,newpi )
}