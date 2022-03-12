package medium

func CountPoints(points [][]int, queries [][]int) []int {
	pow := func(x int) int {
		return x * x
	}

	var res []int
	for _, query := range queries {
		i, max := 0, pow(query[2])
		for _, point := range points {
			if pow(point[0]-query[0])+pow(point[1]-query[1]) <= max {
				i++
			}
		}
		res = append(res, i)
	}
	return res
}
