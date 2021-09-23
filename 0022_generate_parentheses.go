package main

func generateParenthesis(n int) []string {
	solution := []string{}

	generate(n, &solution, []byte{}, 0)

	return solution
}

func generate(n int, solution *[]string, current []byte, pos int) {
	if len(current) == 2*n {
		if valid(current) {
			*solution = append(*solution, string(current))
		}
		return
	}
	current = append(current, '(')
	generate(n, solution, current, pos+1)
	current = current[:len(current)-1]
	current = append(current, ')')
	generate(n, solution, current, pos+1)
}

func valid(current []byte) bool {
	balance := 0
	for _, c := range current {
		if c == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
	}

	return balance == 0
}

// func main() {
// 	fmt.Println(generateParenthesis(3))
// }
