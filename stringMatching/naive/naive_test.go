package naive

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

func TestNaiveStringMatcher(t *testing.T) {

	tests := []struct {
		pattern string
		text    string
		want    []int
	}{
		{"123", "1234", []int{0}},
		{"123", "1231234", []int{0, 3}},
		{"123", "12", []int{}},
	}
	for _, test := range tests {
		got := NaiveStringMatcher(test.pattern, test.text)
		if !equal(got, test.want) {
			t.Errorf("NaiveStringMatcher(%v,%v) got %v, want:%v",
				test.pattern, test.text, got, test.want)
		}
	}

}
