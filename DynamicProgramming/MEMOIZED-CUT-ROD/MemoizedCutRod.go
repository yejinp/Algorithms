package memoizedCutRod

import (
	"math"
)

func MemoizedCutRod(p []int) int {
	r := make([]int, len(p)+1)

	for i := 0; i < len(r); i++ {
		r[i] = math.MinInt
	}

	return MemoizedCutRodAux(p, r)
}

func MemoizedCutRodAux(p, r []int) int {
	n := len(p)
	if r[n] >= 0 {
		return r[n]
	}
	q := math.MinInt
	if n == 0 {
		q = 0
	} else {
		for i := 0; i < n; i++ {
			q = max(q, p[i]+MemoizedCutRodAux(p[:n-i-1], r))
		}
	}
	r[n] = q
	return q
}

func BottomUpCutRod(p []int) int {
	r := make([]int, len(p)+1)
	r[0] = 0

	for j := 1; j <= len(p); j++ {
		q := math.MinInt
		for i := 0; i < j; i++ {
			q = max(q, p[i]+r[j-i-1])
		}
		r[j] = q
	}

	return r[len(p)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
