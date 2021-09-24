package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}

		if max < nums[i]+i {
			max = nums[i] + i
		}
	}
	return true
}

func main() {
	// nums := []int{2, 3, 1, 1, 4}
	// nums := []int{3, 2, 1, 0, 4}
	nums := []int{2, 0, 0}
	fmt.Println(canJump(nums))
}
