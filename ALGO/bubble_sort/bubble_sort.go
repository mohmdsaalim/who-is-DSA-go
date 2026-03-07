package bubblesort

func BubbleSort(a []int) []int {
	n := len(a)
	arr := a
	for i := 0; i < n-1; i++{
		for j := 0; j < n-i-1; j++{
			if arr[j] > arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}
	return  arr
}
// 