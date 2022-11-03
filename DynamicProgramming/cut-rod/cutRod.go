package cutrod

import (
	"math"
)

func cutRod(p []int) int {
	if len(p) == 0 {
		return 0
	}

	q := math.MinInt

	for i, v := range p {
		tc := cutRod(p[:len(p)-i-1])
		q = max(q, v+tc)
	}

	return q
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
