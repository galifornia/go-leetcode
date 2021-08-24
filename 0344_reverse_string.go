package main

func reverseString(s []byte) {
	size := len(s)
	left := 0
	right := size - 1

	for left < right {
		aux := s[right]
		s[right] = s[left]
		s[left] = aux
		left++
		right--
	}
}

// func main() {
// 	s := []byte{'h', 'e', 'l', 'l', 'o'}
// 	// ["o","l","l","e","h"]
// 	fmt.Println(s)
// 	reverseString(s)
// }
