package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	var (
		s, expect, result string
		rows              int
	)

	asst := assert.New(t)

	s = "AB"
	rows = 1
	expect = "AB"
	result = convert(s, rows)
	asst.Equal(expect, result, "wrong answer for %s.", s)

	s = "ABC"
	rows = 2
	expect = "ACB"
	result = convert(s, rows)
	asst.Equal(expect, result, "wrong answer for %s.", s)

	s = "ABCDE"
	rows = 4
	expect = "ABCED"
	result = convert(s, rows)
	asst.Equal(expect, result, "wrong answer for %s.", s)

	s = "LEETCODEISHIRING"
	rows = 3
	expect = "LCIRETOESIIGEDHN"
	result = convert(s, rows)
	asst.Equal(expect, result, "wrong answer for %s.", s)

	s = "LEETCODEISHIRING"
	rows = 4
	expect = "LDREOEIIECIHNTSG"
	result = convert(s, rows)
	asst.Equal(expect, result, "wrong answer for %s.", s)

	s = "PAYPALISHIRING"
	rows = 3
	expect = "PAHNAPLSIIGYIR"
	result = convert(s, rows)
	asst.Equal(expect, result, "wrong answer for %s.", s)
}
