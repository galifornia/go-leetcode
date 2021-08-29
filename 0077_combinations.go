package main

func combine(n int, k int) [][]int {
	combinations := make([][]int, 0)

	if k > n || k < 0 {
		return combinations
	}

	if k == 0 {
		combinations = append(combinations, make([]int, 0))
		return combinations
	}

	combinations = combine(n-1, k-1)
	for i := range combinations {
		combinations[i] = append(combinations[i], n)
	}

	middleCombinations := combine(n-1, k)
	combinations = append(combinations, middleCombinations...)

	return combinations
}

// func main() {
// 	n := 4
// 	k := 2
// 	fmt.Println(combine(n, k))
// }
