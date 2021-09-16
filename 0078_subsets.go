package main

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{}) // first empty element

	// we loop through each number in array & add it to each subarray contained in result
	for _, num := range nums {
		for _, subarray := range result {
			// if we do not create a new slice weird things happen
			// cannot do append(subarray, num)
			newSubset := append([]int{}, subarray...)
			newSubset = append(newSubset, num)
			result = append(result, newSubset)
		}
	}

	return result
}

// func main() {
// 	arr := []int{9, 0, 3, 5, 7}
// 	fmt.Println(subsets(arr))
// }
