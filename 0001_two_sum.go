package main

func twoSum(nums []int, target int) []int {
	size := len(nums)
	for i := 0; i < size; i++ {
		for k := 0; k < size; k++ {
			if k != i && nums[i]+nums[k] == target {
				return []int{i, k}
			}
		}
	}
	return nil
}

// func main() {
// 	arr := []int{3, 2, 4}
// 	fmt.Println(twoSum(arr, 6))
// }
