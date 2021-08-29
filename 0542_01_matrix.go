package main

import (
	"math"
)

func updateMatrix(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])
	// quick exit
	if rows == 0 {
		return mat
	}
	dist := make([][]int, rows)
	// initialize distance matrix to max integers
	for i := 0; i < rows; i++ {
		dist[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			dist[i][j] = math.MaxInt32
		}
	}

	// We need to do two passes, in this first one we go from left to right and top to bottom
	// checking the previous minimum distance from the previous row (i) and column (j)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				dist[i][j] = 0
			} else {
				if i > 0 {
					// check if current position or go left to previous walked row (i-1) plus 1 and keep min distance
					dist[i][j] = minDistance(dist[i][j], dist[i-1][j]+1)
				}

				if j > 0 {
					// check if current position or go top to previous walked row (j-1) plus 1 and keep min distance
					dist[i][j] = minDistance(dist[i][j], dist[i][j-1]+1)
				}
			}

		}
	}

	// In the second we walk the matrix from bottom-right to top-left
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if i < rows-1 {
				// check if current position or go right to previous walked row (i+1) plus 1 has less distance
				dist[i][j] = minDistance(dist[i][j], dist[i+1][j]+1)
			}

			if j < cols-1 {
				// check if current position or go bottom to previous walked column(j+1) plus 1 has less distance
				dist[i][j] = minDistance(dist[i][j], dist[i][j+1]+1)
			}
		}
	}
	return dist
}

func minDistance(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

// func main() {
// 	arr := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
// 	fmt.Println(updateMatrix(arr))
// }
