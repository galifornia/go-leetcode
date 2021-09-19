package main

import (
	"math"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	permutations := make([][]int, 0)

	sort.Ints(nums)

	permuteRec(nums, []int{}, &permutations)

	return permutations
}

func permuteRec(nums, seq []int, permutations *[][]int) {
	if len(seq) == len(nums) {
		validSolution := make([]int, len(seq))
		copy(validSolution, seq)
		*permutations = append(*permutations, validSolution)
		return
	}

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		prev := math.MaxInt32
		if i > 0 {
			prev = nums[i-1]
		}

		if num == math.MinInt32 || num == prev {
			continue
		}

		nums[i] = math.MinInt32
		seq = append(seq, num)
		permuteRec(nums, seq, permutations)
		seq = seq[:len(seq)-1]
		nums[i] = num
	}
}

// func main() {
// 	arr := []int{3, 3, 0, 3}
// 	fmt.Println(permuteUnique(arr))
// }
