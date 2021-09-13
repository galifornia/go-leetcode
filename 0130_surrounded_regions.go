package main

func solve(board [][]byte) {
	size := len(board)
	if size < 2 {
		return
	}

	isNotSurroundedMap := make([][]bool, size)
	for i := range isNotSurroundedMap {
		isNotSurroundedMap[i] = make([]bool, len(board[0]))
	}

	// check first & last row
	for i := 0; i < size; i += size - 1 {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, isNotSurroundedMap)
		}
	}

	// check first & las column
	for j := 0; j < len(board[0]); j += len(board[0]) - 1 {
		for i := 0; i < size; i++ {
			dfs(board, i, j, isNotSurroundedMap)
		}
	}

	// modify board accordinly for regions that are not reached from border
	for i := 0; i < size; i++ {
		for j := 0; j < len(board[0]); j++ {
			if !isNotSurroundedMap[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(grid [][]byte, i, j int, isNotSurroundedMap [][]bool) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 'O' || isNotSurroundedMap[i][j] {
		return
	}

	isNotSurroundedMap[i][j] = true

	dfs(grid, i-1, j, isNotSurroundedMap)
	dfs(grid, i+1, j, isNotSurroundedMap)
	dfs(grid, i, j-1, isNotSurroundedMap)
	dfs(grid, i, j+1, isNotSurroundedMap)
}

// func main() {
// 	// board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
// 	// board := [][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}}
// 	board := [][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}, {'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}}
// 	fmt.Println(board)
// 	solve(board)
// 	fmt.Println(board)
// }
