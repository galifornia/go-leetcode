package main

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}

	sum := 1
	i := 1
	for {
		sum *= 2
		if sum == n {
			return true
		} else if sum > n {
			return false
		}
		i++
	}
}
