package main

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	// create and initialize value map
	r := make(map[int]int)
	r[1] = 1
	r[2] = 2
	r[3] = 3

	for i := 4; i <= n; i++ {
		r[i] = r[i-1] + r[i-2]
	}
	return r[n]
}

// func main() {
// 	n := 7
// 	fmt.Println(climbStairs(n))
// }
