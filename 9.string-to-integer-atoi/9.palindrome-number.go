package solution

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	intString := strconv.Itoa(x)
	n := len(intString)
	for i := 0; i < len(intString)/2; i++ {
		if rune(intString[i]) != rune(intString[n-1-i]) {
			return false
		}
	}
	return true
}
