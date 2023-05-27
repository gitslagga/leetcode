package solution

import (
	"testing"
)

func TestConvert(t *testing.T) {
	input := []string{"42", "   -42", "4193 with words11"}
	expectResult := []int{42, -42, 4193}

	for k, v := range input {
		result := MyAtoi(v)
		if result != expectResult[k] {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
