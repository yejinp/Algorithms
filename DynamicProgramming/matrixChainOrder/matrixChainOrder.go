package matrixchainorder

import (
	"math"
)

func MatrixChainOrder(p []int) ([][]int, [][]int) {
	n := len(p)
	m := make([][]int, n+1)
	s := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		m[i] = make([]int, n+1)
		s[i] = make([]int, n+1)
	}

	for l := 2; l <= n; l++ { //chain length l,  l is the chain length
		for i := 1; i < n-l+1; i++ { // chain begins at Ai
			j := i + l - 1 // chain ends at Aj
			m[i][j] = math.MaxInt
			for k := i; k <= j-1; k++ { // try Ai:k A(k+1):j
				q := m[i][k] + m[k+1][j] + p[i-1]*p[k]*p[j]
				if q < m[i][j] {
					m[i][j] = q // remember this cost
					s[i][j] = k // remember this index
				}
			}
		}
	}

	return m, s
}
