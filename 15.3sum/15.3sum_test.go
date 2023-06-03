package solution

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	input := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{0, 1, 1},
		{0, 0, 0},
		{-2, 0, 0, 2, 2},
	}
	expectResult := [][][]int{
		{{-1, -1, 2}, {-1, 0, 1}},
		nil,
		{{0, 0, 0}},
		{{-2, 0, 2}},
	}

	for k, v := range input {
		result := ThreeSum(v)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
