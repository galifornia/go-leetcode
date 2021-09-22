package main

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	return genCombinations(candidates, target)
}

func genCombinations(candidates []int, target int) [][]int {
	solution := make([][]int, 0)
	for i, candidate := range candidates {
		// we skip the first & the same consecutive candidates
		if i != 0 && candidates[i] == candidates[i-1] {
			continue
		}
		if target == candidate {
			return append(solution, []int{target})
		} else if target < candidate {
			return solution
		}
		tmp := genCombinations(candidates[i+1:], target-candidate)
		for _, val := range tmp {
			solution = append(solution, append(val, candidate))
		}
	}
	return solution
}

// func main() {
// 	candidates := []int{10, 1, 2, 7, 6, 1, 5}
// 	target := 8
// 	fmt.Println(combinationSum2(candidates, target))
// }
