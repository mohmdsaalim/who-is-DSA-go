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


func TwoSum(arr []int, a int) []int {

	m := make(map[int]int) 
	for i, v := range arr{
		n := a - v
		// fmt.Println(i, v)
		if idx, ok := m[n]; ok{ 
			return []int{idx, i}
		}
		m[v]=i 
	}
	return []int{}
}