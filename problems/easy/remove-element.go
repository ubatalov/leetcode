package easy

func RemoveElement(nums []int, val int) int {
	var i int
	for _, n := range nums {
		if n != val {
			nums[i] = n
			i++
		}
	}
	return i
}
