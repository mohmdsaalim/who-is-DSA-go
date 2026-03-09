package Array_hashing

import "fmt"

// hash table using map in go
func Hashmapping(a []int,b int)  {
	m := make(map[int]int)

	for _, v := range a{
		m[v]++
	}
	fmt.Println("Count", m[b])
}