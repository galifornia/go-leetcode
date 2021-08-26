package main

func checkInclusion(s1 string, s2 string) bool {
	// Return if s1 is smaller
	if len(s1) > len(s2) {
		return false
	}

	// map that contains number of times of each character in s1
	runeMap := make(map[rune]int)
	for _, c := range s1 {
		runeMap[c]++
	}

	// loop through s2 using a sliding window of size len(s1)
	for i := 0; i < len(s2)-len(s1)+1; i++ {
		// create count map for this sliding window
		workingMap := make(map[rune]int)
		for j := i; j < i+len(s1); j++ {
			workingMap[rune(s2[j])]++
		}
		// Check if there is match between maps
		ok := matches(&runeMap, &workingMap)
		if ok {
			return true
		}
	}

	return false
}

func matches(m1, m2 *map[rune]int) bool {
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
// 	s1 := "eidboaoo"
// 	s2 := "ab"
// 	fmt.Println(checkInclusion(s1, s2))
// }
