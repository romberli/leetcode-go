package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	var (
		x, expect, result int
	)

	asst := assert.New(t)

	x = -123
	expect = -321
	result = reverse(x)
	asst.Equal(expect, result, "wrong answer for %d.", x)

	x = -2147483412
	expect = -2143847412
	result = reverse(x)
	asst.Equal(expect, result, "wrong answer for %d.", x)
}
