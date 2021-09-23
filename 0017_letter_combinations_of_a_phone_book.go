package main

import "fmt"

func letterCombinations(digits string) []string {
	letters := make(map[byte]*[]byte)
	letters['2'] = &[]byte{'a', 'b', 'c'}
	letters['3'] = &[]byte{'d', 'e', 'f'}
	letters['4'] = &[]byte{'g', 'h', 'i'}
	letters['5'] = &[]byte{'j', 'k', 'l'}
	letters['6'] = &[]byte{'m', 'n', 'o'}
	letters['7'] = &[]byte{'p', 'q', 'r', 's'}
	letters['8'] = &[]byte{'t', 'u', 'v'}
	letters['9'] = &[]byte{'w', 'x', 'y', 'z'}

	solution := []string{}

	if digits == "" {
		return solution
	}

	recursiveLetterCombinations(digits, letters, "", 0, &solution)

	return solution

}

func recursiveLetterCombinations(digits string, letters map[byte]*[]byte, current string, i int, solution *[]string) {
	if i == len(digits) {
		*solution = append(*solution, current)
		return
	}

	for _, ch := range *letters[digits[i]] {
		recursiveLetterCombinations(digits, letters, current+string(ch), i+1, solution)
	}

}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
