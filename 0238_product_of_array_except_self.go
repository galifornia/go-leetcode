package main

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	aux := make([]int, len(nums))

	size := len(nums)

	answer[0] = 1
	for i := 1; i < size; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	aux[size-1] = 1
	for j := size - 2; j >= 0; j-- {
		aux[j] = aux[j+1] * nums[j+1]
	}

	for i := range nums {
		answer[i] *= aux[i]
	}

	return answer
}

// func main() {
// 	arr := []int{1, 2, 3, 4}
// 	fmt.Println(productExceptSelf(arr))
// }
