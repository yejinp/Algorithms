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
		{"B", "B", 1},
		{"B", "BA", 1},
		{"BA", "BA", 2},
		{"BAB", "BA", 2},
		{"BA", "BAB", 2},
		{"ABCD", "AD", 2},
		{"ABCD", "ACBD", 3},
		{"ABCD", "ACBD", 3},
		{"ABCD", "ABCD", 4},
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
