package mySort

import "testing"

func intCmp(a, b interface{}) int {
	return a.(int) - b.(int)
}

func equal(a []interface{}, b []int) bool {
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

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input []interface{}
		want  []int
	}{
		{[]interface{}{6, 7, 8, 9, 2, 3, 4}, []int{2, 3, 4, 6, 7, 8, 9}},
	}
	for _, test := range tests {
		MergeSort(test.input, intCmp)
		if !equal(test.input, test.want) {
			t.Errorf("mergesort(%v)", test.input)
		}
	}
}
