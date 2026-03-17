package main

import "fmt"

func heapSort(arr []int) {
	n := len(arr)

	// Step 1: Build Max Heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Step 2: Extract elements from heap
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		arr[0], arr[i] = arr[i], arr[0]

		// Call heapify on reduced heap
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// Check left child
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// Check right child
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]

		// Recursively heapify the affected subtree
		heapify(arr, n, largest)
	}
}

func main() {
	arr := []int{6, 5, 4, 3, 2, 1}
	heapSort(arr)
	fmt.Println(arr)
}