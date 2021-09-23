package main

func exist(board [][]byte, word string) bool {
	containedInBoard := false

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if exists(board, i, j, word, 0) {
				return true
			}
		}
	}

	return containedInBoard
}

func exists(board [][]byte, i, j int, word string, idx int) bool {
	if idx == len(word) {
		return true
	}

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}

	if board[i][j] != word[idx] {
		return false
	}

	aux := board[i][j]
	board[i][j] = '*'
	present := exists(board, i+1, j, word, idx+1) || exists(board, i, j-1, word, idx+1) || exists(board, i-1, j, word, idx+1) || exists(board, i, j+1, word, idx+1)
	board[i][j] = aux

	return present
}

// func main() {
// 	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
// 	word := "ABCCED"

// 	fmt.Println(exist(board, word))
// }
