package solution

import (
	"reflect"
	"testing"
)

func TestDivide(t *testing.T) {
	input := []struct {
		dividend int
		divisor  int
	}{
		{10, 3},
		{7, -3},
	}
	expectResult := []int{3, -2}

	for k, v := range input {
		result := Divide(v.dividend, v.divisor)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
