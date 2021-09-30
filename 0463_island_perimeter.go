package main

func islandPerimeter(grid [][]int) int {
	perimeter := 0

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				dfs(grid, i, j, &perimeter)
				return perimeter
			}
		}
	}

	return perimeter
}

func dfs(grid [][]int, row, col int, perimeter *int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] != 1 {
		return
	}

	grid[row][col] = 2

	addToPerimeter := 4 //default

	if row+1 < len(grid) && grid[row+1][col] != 0 {
		addToPerimeter--
	}

	if row > 0 && grid[row-1][col] != 0 {
		addToPerimeter--
	}

	if col > 0 && grid[row][col-1] != 0 {
		addToPerimeter--
	}

	if col+1 < len(grid[0]) && grid[row][col+1] != 0 {
		addToPerimeter--
	}

	*perimeter += addToPerimeter

	dfs(grid, row+1, col, perimeter)
	dfs(grid, row-1, col, perimeter)
	dfs(grid, row, col-1, perimeter)
	dfs(grid, row, col+1, perimeter)
}
