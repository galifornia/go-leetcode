package main

import "fmt"

func searchRange(nums []int, target int) []int {
	size := len(nums)
	// quick return when there are no items in array
	if size == 0 {
		return []int{-1, -1}
	}

	first := 0
	last := 0
	for i := 0; i < size; i++ {
		if target == nums[i] {
			// first time we update both first & last
			if nums[first] != target {
				first = i
				last = i
				// we update last index of subsequent findings
			} else {
				last = i
			}
		}
	}

	// we doble check we find the target
	if nums[first] == target {
		return []int{first, last}
	}

	return []int{-1, -1}
}

func main() {
	arr := []int{1}
	fmt.Println(searchRange(arr, 1))
}
