package leetcode

func findInsertIndex(nums []int, n, start, end int) int {
	if n <= nums[start] {
		return start
	}

	if n >= nums[end] {
		return end + 1
	}

	if end-start == 1 {
		if n < nums[start] {
			return start
		}

		if n > nums[end] {
			return end + 1
		}

		return end
	}

	median := (end-start)/2 + start

	if n >= nums[median] {
		start = median
	} else {
		end = median
	}

	return findInsertIndex(nums, n, start, end)
}

func AppendSliceBeforeIndex(nums []int, n int, index int) []int {
	var result []int

	if index < 0 {
		result = append(result, n)
		result = append(result, nums[:]...)
		return result
	}

	if index > len(nums) {
		return append(nums, n)
	}

	result = append(result, nums[:index]...)
	result = append(result, n)
	result = append(result, nums[index:]...)

	return result
}

func findMedianSortedArray(nums []int) float64 {
	if nums == nil || len(nums) == 0 {
		return 0.0
	}

	length := len(nums)
	if length%2 == 0 {
		return float64(nums[length/2-1]+nums[length/2]) / 2.0
	}

	return float64(nums[length/2])
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if nums1 == nil || len(nums1) == 0 {
		return findMedianSortedArray(nums2)
	}

	if nums2 == nil || len(nums2) == 0 {
		return findMedianSortedArray(nums1)
	}

	len1 := len(nums1)
	len2 := len(nums2)
	end1 := len1 - 1
	end2 := len2 - 1

	if nums1[0] >= nums2[end2] {
		nums2 = append(nums2, nums1...)
		return findMedianSortedArray(nums2)
	}

	if nums2[0] >= nums1[end1] {
		nums1 = append(nums1, nums2...)
		return findMedianSortedArray(nums1)
	}

	start := 0

	for i, n := range nums2 {
		index := findInsertIndex(nums1, n, start, end1)

		if index > end1+1 {
			nums1 = append(nums1, nums2[i:]...)
			break
		}

		nums1 = AppendSliceBeforeIndex(nums1, n, index)

		start = index

		if index == start && start > 0 {
			start--
		}
		end1++
	}

	return findMedianSortedArray(nums1)
}
