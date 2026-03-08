package main

import "fmt"

func main() {
	arr := []int{1,3,4,1,1,3,4,45,45,5,6,567,5,4,334,34,1,2}
	a  := BubbleSort(arr)  
	// b := selection_sort(arr)
	fmt.Println(a)  
	// fmt.Println(b)

}

func BubbleSort(a []int) []int {
	n := len(a)

	for i := 0; i < n-1; i ++{
		for j := 0; j < n-i-1; j++{
			if a[j] > a[j+1]{
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

// swapping 
func selection_sort(a []int) []int {
	
	arr := a
	n := len(arr)
	for i := 0; i < n-1; i++{
		min := i

		for j := i+1; j < n; j++{
			// fmt.Println(i)
			if arr[j] < arr[min]{
				arr[min] = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	fmt.Println(arr)
	return arr
}

