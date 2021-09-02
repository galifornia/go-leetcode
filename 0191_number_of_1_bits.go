package main

func hammingWeight(num uint32) int {
	count := 0
	// loop until we remove all 1
	for num != 0 {
		// XOR original number with number - 1 which changes rightmost 1 to 0
		num &= num - 1
		count++
	}

	return count
}
