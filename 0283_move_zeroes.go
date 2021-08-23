package main

import "fmt"

func moveZeroes(nums []int) {
	size := len(nums)
	if size < 2 {
		return
	}

	left := 0
	count := 0

	// move 0 to end of list
	for left < size-count {

		if nums[left] == 0 {
			count++
			for i := left; i < size-1; i++ {
				nums[i] = nums[i+1]
			}
			nums[size-1] = 0
			left--
		}
		left++
	}
}

func main() {
	arr := []int{0, 1, 0, 0, 3, 0, 12}
	// arr := []int{0, 0, 1}
	moveZeroes(arr)
	fmt.Println(arr)
}
