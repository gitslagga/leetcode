package solution

func LengthOfLongestSubstring(s string) int {
	var (
		ans     int
		left    int
		chatSet [128]int
	)
	for right, ch := range s {
		if chatSet[ch] > left {
			left = chatSet[ch]
		}
		chatSet[ch] = right + 1
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
