package main

import "fmt"

type User struct{
	Name string
	age int
} 

var neww []int
func main() {
	arr := []int{6,2,3,4,5}
	r(arr, 10)
}

func r(a []int, b int){
	for i := 0; i < len(a); i++{
		for j := i+1; j < len(a); j++{
			if a[i]+a[j] == b{
				neww = append(neww, a[i], a[j])
			}
		}
		
	}
	fmt.Println(neww)
}