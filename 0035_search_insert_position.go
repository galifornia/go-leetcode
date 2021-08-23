package main

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		pivot := left + (right-left+1)/2

		if nums[pivot] == target {
			return pivot
		}

		if nums[pivot] > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	return left + (right-left+1)/2
}

// func main() {
// 	arr := []int{1, 3, 5, 6}
// 	fmt.Println(searchInsert(arr, 4))
// }
