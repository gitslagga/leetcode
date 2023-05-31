package solution

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	target := 0
	sort.Ints(nums)
	n := len(nums)
	var list [][]int

	for i := range nums {
		// Remove the duplicates from the nums array.
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		// Now take the two pointer approach to solve the linear array problem.
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				list = append(list, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return list
}
