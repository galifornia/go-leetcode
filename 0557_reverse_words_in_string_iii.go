package main

func reverseWords(s string) string {
	phrase := []byte(s)
	size := len(phrase)
	if size < 2 {
		return s
	}

	left := 0
	right := 0

	for left < size {
		if right+1 >= size {
			break
		}
		for right+1 < size && s[right+1] != ' ' {
			right++
		}
		for k := right; k > left; k-- {
			if right <= left {
				break
			}
			aux := s[k]
			phrase[k] = s[left]
			phrase[left] = aux
			left++
		}
		right += 2 // jump over space
		left = right
	}

	return string(phrase)
}

// func main() {
// 	// s := "Let's take LeetCode contest"
// 	s := "God Ding"
// 	fmt.Println(reverseWords(s))
// }
