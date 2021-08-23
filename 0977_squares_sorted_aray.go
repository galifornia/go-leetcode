package main

func sortedSquares(nums []int) []int {
	size := len(nums)
	left := 0
	right := size - 1
	newArr := make([]int, size)
	i := size - 1

	for left <= right {
		if nums[left]*nums[left] <= nums[right]*nums[right] {
			newArr[i] = nums[right] * nums[right]
			right--
		} else {
			newArr[i] = nums[left] * nums[left]
			left++
		}
		i--
	}

	return newArr
}

// func main() {
// 	// arr := []int{-7, -3, 2, 3, 11}
// 	// arr := []int{-4, -1, 0, 3, 10}
// 	arr := []int{-5, -3, -2, -1}
// 	fmt.Println(sortedSquares((arr)))

// }
