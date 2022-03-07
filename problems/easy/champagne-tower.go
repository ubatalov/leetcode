package easy

import "fmt"

func ChampagneTower(poured, queryRow, queryGlass int) float64 {
	var cur []float64
	pre := []float64{float64(poured)}
	for i, added := 1, 0.0; i < queryRow+1; i++ {
		cur = make([]float64, i+1)
		for j := 0; j < i; j++ {
			if pre[j] > 1 {
				added = (pre[j] - 1) / 2
				cur[j] += added
				cur[j+1] += added
				fmt.Println(cur)
			}
		}
		pre = cur
	}

	if pre[queryGlass] < 1 {
		return pre[queryGlass]
	} else {
		return 1
	}
}
