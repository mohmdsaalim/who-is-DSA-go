package main

import "fmt"



func quick(arr []int, low, high int)  {
	if low < high{
		pi := partition(arr, low, high)

		quick(arr, low, pi-1)
		quick(arr, pi+1, high)
	}
}
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low-1

	for j := low; j < high; j++{
		if arr[j] < pivot{
			i++
			arr[i],arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i+1
}

func main() {
	arr := []int{6,5,4,3,2,1}

	quick(arr, 0, len(arr)-1)

fmt.Println(arr)
}
