package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	var (
		nums1, nums2   []int
		expect, result float64
	)

	asst := assert.New(t)

	// nums1 = []int{1, 3, 5, 7, 9}
	// nums2 = []int{2, 4}
	// expect = float64(4)
	// result = findMedianSortedArrays(nums1, nums2)
	// asst.Equal(expect, result, "wrong answer.")
	//
	// nums1 = []int{1, 2}
	// nums2 = []int{3, 4}
	// expect = float64(2.5)
	// result = findMedianSortedArrays(nums1, nums2)
	// asst.Equal(expect, result, "wrong answer.")

	// nums1 = []int{1, 2}
	// nums2 = []int{-1, 3}
	// expect = float64(1.5)
	// result = findMedianSortedArrays(nums1, nums2)
	// asst.Equal(expect, result, "wrong answer.")

	nums1 = []int{4}
	nums2 = []int{1, 2, 3, 5}
	expect = float64(3)
	result = findMedianSortedArrays(nums1, nums2)
	asst.Equal(expect, result, "wrong answer.")
}
