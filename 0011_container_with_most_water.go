package main

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	water := 0

	for i < j {
		water = max(water, (j-i)*min(height[i], height[j]))

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return water
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
