package easy

func IsValid(s string) bool {
	stack, m := make([]rune, 0, 1), map[rune]rune{'{': '}', '(': ')', '[': ']'}

	for _, r := range s {
		switch r {
		case '[', '(', '{':
			stack = append(stack, m[r])
		case '}', ']', ')':
			if len(stack) == 0 || stack[len(stack)-1] != r {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
