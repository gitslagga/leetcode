package solution

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	input := []int{121, -121, 10}
	expectResult := []bool{true, false, false}

	for k, v := range input {
		result := IsPalindrome(v)
		if result != expectResult[k] {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
