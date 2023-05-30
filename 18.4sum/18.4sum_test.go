package solution

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	input := []struct {
		nums   []int
		target int
	}{
		{[]int{1, 0, -1, 0, -2, 2}, 0},
		{[]int{2, 2, 2, 2, 2}, 8},
	}
	expectResult := [][][]int{
		{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		{{2, 2, 2, 2}},
	}

	for k, v := range input {
		result := FourSum(v.nums, v.target)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
