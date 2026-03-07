package main

import (
	bubblesort "dsa-algo/bubble_sort"
	insertionsort "dsa-algo/insertion_sort"
	selectionsort "dsa-algo/selection_sort"
	"fmt"
)

func main() {
	arr := []int{2,4,3,1,5,7,6,934,5,5,4564,456,456,456,546,4,456,56,564,6,66}

	bubble_sort := bubblesort.BubbleSort(arr) // swapping with next value
	selection_sort := selectionsort.Selection(arr) // find smallest number and swap
	insertion_sort := insertionsort.Insertion(arr) // comparing with neibrs

	fmt.Println(bubble_sort)
	fmt.Println(selection_sort)
	fmt.Println(insertion_sort)
}
// 