package main

import "fmt"

func mergesort(arr []int) []int {
	if len(arr) <= 1{
		return arr
	}
	mid := len(arr) / 2
	
	left := mergesort(arr[:mid])
	right := mergesort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	r := []int{}

	j := 0
	i := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			r = append(r, left[i])
			i++
		}else{
			r = append(r, right[j])
			j++
		}
	}
	r = append(r, left[i:]...)
	r = append(r, right[j:]...)
	return r
}
func main() {
	arr := []int{6,5,4,3,2,1}
	a := mergesort(arr)
	fmt.Println(a)
	
}
// completed the dsa sylabus eni ipppo onum illa after job