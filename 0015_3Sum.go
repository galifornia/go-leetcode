package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	size := len(nums)
	if size < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	result := make([][]int, 0)

	for i := 0; i+2 < size; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, size-1
		target := -nums[i]

		for j < k {
			sum := nums[j] + nums[k]
			if sum == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j-1] == nums[j] {
					j++
				}
				for j < k && nums[k+1] == nums[k] {
					k--
				}
			} else if nums[j]+nums[k] > target {
				k--
			} else {
				j++
			}
		}
	}

	return result
}

func main() {
	// arr := []int{-1, 0, 1, 2, -1, -4}
	arr := []int{3, 0, -2, -1, 1, 2}
	fmt.Println(threeSum(arr))
}
