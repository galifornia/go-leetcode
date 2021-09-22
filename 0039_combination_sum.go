package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	// solution array
	combinations := make([][]int, 0)
	// we sort the list of integers
	sort.Ints(candidates)

	// backtracking
	backtrack(&candidates, &combinations, []int{}, target, 0)

	return combinations
}

func backtrack(candidates *[]int, combinations *[][]int, currentSum []int, remain int, start int) {
	// if too big return
	if remain < 0 {
		return
	}

	// if we have the target add it to combinations & return
	if remain == 0 {
		newComb := make([]int, 0)
		newComb = append(newComb, currentSum...)
		*combinations = append((*combinations), newComb)
		return
	}

	// walk through all remaining candidates & use recusion
	for i := start; i < len(*candidates); i++ {
		currentSum = append(currentSum, (*candidates)[i])
		backtrack(candidates, combinations, currentSum, remain-(*candidates)[i], i)
		currentSum = currentSum[:len(currentSum)-1]
	}
}

// func main() {
// 	candidates := []int{2, 3, 6, 7}
// 	target := 7

// 	fmt.Println(combinationSum(candidates, target))
// }
