package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	var (
		s      string
		p      string
		result bool
	)

	asst := assert.New(t)

	s = "b"
	p = "b.*c"
	result = isMatch(s, p)
	asst.False(result, "wrong answer for %s.", s)

	s = "a"
	p = "..*"
	result = isMatch(s, p)
	asst.True(result, "wrong answer for %s.", s)

	s = ""
	p = "c*c*"
	result = isMatch(s, p)
	asst.True(result, "wrong answer for %s.", s)

	s = "a"
	p = "..a*"
	result = isMatch(s, p)
	asst.False(result, "wrong answer for %s.", s)

	s = "aaa"
	p = "aaaa"
	result = isMatch(s, p)
	asst.False(result, "wrong answer for %s.", s)

	s = "a"
	p = "ab*"
	result = isMatch(s, p)
	asst.True(result, "wrong answer for %s.", s)

	s = "aa"
	p = "a*"
	result = isMatch(s, p)
	asst.True(result, "wrong answer for %s.", s)

	s = "aaa"
	p = "a*a"
	result = isMatch(s, p)
	asst.True(result, "wrong answer for %s.", s)

	s = "aaa"
	p = "a.a"
	result = isMatch(s, p)
	asst.True(result, "wrong answer for %s.", s)

	s = "ab"
	p = ".*c"
	result = isMatch(s, p)
	asst.False(result, "wrong answer for %s.", s)

	s = "abcd"
	p = "d*"
	result = isMatch(s, p)
	asst.False(result, "wrong answer for %s.", s)

	s = "abcdefg"
	p = "c*a*b"
	result = isMatch(s, p)
	asst.False(result, "wrong answer for %s.", s)

	s = "abcdefg"
	p = "a.*"
	result = isMatch(s, p)
	asst.True(result, "wrong answer for %s.", s)

	s = "sippi"
	p = "s*p*."
	result = isMatch(s, p)

	s = "mississippi"
	p = "mis*is*p*."
	result = isMatch(s, p)
	asst.False(result, "wrong answer for %s.", s)
}
