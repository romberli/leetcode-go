package leetcode

func isPalindrome(x int) bool {
	var numList []int

	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	for x > 0 {
		remainder := x % 10
		numList = append(numList, remainder)

		x = x / 10
	}

	length := len(numList)
	if length == 1 {
		return true
	}

	left := 0
	right := length - 1
	for i := 0; i < length; i++ {
		if numList[left] != numList[right] {
			return false
		}

		if right-left <= 2 {
			return true
		}

		left++
		right--
	}

	return false
}
