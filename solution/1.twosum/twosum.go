package solution

func TwoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for k, v := range nums {
		if i, ok := numsMap[target-v]; ok {
			return []int{i, k}
		}
		numsMap[v] = k
	}
	return nil
}
