package leetCode

/*
Given an array of integers nums and an integer target, return indices
of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution,
and you may not use the same element twice.

You can return the answer in any order.
*/

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
