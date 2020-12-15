package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	var (
		num    int
		result bool
	)

	asst := assert.New(t)

	num = 121
	result = isPalindrome(num)
	asst.True(result, "wrong answer for %d.", num)

	num = 444
	result = isPalindrome(num)
	asst.True(result, "wrong answer for %d.", num)

	num = 123
	result = isPalindrome(num)
	asst.False(result, "wrong answer for %d.", num)

	num = 1221
	result = isPalindrome(num)
	asst.True(result, "wrong answer for %d.", num)

	num = 123454321
	result = isPalindrome(num)
	asst.True(result, "wrong answer for %d.", num)

	num = 121535252
	result = isPalindrome(num)
	asst.False(result, "wrong answer for %d.", num)
}
