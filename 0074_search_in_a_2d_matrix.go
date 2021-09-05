package main

func searchMatrix(matrix [][]int, target int) bool {
	cols := len(matrix[0])
	if cols == 0 {
		return false
	}

	row := 0
	// find right row to search
	for row < len(matrix) && target > matrix[row][cols-1] {
		row++
	}

	// exit if target is not in any row
	if row == len(matrix) {
		return false
	}

	// do binary search on the row array
	left := 0
	right := len(matrix[0]) - 1

	for left <= right {
		pivot := left + (right-left+1)/2

		if matrix[row][pivot] == target {
			return true
		}

		if matrix[row][pivot] > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	return false
}
