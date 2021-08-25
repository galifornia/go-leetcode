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

	// initialize first 2 values in map
	memory := make(map[int]int)
	memory[0] = 1
	memory[1] = 1

	// build each new value from the previous two numbers already calculated in the memory map
	for i := 2; i < n; i++ {
		memory[i] = memory[i-1] + memory[i-2]
	}

	return memory[n-1]
}

func main() {
	fmt.Println(fib(5))
}
