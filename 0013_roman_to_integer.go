package main

func romanToInt(s string) int {
	num := 0
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	size := len(s)
	for i := 0; i < size; i++ {
		if i+1 < size && s[i] == 'I' {
			if s[i+1] == 'X' {
				num += 9
				i++
				continue
			} else if s[i+1] == 'V' {
				num += 4
				i++
				continue
			}
		} else if i+1 < size && s[i] == 'X' {
			if s[i+1] == 'L' {
				num += 40
				i++
				continue
			} else if s[i+1] == 'C' {
				num += 90
				i++
				continue
			}

		} else if i+1 < size && s[i] == 'C' {
			if s[i+1] == 'D' {
				num += 400
				i++
				continue
			} else if s[i+1] == 'M' {
				num += 900
				i++
				continue
			}
		}
		num += values[s[i]]
	}

	return num
}
