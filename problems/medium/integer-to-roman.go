package medium

func IntToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var res string
	i := 0
	for num > 0 {
		if values[i] > num {
			i++
		}
		res += romans[i]
		num -= values[i]
	}

	return res
}
