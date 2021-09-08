package main

func backspaceCompare(s string, t string) bool {
	s1 := backspaceString(s)
	s2 := backspaceString(t)

	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func backspaceString(s string) string {
	transformedString := make([]byte, 0)

	walk := 0
	back := 0

	for walk < len(s) {
		if s[walk] == '#' {
			back++
		} else {
			for len(transformedString) > 0 && back > 0 {
				transformedString = transformedString[0 : len(transformedString)-1]
				back--
			}
			if len(transformedString) == 0 {
				back = 0
			}
			transformedString = append(transformedString, s[walk])
		}
		walk++
	}
	for len(transformedString) > 0 && back > 0 {
		transformedString = transformedString[0 : len(transformedString)-1]
		back--
	}
	return string(transformedString)
}

// func main() {
// 	s1, s2 := "a##c", "#a#c"
// 	fmt.Println(backspaceCompare(s1, s2))
// }
