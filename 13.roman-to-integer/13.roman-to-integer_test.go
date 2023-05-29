package solution

import (
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	input := []string{"III", "LVIII", "MCMXCIV"}
	expectResult := []int{3, 58, 1994}

	for k, v := range input {
		result := RomanToInteger(v)
		if result != expectResult[k] {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
