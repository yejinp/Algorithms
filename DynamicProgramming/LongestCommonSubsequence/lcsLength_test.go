package lcs

import "testing"

func equal(a []byte, b []byte) bool {
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

func equal2(a [][]byte, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !equal(v, a[i]) {
			return false
		}
	}
	return true
}

func TestLcsLength(t *testing.T) {
	tests := []struct {
		X    string
		Y    string
		want int
	}{
		{"ABCD", "ACBD", 3},
		{"BDCABA", "ABCBDAB", 4},
		{"ACCGGTCGAGTGCGCGGAAGCCGGCCGAA", "GTCGTTCGGAATGCCGTTGCTCTGTAAA", 20},
	}

	for _, test := range tests {
		got := lcsLength(test.X, test.Y)
		if got != test.want {
			t.Errorf("lcsLength(%v,%v) got %v, want:%v",
				test.X, test.Y, got, test.want)
		}
	}
}

func TestLcsLength1(t *testing.T) {
	tests := []struct {
		X    string
		Y    string
		want int
	}{
		{"ABCD", "ACBD", 3},
		{"ABCBDAB", "BDCABA", 4},
		{"ACCGGTCGAGTGCGCGGAAGCCGGCCGAA", "GTCGTTCGGAATGCCGTTGCTCTGTAAA", 20},
	}

	for _, test := range tests {
		got := lcsLength1(test.X, test.Y)
		if got != test.want {
			t.Errorf("lcsLength(%v,%v) got %v, want:%v",
				test.X, test.Y, got, test.want)
		}
	}
}
