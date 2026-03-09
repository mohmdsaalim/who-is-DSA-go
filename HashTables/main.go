package main

import (
	"fmt"
	Array_hashing "hashing/array_hashing"
)



func main() {
	arr := []int{1,2,12,1,2,12,1,2,12,1,2,12,1,21,2,12}
	Array_hashing.Hashmapping(arr, 1)
	ar := Array_hashing.TwoSum(arr, 13)
	fmt.Println(ar)
}
