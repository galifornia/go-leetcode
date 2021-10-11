package main

import "math"

func maxSubArray(nums []int) int {
	sum := 0
	max := math.MinInt32

	for i := 0; i < len(nums); i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		if sum > max {
			max = sum
		}
	}

	return max
}
