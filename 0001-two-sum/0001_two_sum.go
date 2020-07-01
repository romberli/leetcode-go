package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, x := range nums {
		difference := target - x

		if j, ok := m[difference]; ok {
			return []int{j, i}
		}

		m[x] = i
	}

	return nil
}
