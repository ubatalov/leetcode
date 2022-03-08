package medium

func MinPartitions(n string) int {
	var max int32
	for _, c := range n {
		if c > max {
			max = c
		}
	}
	return int(max - '0')
}
