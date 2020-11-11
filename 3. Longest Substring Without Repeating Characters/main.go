func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	maxCount := 0
	curr := 0
	for i := 0; i < len(s); i++ {
		ch := s[i]

		if m[ch] != 0 && m[ch] >= curr {
			curr = m[ch]
		}

		if i+1-curr >= maxCount {
			maxCount = i + 1 - curr
		}

		m[ch] = i + 1
	}
	return maxCount
}
