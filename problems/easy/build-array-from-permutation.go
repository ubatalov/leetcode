package easy

func BuildArray(nums []int) []int {
	res := make([]int, len(nums))
	for i, pos := range nums {
		res[i] = nums[pos]
	}
	return res
}
