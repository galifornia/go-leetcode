package main

func maxAreaOfIsland(grid [][]int) int {
	max := 0 // result

	// loop until we find land
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// skip already visited coordinates
			if grid[i][j] == 1 {
				islandSize := dfs(grid, i, j)
				if islandSize > max {
					max = islandSize
				}
			}
		}
	}

	return max
}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0

	islandSize := 1
	islandSize += dfs(grid, i-1, j)
	islandSize += dfs(grid, i+1, j)
	islandSize += dfs(grid, i, j-1)
	islandSize += dfs(grid, i, j+1)
	return islandSize
}

// func main() {
// 	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
// 	// grid := [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}
// 	// grid := [][]int{{0}, {0}}
// 	fmt.Println(maxAreaOfIsland(grid))
// }
