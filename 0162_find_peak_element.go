package main

func findPeakElement(nums []int) int {
	return search(nums, 0, len(nums)-1)
}

func search(nums []int, left int, right int) int {
	if left == right {
		return left
	}

	mid := (left + right) / 2
	if nums[mid] > nums[mid+1] {
		return search(nums, left, mid)
	} else {
		return search(nums, mid+1, right)
	}
}
