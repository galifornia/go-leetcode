package main

func isPalindrome(x int) bool {
	original := x
	reversed := 0

	if x < 0 {
		return false
	}

	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return original == reversed
}

// func main() {
// 	x := 121
// 	fmt.Println(isPalindrome(x))
// }
