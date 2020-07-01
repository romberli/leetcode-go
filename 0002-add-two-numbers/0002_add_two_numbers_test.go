package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func listToListNode(l []int) *ListNode {
	result := &ListNode{l[0], nil}
	tmp := result

	for i := 1; i < len(l); i++ {
		tmp.Next = &ListNode{l[i], nil}
		tmp = tmp.Next
	}

	return result
}

func TestAddTwoNumbers(t *testing.T) {
	asst := assert.New(t)

	var l1, l2, expect, result *ListNode

	// n1 := []int{5, 4, 6}
	// n2 := []int{5, 5, 3}
	// sum := []int{0, 0, 0, 1}
	n1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	n2 := []int{5, 6, 4}
	sum := []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

	expect = listToListNode(sum)

	l1 = listToListNode(n1)
	l2 = listToListNode(n2)
	result = addTwoNumbers(l1, l2)

	t.Logf("num1: %v\nnum2: %v\nexpect: %v\nresult: %v.", n1, n2, expect, result)

	for {
		if result == nil && expect == nil {
			return
		}

		asst.Equal(result.Val, expect.Val, "wrong answer.")
		result = result.Next
		expect = expect.Next
	}
}
