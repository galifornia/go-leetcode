package main

import "fmt"

func singleNumber(nums []int) int {
	num := nums[0]
	for i := 1; i < len(nums); i++ {
		num = num ^ nums[i]
	}
	return num
}

// func singleNumber(nums []int) int {
// 	countMap := make(map[int]int)
// 	for _, v := range nums {
// 		countMap[v]++
// 	}

// 	for k, v := range countMap {
// 		if v == 1 {
// 			return k
// 		}
// 	}
// 	return -1
// }

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
