package solution

import (
	"testing"
)

func TestConvert(t *testing.T) {
	input := []int{123, -123, 120}
	expectResult := []int{321, -321, 21}

	for k, v := range input {
		result := Reverse(v)
		if result != expectResult[k] {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
