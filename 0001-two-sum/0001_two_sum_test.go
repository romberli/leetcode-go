package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	asst := assert.New(t)

	nums := []int{2, 8, 17, 5, 1}
	target := 7
	expect := []int{0, 3}

	result := twoSum(nums, target)

	t.Logf("nums: %v\ntarget: %d\nexpect: %v\nresult: %v.", nums, target, expect, result)

	asst.Equal(len(expect), len(result), "wrong answer.")

	for i := 0; i < len(expect); i++ {
		asst.Equal(expect[i], result[i], "wrong answer.")
	}
}
