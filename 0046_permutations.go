package main

func permute(nums []int) [][]int {
	var res [][]int
	permuteRec([]int{}, nums, &res)
	return res
}

// We use a pointer for the result so we don't need to worry returning it.
func permuteRec(currComb, left []int, res *[][]int) {
	// We know that we found a new combination when we have no elements left.
	if len(left) == 0 {
		*res = append(*res, currComb)
		return
	}
	// For the next iteration we consider all the left elements but the current one (idx).
	for idx, l := range left {
		permuteRec(
			append(currComb, l),
			append(append([]int{}, left[:idx]...), left[idx+1:]...), // Make sure to allocate a new slice.
			res,
		)
	}
}

// func main() {
// 	nums := []int{1, 2, 3}
// 	fmt.Println(permute(nums))
// }
