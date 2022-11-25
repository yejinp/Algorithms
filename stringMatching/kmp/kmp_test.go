package kmp

import "testing"

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

func TestKmpStringMatcher(t *testing.T) {

	tests := []struct {
		pattern string
		text    string
		want    []int
	}{
		{"12123", "12", []int{}},
		{"12123", "12123", []int{0}},
		{"123", "1231234", []int{0, 3}},
	}
	for _, test := range tests {
		got := KmpStringMatcher(test.pattern, test.text)
		//got := Kmp(test.text, test.pattern)

		if !equal(got, test.want) {
			t.Errorf("KmpStringMatcher(%v,%v) got %v, want:%v",
				test.pattern, test.text, got, test.want)
		}
	}

}

func TestComputePrefix(t *testing.T) {

	tests := []struct {
		pattern string
		want    []int
	}{
		{"a", []int{-1}},
		{"11", []int{-1, 0}},
		{"aba", []int{-1, -1, 0}},
		{"111", []int{-1, 0, 1}},
		{"124", []int{-1, -1, -1}},
		{"1212", []int{-1, -1, 0, 1}},
		{"12123", []int{-1, -1, 0, 1, -1}},
		{"1231231231234", []int{-1, -1, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, -1}},
	}
	for _, test := range tests {
		got := computePrefix(test.pattern)

		if !equal(got, test.want) {
			t.Errorf("computePrefix(%v) got %v, want:%v",
				test.pattern, got, test.want)
		}
	}

}
