package easy

func RomanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}

	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	p := m[s[len(s)-1]]

	for i := len(s) - 2; i >= 0; i-- {
		if m[s[i]] < m[s[i+1]] {
			p -= m[s[i]]
		} else {
			p += m[s[i]]
		}
	}

	return p
}
