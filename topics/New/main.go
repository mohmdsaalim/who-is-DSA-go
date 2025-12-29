package main

import "fmt"
// new 
func main() {
	Data := new(int)
	fmt.Println(*Data)
	*Data = 60
	fmt.Println(*Data)
	str()
	mk()
	mp()
}
// sturct/new
type n struct{
	name string
	age int
}

func str(){
	super := new(n)
	super.name = "Saalim"
	super.age = 30
	fmt.Println(super)
}

// make
// build a slice with make func
func mk()  {
	slice := make([]int, 2)
	slice[0] = 10
	slice = append(slice, 10)
	slice = append(slice, 3)
	fmt.Println(slice)
}

func mp(){
	Data := make(map[string]any, 10)
	Data["messi"] = "footballer"
	fmt.Println(Data)
}