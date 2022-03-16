package easy

import "sort"

func MinimumSum(num int) int {
	numbers := []int{0, 0, 0, 0}
	for i := 0; num > 0; i++ {
		numbers[i] = num % 10
		num /= 10
	}
	sort.Ints(numbers)

	return numbers[0]*10 + numbers[2] + numbers[1]*10 + numbers[3]
}
