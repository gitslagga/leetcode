package solution

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	input := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
	}
	expectResult := []string{"fl", ""}

	for k, v := range input {
		result := LongestCommonPrefix(v)
		if result != expectResult[k] {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
