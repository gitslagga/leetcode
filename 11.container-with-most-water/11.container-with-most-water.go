package solution

func MaxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		ans = max(ans, area)
		if height[l] <= height[r] {
			l += 1
		} else {
			r -= 1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
