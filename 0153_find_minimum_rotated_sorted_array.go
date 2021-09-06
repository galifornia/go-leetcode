package main

func findMin(nums []int) int {
	size := len(nums)
	left := 0
	right := size - 1

	// find rotation point
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
