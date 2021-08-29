package main

type Point struct {
	row, col int
}

var (
	directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return -1
	}
	cols := len(grid[0])
	cfresh := 0
	queue := []*Point{}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, &Point{row: i, col: j})
			} else if grid[i][j] == 1 {
				cfresh++
			}
		}
	}

	if cfresh == 0 {
		return 0
	}

	count := 0

	for len(queue) > 0 {
		count++
		size := len(queue)
		for i := 0; i < size; i++ {
			point := queue[0]
			queue = queue[1:]
			for _, dir := range directions {
				row := point.row + dir[0]
				col := point.col + dir[1]

				if row < 0 || col < 0 || row >= rows || col >= cols || grid[row][col] == 0 || grid[row][col] == 2 {
					continue
				}

				grid[row][col] = 2
				queue = append(queue, &Point{row: row, col: col})
				cfresh--
			}
		}
	}

	if cfresh == 0 {
		return count - 1
	} else {
		return -1
	}
}

// func main() {
// 	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
// 	fmt.Println(orangesRotting(grid))
// }
