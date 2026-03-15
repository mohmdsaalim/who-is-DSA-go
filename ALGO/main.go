package main

import (
	// bubblesort "dsa-algo/bubble_sort"
	insertionsort "dsa-algo/insertion_sort"
	// selectionsort "dsa-algo/selection_sort"
	"fmt"
)

func main() {
	arr := []int{2,4,3,1,5,7,6,934,5,5,4564,456,456,456,546,4,456,56,564,6,66}
	
	fmt.Println(arr)
	// bubble_sort := bubblesort.BubbleSort(arr) // swapping with next value
	// selection_sort := selectionsort.Selection(arr) // find smallest number and swap
	insertion_sort := insertionsort.Insertion(arr) // comparing with neibrs


	// fmt.Println(bubble_sort)
	// selection_sort_test(arr)
	// fmt.Println(selection_sort)
	fmt.Println(insertion_sort)
}

// bubble sort -> cechking with neighers
func selection_sort_test(arr []int)  {
	n := len(arr)

	for i := 0; i < n-1; i++{
		min := i 
		for j := 0; j < n; j++{
			if arr[j] < arr[min]{
				min = j
			}
		} // 0 != j
		arr[i], arr[min] = arr[min], arr[i]
	}
	fmt.Println(arr)
}