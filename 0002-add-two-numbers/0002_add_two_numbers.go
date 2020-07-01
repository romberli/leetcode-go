package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func processCarry(l *ListNode) {
	if l == nil || l.Val < 10 {
		return
	}

	l.Val %= 10
	l.Next = &ListNode{1, nil}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var valL1, valL2 int
	var tmp, result *ListNode

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	tmp = &ListNode{l1.Val + l2.Val, nil}
	processCarry(tmp)
	result = tmp

	for {
		if l1.Next == nil && l2.Next == nil {
			return result
		}

		if l1.Next == nil {
			valL1 = 0
		} else {
			l1 = l1.Next
			valL1 = l1.Val
		}

		if l2.Next == nil {
			valL2 = 0
		} else {
			l2 = l2.Next
			valL2 = l2.Val
		}

		if tmp.Next == nil {
			tmp.Next = &ListNode{valL1 + valL2, nil}
			tmp = tmp.Next
			processCarry(tmp)

			continue
		}

		tmp.Next.Val += valL1 + valL2
		tmp = tmp.Next
		processCarry(tmp)
	}
}
