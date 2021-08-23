package main

import (
	"math"
)

func reverse(x int) int {
	reversedX := 0

	for x != 0 {
		reversedX = reversedX*10 + x%10
		x /= 10
	}

	if reversedX < math.MinInt32 || reversedX > math.MaxInt32 {
		return 0
	}
	return reversedX
}

// func main() {
// 	x := 123
// 	// x := 901000
// 	fmt.Println(reverse(x))
// }
