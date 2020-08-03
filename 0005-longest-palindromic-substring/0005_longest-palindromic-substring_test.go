package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	var (
		s, expect, result string
	)

	asst := assert.New(t)

	s = ""
	expect = ""
	result = longestPalindrome(s)
	asst.Equal(expect, result, "wrong answer for %s.", s)

	s = "abbbbbc"
	expect = "bbbbb"
	result = longestPalindrome(s)
	asst.Equal(expect, result, "wrong answer for %s.", s)

	s = "a"
	expect = "a"
	result = longestPalindrome(s)
	asst.Equal(expect, result, "wrong answer for %s.", s)

	s = "ac"
	expect = "a"
	result = longestPalindrome(s)
	asst.Equal(expect, result, "wrong answer for %s.", s)

	s = "ccc"
	expect = "ccc"
	result = longestPalindrome(s)
	asst.Equal(expect, result, "wrong answer for %s.", s)

	s = "babad"
	expect = "bab"
	result = longestPalindrome(s)
	asst.Equal(expect, result, "wrong answer for %s.", s)

	s = "cbbd"
	expect = "bb"
	result = longestPalindrome(s)
	asst.Equal(expect, result, "wrong answer for %s.", s)

	s = "caba"
	expect = "aba"
	result = longestPalindrome(s)
	asst.Equal(expect, result, "wrong answer for %s.", s)

	s = "bananas"
	expect = "anana"
	result = longestPalindrome(s)
	asst.Equal(expect, result, "wrong answer for %s.", s)
}
