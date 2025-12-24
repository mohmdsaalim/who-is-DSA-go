package main

import "fmt"

type User struct{
	Name string
	age int
} 

var neww []int// *
var arrr =  make([]int, 10)// initializing  a Slice -  O(n)ST

func main() {
	arrr[0] = 10 //. SET opertaion in here - O(1)ST , when the value is updated by for loop the time will br O(n) becouse of thr loop,for
	// GET also similar to the SET and also in the case of the for loop are also same
	// TRAVERSE the space is not effect it will O(1)S and time is O(n)T. 
	fmt.Println(arrr)
	arr := []int{6,2,3,4,5}
	r(arr, 10)
	lol(arr)
}

func r(a []int, b int){
	for i := 0; i < len(a); i++{
		for j := i+1; j < len(a); j++{
			if a[i]+a[j] == b{
				neww = append(neww, a[i], a[j])// *
			}
		}
		
	}
	fmt.Println(neww)
}

// find the min/max number of array
func lol(arr []int) {
	max := arr[0]
	for _,v := range arr{
		if max < v{
			max = v
		}
	}
	fmt.Println(max)
}

