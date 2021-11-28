package easy

import "strconv"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}

	num := strconv.Itoa(x)
	end := len(num) - 1

	for i := 0; i <= len(num)/2; i++ {
		if num[i] != num[end] {
			return false
		}

		end--
	}

	return true
}
