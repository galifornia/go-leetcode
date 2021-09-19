package main

func combinationSum(candidates []int, target int) [][]int {
	results := make([][]int, 0)
	countMap := make(map[int]int, len(candidates))

	combinationRec(candidates, target, []int{}, countMap, 0, &results)

	return results
}

func combinationRec(candidates []int, target int, seq []int, countMap map[int]int, sum int, results *[][]int) {
	if sum == target {
		for i := range *results {
			if areMapsCountEqual(countMap, (*results)[i]) {
				return
			}
		}
		result := make([]int, len(seq))
		copy(result, seq)
		*results = append(*results, result)
		return
	}

	if sum > target {
		return
	}

	for _, num := range candidates {
		seq = append(seq, num)
		countMap[num]++
		combinationRec(candidates, target, seq, countMap, sum+num, results)
		countMap[num]--
		seq = seq[:len(seq)-1]
	}
}

func areMapsCountEqual(count1 map[int]int, arr []int) bool {
	count2 := make(map[int]int, len(count1))

	for _, num := range arr {
		count2[num]++
	}

	for k := range count2 {
		if count1[k] != count2[k] {
			return false
		}
	}

	return true
}

// func main() {
// 	candidates := []int{2, 3, 6, 7}
// 	target := 7

// 	fmt.Println(combinationSum(candidates, target))
// }
