package quickSort

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

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input []interface{}
		want  []int
	}{
		{[]interface{}{128, 90}, []int{90, 128}},
		{[]interface{}{30, 12, 90}, []int{12, 30, 90}},
		{[]interface{}{89, 56, 256}, []int{56, 89, 256}},
		{[]interface{}{6, 7, 8, 1, 9, 2, 3, 4}, []int{1, 2, 3, 4, 6, 7, 8, 9}},
	}
	for _, test := range tests {
		origin := []interface{}{}
		origin = append(origin, test.input...)
		QuickSort(test.input, intCmp)
		if !equal(test.input, test.want) {
			t.Errorf("QuickSort(%v), got:%v, want:%v", origin, test.input, test.want)
		}
	}
}
