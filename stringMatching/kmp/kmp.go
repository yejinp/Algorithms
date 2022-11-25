package kmp

func KmpStringMatcher(pattern, text string) []int {
	ans := []int{}
	n, m := len(text), len(pattern)
	next := computePrefix(pattern)

	q := 0 // number of charaters matched

	for i := 0; i < n; i++ { // scan the text from left to right
		for 0 < q && pattern[q] != text[i] {
			q = next[q] + 1 // next character does not match
		}
		if pattern[q] == text[i] {
			q++ // next character matches
		}
		if q == m { // is all of P matched ?
			ans = append(ans, i-m+1)
			q = next[q-1] + 1 // look for the next match
		}
	}
	return ans
}

func computePrefix(pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}
	next := make([]int, len(pattern))
	next[0] = -1
	k, m := -1, len(pattern)

	for q := 1; q < m; q++ {
		for 0 <= k && pattern[k+1] != pattern[q] {
			k = next[k]
		}
		if pattern[k+1] == pattern[q] {
			k++
		}
		next[q] = k
	}
	return next
}
