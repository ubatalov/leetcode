package easy

func SingleNumber(nums []int) int {
	var a int
	for _, n := range nums {
		a = a ^ n
	}
	return a
}
