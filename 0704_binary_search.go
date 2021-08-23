package main

func search(nums []int, target int) int {
	numsLength := len(nums)
	left := 0
	right := numsLength - 1

	for left <= right {
		pivot := left + (right-left+1)/2

		if nums[pivot] == target {
			return pivot
		}

		if nums[pivot] < target {
			left = pivot + 1
		} else if nums[right] > target {
			right = pivot - 1
		}
	}
	return -1
}
