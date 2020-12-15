package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	var (
		str    string
		result int
		expect int
	)

	asst := assert.New(t)

	str = "words and 987"
	expect = 0
	result = myAtoi(str)
	asst.Equal(expect, result, "wrong answer for %s.", str)

	str = "42"
	expect = 42
	result = myAtoi(str)
	asst.Equal(expect, result, "wrong answer for %s.", str)

	str = "-12300abc"
	expect = -12300
	result = myAtoi(str)
	asst.Equal(expect, result, "wrong answer for %s.", str)

	str = "-abc"
	expect = 0
	result = myAtoi(str)
	asst.Equal(expect, result, "wrong answer for %s.", str)

	str = "-9876543210"
	expect = MinInt
	result = myAtoi(str)
	asst.Equal(expect, result, "wrong answer for %s.", str)

	str = "9876543210"
	expect = MaxInt
	result = myAtoi(str)
	asst.Equal(expect, result, "wrong answer for %s.", str)
}
