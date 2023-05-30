package solution

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	best := math.MaxInt32
	ans := 0

	for i := range nums {
		// Remove the duplicates from the nums array.
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		// Now take the two pointer approach to solve the linear array problem.
		j, k := i+1, n-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			diff := abs(target - sum)
			if diff < best {
				ans = sum
				best = diff
			}

			if sum == target {
				return target
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
