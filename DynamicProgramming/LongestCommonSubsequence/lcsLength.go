package lcs

import "fmt"

func lcsLength(x, y string) int {
	lx, ly := len(x), len(y)
	b, c := make([][]byte, lx), make([][]int, lx+1)

	for i := 0; i < lx; i++ {
		b[i] = make([]byte, ly)
	}

	for i := 0; i < lx+1; i++ {
		c[i] = make([]int, ly+1)
	}

	for i := 0; i < lx; i++ {
		for j := 0; j < ly; j++ {
			if x[i] == y[j] {
				c[i+1][j+1] = c[i][j] + 1
				b[i][j] = '\\'
			} else if c[i][j+1] >= c[i+1][j] {
				c[i+1][j+1] = c[i][j+1]
				b[i][j] = '^'
			} else {
				c[i+1][j+1] = c[i+1][j]
				b[i][j] = '<'
			}
		}
	}

	return c[lx][ly]
}

func lcsLength1(x, y string) int {
	lx, ly := len(x), len(y)
	c1, c2 := make([]int, lx+1), make([]int, lx+1)

	for i := 0; i < lx; i++ {
		for j := 0; j < ly; j++ {
			if x[i] == y[j] {
				c2[j+1] = c1[j] + 1
			} else if c2[j] >= c1[j+1] {
				c2[j+1] = c2[j]
			} else {
				c2[j+1] = c1[j+1]
			}
		}
		c1, c2 = c2, c1

	}
	return c1[ly]
}

func printLCS(b [][]byte, X string) {
	i, j := len(b), len(X)
	if i == 0 || j == 0 {
		return
	}

	if b[i][j] == '\\' {
		printLCS(b[:][:i-1], X[:j-1])
		fmt.Printf("\\")
	} else if b[i][j] == '^' {
		printLCS(b[:][:i-1], X[:j-1])
	} else {
		printLCS(b[:][:], X[:j-1])
	}
}
