package recursion

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Sum(arr []int, i int) int {
	if i == len(arr) {
		return 0
	}
	return arr[i] + Sum(arr, i+1)
}