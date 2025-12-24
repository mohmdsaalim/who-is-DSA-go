// package main

// import "fmt"

// var newarray []int
// var firstarray []int

// func main() {
// 	arr := []int{6,2,3,4,6,7,9,6,7,6}
// 	for _,v := range arr{
// 		if v == 6{
// 			newarray = append(newarray, v)
// 			}else{
// 				firstarray = append(firstarray, v)
// 			}
// 		}
// 		mainarr := append(firstarray, newarray...)
// 	fmt.Println(mainarr)

// }

package main

import "fmt"

func main() {
	arr := []int{2,7,11,15}
	fmt.Println(twoSum(arr, 9))
}

// func last(arr []int) []int {
// 	arr1 := make([]int, 0,len(arr))
// 	for _,v := range arr{
// 		if v != 6{
// 			arr1 = append(arr1, v)
// 		}

// 	}
// 		for _,v :=  range arr{
// 			if v == 6{

// 				arr1 = append(arr1, v)
// 			}
// 		}
// 	return arr1
// }

// func twoSum(nums []int, target int) []int {
//     newarray := []int{}
//     for i := 0; i < len(nums);i++{ // o(n)
//         for j := i+1; j < len(nums); j++{ // O(n) o(n2)
//                 if nums[i] + nums[j] == target{
//                     newarray = append(newarray, i,j)
//                 }
//         }
//     }
//     return newarray
//     }

    func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, v := range nums {

        if idx, ok := m[target-v]; ok {
            return []int{idx, i}
        }
        m[v] = i
    }
    return []int{}
}