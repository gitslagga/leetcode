package solution

func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for k, v := range nums {
		if i, ok := hashMap[target-v]; ok {
			return []int{i, k}
		}
		hashMap[v] = k
	}
	return nil
}
