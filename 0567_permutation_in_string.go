package main

func checkInclusion(s1 string, s2 string) bool {
	// Return if s1 is smaller
	if len(s1) > len(s2) {
		return false

	}

	size1 := len(s1)
	size2 := len(s2)

	// create maps for s1 (permutation string) and s2
	countMap1 := make(map[byte]int)
	countMap2 := make(map[byte]int)

	// initialize map counting characters in both s1 & first s2 window
	for i := 0; i < len(s1); i++ {
		countMap1[s1[i]]++
		countMap2[s2[i]]++
	}

	// check if we are lucky and our first window matches
	if matches(&countMap1, &countMap2) {
		return true
	}

	i, j := 1, size1
	// loop until window is at the end of string s2
	for j < size2 {
		// we move the window 1 so we remove the character that gets out of the window
		countMap2[s2[i-1]]--
		countMap2[s2[j]]++

		if matches(&countMap1, &countMap2) {
			return true
		}

		i++
		j++
	}

	return false
}

func matches(m1, m2 *map[byte]int) bool {
	for k, v := range *m1 {
		if v != (*m2)[k] {
			return false
		}
	}
	return true
}

// func main() {
// 	// s1 := "hello"
// 	// s2 := "ooolleoooleh"
// 	// s1 := "adc"
// 	// s2 := "dcda"
// 	s2 := "eidbaooo"
// 	s1 := "ab"
// 	fmt.Println(checkInclusion(s1, s2))
// }
