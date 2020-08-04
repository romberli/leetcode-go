package leetcode

func findLongestPalindromicWithIndex(s string, length int, index int) (string, int) {
	left := index
	right := index

	leftEdge := 0
	rightEdge := length - 1

	for left > leftEdge && s[left-1] == s[index] {
		left--
	}

	for right < rightEdge && s[right+1] == s[index] {
		right++
	}

	if left == leftEdge {
		return s[:right+1], right
	}

	for {
		if left < leftEdge || right > rightEdge || s[left] != s[right] {
			return s[left+1 : right], right - 1
		}

		left--
		right++
	}

}

func longestPalindrome(s string) string {
	result := ""

	length := len(s)

	if length <= 1 {
		return s
	}

	for i, _ := range s {
		str, right := findLongestPalindromicWithIndex(s, length, i)

		if len(str) > len(result) {
			result = str
		}

		if right == length-1 {
			return result
		}
	}

	return result
}
