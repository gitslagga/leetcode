package solution

func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		index := i
		j := 0
		for ; j < len(needle) && needle[j] == haystack[index]; j++ {
			index++
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}
