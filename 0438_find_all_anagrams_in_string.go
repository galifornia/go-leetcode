package main

func findAnagrams(s string, p string) []int {
	window := make(map[byte]int)
	objective := createCharMap(p)
	indexes := make([]int, 0)

	k := 0
	w := len(p) - 1
	for w < len(s) {
		slice := createCharMap(s[k : w+1])
		if isAnagram(slice, objective) {
			indexes = append(indexes, k)
		}
		window[s[w]]++
		if k > 0 {
			window[s[k-1]]--
		}
		k++
		w++
	}
	return indexes
}

func isAnagram(tentative map[byte]int, objetive map[byte]int) bool {
	for key := range tentative {
		if tentative[key] != objetive[key] {
			return false
		}
	}
	return true
}

func createCharMap(s string) map[byte]int {
	charMap := make(map[byte]int, 0)
	for i := range s {
		charMap[s[i]]++
	}
	return charMap
}

// func main() {
// 	s := "cbaebabacd"
// 	p := "abc"
// 	fmt.Println(findAnagrams(s, p))
// }
