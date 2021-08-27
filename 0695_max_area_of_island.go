package main

type Point struct {
	x, y int
}

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	max := 0 // result
	// array to avoid visiting same island twice
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	// loop until we find land
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// skip already visited coordinates
			if visited[i][j] || grid[i][j] == 0 {
				continue
			}
			visited[i][j] = true
			// land has been found
			if grid[i][j] == 1 {
				queue := make([]Point, 0)
				queue = append(queue, Point{x: i, y: j})
				islandSize := 0
				for len(queue) > 0 {
					point := queue[0]
					queue = queue[1:]
					visited[point.x][point.y] = true

					islandSize++
					if point.x-1 >= 0 && grid[point.x-1][point.y] == 1 && !visited[point.x-1][point.y] {
						visited[point.x-1][point.y] = true
						queue = append(queue, Point{x: point.x - 1, y: point.y})
					}
					if point.x+1 < m && grid[point.x+1][point.y] == 1 && !visited[point.x+1][point.y] {
						visited[point.x+1][point.y] = true
						queue = append(queue, Point{x: point.x + 1, y: point.y})
					}
					if point.y-1 >= 0 && grid[point.x][point.y-1] == 1 && !visited[point.x][point.y-1] {
						visited[point.x][point.y-1] = true
						queue = append(queue, Point{x: point.x, y: point.y - 1})
					}
					if point.y+1 < n && grid[point.x][point.y+1] == 1 && !visited[point.x][point.y+1] {
						visited[point.x][point.y+1] = true
						queue = append(queue, Point{x: point.x, y: point.y + 1})
					}
				}
				if islandSize > max {
					max = islandSize
				}
			}
		}
	}

	return max
}

// func main() {
// 	// grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
// 	// grid := [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}
// 	grid := [][]int{{0}, {0}}
// 	fmt.Println(maxAreaOfIsland(grid))
// }
