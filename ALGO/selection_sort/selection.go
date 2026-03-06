package selectionsort

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
	return a
}

// func(a){
// for i = 0; i < n-1; i++{
// 		min := i
// 		for j := 0; j < n; j ++{
// 			if arr[j] < a[min]{
// 			a[min] = j	
// }		
//}
// a[j], a[min] = a[min] = a[j]
// }
// }


