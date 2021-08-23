package main

func buddyStrings(s string, goal string) bool {
	size := len(s)
	if size != len(goal) {
		return false
	}
	count := 0
	hasSameCharDifferentPositions := false
	countMap := make(map[byte]int)

	for i := 0; i < size; i++ {
		countMap[s[i]]++
		if countMap[s[i]] > 1 {
			hasSameCharDifferentPositions = true
		}

		if s[i] != goal[i] {
			count++
			if count > 2 {
				return false
			}
		}
	}

	for k := 0; k < size; k++ {
		countMap[goal[k]]--
	}

	containsSameCharacters := true
	for _, v := range countMap {
		if v != 0 {
			containsSameCharacters = false
			break
		}
	}

	return containsSameCharacters && count == 2 || (containsSameCharacters && hasSameCharDifferentPositions)
}

// func main() {
// 	s := "abcaa"
// 	goal := "abcbb"
// 	fmt.Println(buddyStrings(s, goal))
// }
