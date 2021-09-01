package main

import (
	"strings"
)

func letterCasePermutation(s string) []string {
	output := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if !CheckIfNumber(s[i]) {
			permString := make([]byte, len(s))
			for idx := range permString {
				permString[idx] = s[idx]
			}

			if strings.ToLower(string(s[i]))[0] == s[i] {
				permString[i] = strings.ToUpper(string(s[i]))[0]
			} else {
				permString[i] = strings.ToLower(string(s[i]))[0]
			}

			output = append(output, string(permString))
		}
	}
	return output
}

func CheckIfNumber(r byte) bool {

	if r >= 48 && r <= 57 {
		return false
	}

	return true
}

// func main() {
// 	s := "a1b2"
// 	fmt.Println(letterCasePermutation(s))
// }
