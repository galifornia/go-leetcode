package main

type Coordinate struct {
	row    int
	column int
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	m := len(image)
	n := len(image[0])
	queue := []Coordinate{}

	queue = append(queue, Coordinate{row: sr, column: sc})
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}

	for len(queue) > 0 {
		// Take coordinate to do the check
		coordinate := queue[0]
		queue = queue[1:]

		// change color
		image[coordinate.row][coordinate.column] = newColor
		visited[coordinate.row][coordinate.column] = 1

		if coordinate.row-1 >= 0 && image[coordinate.row-1][coordinate.column] == color && visited[coordinate.row-1][coordinate.column] == 0 {
			queue = append(queue, Coordinate{row: coordinate.row - 1, column: coordinate.column})
		}

		if coordinate.row+1 < m && image[coordinate.row+1][coordinate.column] == color && visited[coordinate.row+1][coordinate.column] == 0 {
			queue = append(queue, Coordinate{row: coordinate.row + 1, column: coordinate.column})
		}

		if coordinate.column-1 >= 0 && image[coordinate.row][coordinate.column-1] == color && visited[coordinate.row][coordinate.column-1] == 0 {
			queue = append(queue, Coordinate{row: coordinate.row, column: coordinate.column - 1})
		}

		if coordinate.column+1 < n && image[coordinate.row][coordinate.column+1] == color && visited[coordinate.row][coordinate.column+1] == 0 {
			queue = append(queue, Coordinate{row: coordinate.row, column: coordinate.column + 1})
		}
	}
	return image
}

// func main() {
// 	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
// 	sr := 1
// 	sc := 1
// 	// image := [][]int{{0, 0, 0}, {0, 1, 1}}
// 	// sr := 1
// 	// sc := 1
// 	newColor := 2
// 	fmt.Println(floodFill(image, sr, sc, newColor))
// }
