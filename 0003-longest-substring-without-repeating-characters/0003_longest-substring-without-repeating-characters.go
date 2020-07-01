package leetcode

func moveLeftSide(s string, m map[string]int) (string, map[string]int) {
	d := string(s[0])
	s = s[1:]
	m[d]--

	if m[d] == 0 {
		delete(m, d)
	}

	return s, m
}

func lengthOfLongestSubstring(s string) int {
	result := ""
	m := make(map[string]int)

	for _, r := range s {
		c := string(r)
		result += c

		if _, ok := m[c]; ok {
			result, m = moveLeftSide(result, m)
		} else {
			for _, v := range m {
				if v > 1 {
					result, m = moveLeftSide(result, m)
					break
				}
			}
		}

		m[c]++
	}

	return len(result)
}
