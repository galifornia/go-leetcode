package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {
	return findBiggestReward(nums, len(nums)-1)
}

func findBiggestReward(nums []int, i int) int {
	if i < 0 {
		return 0
	}
	return int(math.Max(float64(findBiggestReward(nums, i-2)+nums[i]), float64(findBiggestReward(nums, i-1))))
}

func main() {
	// nums := []int{1, 2, 3, 1}
	// nums := []int{2, 7, 9, 3, 1}
	nums := []int{1, 4, 1, 2, 3, 1, 1, 5, 1}
	fmt.Println(rob(nums))
}
