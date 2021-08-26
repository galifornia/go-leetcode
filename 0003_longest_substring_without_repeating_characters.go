package main

func lengthOfLongestSubstring(s string) int {
	size := len(s)
	longest := 0
	memory := make(map[rune]int)
	streak := 0

	// loop through input string with index & value
	for i, c := range s {
		// If character is repeated (present in map) && index of previous appeareance
		// is bigger or equal than current streak (which holds index of last non repeating character)
		// we increase streak by 1
		if _, okay := memory[c]; okay && memory[c] >= streak {
			// if current streak is bigger than longest update
			if i-streak > longest {
				longest = i - streak
			}
			streak = memory[c] + 1
		}
		// update latest index for this particular c character
		memory[c] = i
	}
	if size-streak > longest {
		// if current streak is bigger than longest update
		longest = size - streak
	}

	return longest
}

// func main() {
// 	s := "pwwkew"
// 	fmt.Println(lengthOfLongestSubstring(s))
// }
