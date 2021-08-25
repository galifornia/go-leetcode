package main

import "fmt"

func fib(n int) int {
	// return error when n is negative
	if n < 0 {
		return -1
	}

	// special 0 case
	if n == 0 {
		return 0
	}

	// initial cases that generate the sucession (1 & 2)
	if n < 3 {
		return 1
	}

	// initialize two pair of values from n = 3
	prev := 1
	last := 2

	// update last with the sum of the two and remember the last one
	for i := 2; i < n; i++ {
		aux := prev
		prev = last
		last = aux + prev
	}

	return last
}

// version using a map to store calculations

// func fib(n int) int {
// 	if n < 0 {
// 		return -1
// 	}

//     if n == 0 {
// 		return 0
// 	}

// 	if n < 3 {
// 		return 1
// 	}

// 	memory := make(map[int]int)
// 	memory[0] = 1
// 	memory[1] = 1

// 	for i := 2; i < n; i++ {
// 		memory[i] = memory[i-1] + memory[i-2]
// 	}

// 	return memory[n-1]
// }

func main() {
	fmt.Println(fib(5))
}
