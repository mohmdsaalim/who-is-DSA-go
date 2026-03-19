package main

func heapsort(arr []int)  {
	n := len(arr)

	for i := n/2-1; i>=0; i++{
		heapfy(arr, n, i)
	}

	for i := n-1; i > 0 ; i--{
		arr[0], arr[i] = arr[i], arr[0]
		heapfy(arr, i, 0)
	}
}
func heapfy(arr []int, n, i int)  {
	
}
func main() {
	
}