package main

type Visit struct {
	row   int
	col   int
	depth int
}

var (
	directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{-1, 1},
		{1, -1},
		{-1, -1},
	}
)

func shortestPathBinaryMatrix(grid [][]int) int {
	size := len(grid)
	if grid[0][0] != 0 || grid[size-1][size-1] != 0 {
		return -1
	}
	queue := make([]*Visit, 0)
	queue = append(queue, &Visit{row: 0, col: 0, depth: 1})
	grid[0][0] = 1

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		row, col, depth := item.row, item.col, item.depth

		if row == size-1 && col == size-1 {
			return depth
		}

		for _, dir := range directions {
			x := row + dir[0]
			y := col + dir[1]
			if x >= 0 && x < size && y >= 0 && y < size && grid[x][y] == 0 {
				grid[x][y] = 1
				queue = append(queue, &Visit{row: x, col: y, depth: depth + 1})
			}
		}
	}

	return -1
}

// func main() {
// 	grid := [][]int{{0, 1}, {1, 0}}
// 	fmt.Println(shortestPathBinaryMatrix(grid))
// }
