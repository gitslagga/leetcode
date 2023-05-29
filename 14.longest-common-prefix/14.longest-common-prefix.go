package solution

import (
	"sort"
)

func LongestCommonPrefix(v []string) string {
	ans := ""
	sort.Strings(v)
	first := v[0]
	last := v[len(v)-1]

	for i := 0; i < min(len(first), len(last)); i++ {
		if first[i] == last[i] {
			ans += string(first[i])
		} else {
			break
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
