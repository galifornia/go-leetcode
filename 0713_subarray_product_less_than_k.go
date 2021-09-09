package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 0 {
		return 0
	}

	count, product, i := 0, 1, 0
	for j := 0; j < len(nums); j++ {
		product *= nums[j]
		for i <= j && product >= k {
			product /= nums[i]
			i++
		}
		count += j - i + 1
	}
	return count
}

// func main() {
// 	arr := []int{10, 5, 2, 6}
// 	k := 100

// 	fmt.Println(numSubarrayProductLessThanK(arr, k))
// }
