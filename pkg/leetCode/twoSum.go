package leetCode

func TwoSumBruteForce(nums []int, target int) []int {
	for i, n1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			if n1+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSumWithMap(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if j, found := m[target-num]; found {
			return []int{i, j}
		}
		m[num] = i
	}
	return nil
}
