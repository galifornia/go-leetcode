package main

func twoSumII(numbers []int, target int) []int {
	size := len(numbers)
	left := 0
	right := size - 1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left, right}
		}

		if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}

	return nil
}

// func main() {
// 	arr := []int{-4, 2, 3, 4, 11}
// 	// arr := []int{-1, 0}
// 	target := 7

// 	fmt.Println(twoSumII(arr, target))
// }
