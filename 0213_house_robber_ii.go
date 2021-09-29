package main

// func rob(nums []int) int {
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}

// 	maxReward := math.MinInt32
// 	// rob from start until end - 1
// 	curr := math.MinInt32
// 	prev := math.MinInt32
// 	for i := 0; i < len(nums)-1; i++ {
// 		aux := curr
// 		curr = max(prev+nums[i], curr)
// 		prev = aux
// 	}
// 	maxReward = curr

// 	// rob from start+1 until end
// 	curr = math.MinInt32
// 	prev = math.MinInt32
// 	for i := 1; i < len(nums); i++ {
// 		aux := curr
// 		curr = max(prev+nums[i], curr)
// 		prev = aux
// 	}
// 	maxReward = max(maxReward, curr)

// 	return maxReward
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// func main() {
// 	// arr := []int{1, 2, 3, 1}
// 	// arr := []int{1, 2, 3}
// 	// arr := []int{1, 2, 3, 1, 5}
// 	arr := []int{200, 3, 140, 20, 10}
// 	fmt.Println(rob(arr))
// }
