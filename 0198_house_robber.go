package main

import (
	"math"
)

func rob(nums []int) int {
	size := len(nums)
	memo := make([]int, size+1)

	for i := 0; i < size+1; i++ {
		memo[i] = math.MinInt32
	}
	return findBiggestReward(nums, size-1, &memo)
}

func findBiggestReward(nums []int, i int, memo *[]int) int {
	if i < 0 {
		return 0
	}
	if (*memo)[i] >= 0 {
		return (*memo)[i]
	}
	result := int(math.Max(float64(findBiggestReward(nums, i-2, memo)+nums[i]), float64(findBiggestReward(nums, i-1, memo))))
	(*memo)[i] = result

	return result
}

// func main() {
// 	// nums := []int{1, 2, 3, 1}
// 	// nums := []int{2, 7, 9, 3, 1}
// 	nums := []int{1, 4, 1, 2, 3, 1, 1, 5, 1}
// 	fmt.Println(rob(nums))
// }
