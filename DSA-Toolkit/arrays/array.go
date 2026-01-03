package arrays

// Insert value at position
func Insert(arr []int, index int, val int) []int {
	arr = append(arr, 0)
	copy(arr[index+1:], arr[index:])
	arr[index] = val
	return arr
}

// Delete value at position
func Delete(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

// Reverse array
func Reverse(arr []int) []int {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}