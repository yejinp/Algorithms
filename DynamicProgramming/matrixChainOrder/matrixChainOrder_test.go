package matrixchainorder

import (
	"fmt"
	"testing"
)

func TestMatrixChainOrder(t *testing.T) {
	tests := []struct {
		input []int
		wantM [][]int
		wantS [][]int
	}{{[]int{30, 35, 15, 5, 10, 20, 25}, [][]int{}, [][]int{}}}

	for _, test := range tests {
		getM, getS := MatrixChainOrder(test.input)
		fmt.Println(getM, getS)
		/*
			if equal2demi(test.wantM, getM) || equal2demi(test.wantS, getS) {
				t.Errorf("MatrixChainOrder(%v), got:(%v,%v), want:(%v,%v)",
					test.input, getM, getS, test.wantM, test.wantS)
			}
		*/
	}

}

func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func equal2demi(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if equal(v, b[i]) == false {
			return false
		}
	}
	return true
}
