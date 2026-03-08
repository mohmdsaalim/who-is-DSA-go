package selectionsort

import "fmt"

func Selection(arr []int) []int {
	n := len(arr)
	a := arr
	for i := 0; i < n-1; i++{
		min := i
		
		for j := i+1; j < n; j++{
			if a[j] < a[min]{
				a[min] = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
	fmt.Println("this is selection sort")
	return a
}
