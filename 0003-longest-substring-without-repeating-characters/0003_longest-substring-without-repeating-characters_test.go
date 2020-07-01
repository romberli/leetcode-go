package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var (
		s              string
		expect, result int
	)

	asst := assert.New(t)

	s = "ababbbbbbbcde"
	expect = 4
	result = lengthOfLongestSubstring(s)
	asst.Equal(expect, result, "wrong answer.")

	s = "abcbcdef"
	expect = 5
	result = lengthOfLongestSubstring(s)
	asst.Equal(expect, result, "wrong answer.")

	s = "abcade"
	expect = 5
	result = lengthOfLongestSubstring(s)
	asst.Equal(expect, result, "wrong answer.")
}
