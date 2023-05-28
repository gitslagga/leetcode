package solution

import (
	"testing"
)

func TestIntegerToRoman(t *testing.T) {
	input := []int{3, 58, 1994}
	expectResult := []string{"III", "LVIII", "MCMXCIV"}

	for k, v := range input {
		result := IntegerToRoman(v)
		if result != expectResult[k] {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}

func TestIntToRoman(t *testing.T) {
	input := []int{3, 58, 1994}
	expectResult := []string{"III", "LVIII", "MCMXCIV"}

	for k, v := range input {
		result := IntToRoman(v)
		if result != expectResult[k] {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
