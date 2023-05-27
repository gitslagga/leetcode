package solution

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	input := []string{"abcabcbb", "bbbbb", "pwwkew"}
	expectResult := []int{3, 1, 3}

	for k, v := range input {
		result := LengthOfLongestSubstring(v)
		if result != expectResult[k] {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
