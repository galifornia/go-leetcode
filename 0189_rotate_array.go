package main

func rotate(nums []int, k int) {
	size := len(nums)
	k = k % size
	copy(nums, append(nums[size-k:size], nums[:size-k]...))
}

// func main() {
// 	// arr := []int{1, 2, 3, 4, 5, 6, 7}
// 	arr := []int{-1, -100, 3, 99}
// 	rotate(arr, 1)
// 	fmt.Println(arr)
// }
