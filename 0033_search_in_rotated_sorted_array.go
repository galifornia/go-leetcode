package main

func search(nums []int, target int) int {
	left := 0
	size := len(nums)
	right := size - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	rot := left
	left = 0
	right = size - 1
	for left <= right {
		mid := (left + right) / 2
		realmid := (mid + rot) % size
		if nums[realmid] == target {
			return realmid
		}

		if nums[realmid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
