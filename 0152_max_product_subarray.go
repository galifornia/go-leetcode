package main

func maxProduct(nums []int) int {
	_min, _max, res := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			tmp := _max
			_max = _min
			_min = tmp
		}

		_min = min(nums[i], _min*nums[i])
		_max = max(nums[i], _max*nums[i])

		res = max(res, _max)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// func main() {
// 	// arr := []int{2, 3, -2, 4}
// 	// arr := []int{-2, 0, -1}
// 	arr := []int{3, -1, 4}
// 	fmt.Println(maxProduct(arr))
// }
