package main

func twoSum(nums []int, target int) []int {
	size := len(nums)
	m := make(map[int]int, size)

	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{i, j}
		}
		m[n] = i
	}
	return []int{}
}

// func main() {
// 	arr := []int{3, 2, 4}
// 	fmt.Println(twoSum(arr, 6))
// }
