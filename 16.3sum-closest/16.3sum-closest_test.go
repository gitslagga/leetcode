package solution

import (
	"reflect"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	input := []struct {
		nums   []int
		target int
	}{
		{[]int{-1, 2, 1, -4}, 1},
		{[]int{0, 0, 0}, 1},
		{[]int{1, 1, 1, 0}, 100},
	}
	expectResult := []int{2, 0, 3}

	for k, v := range input {
		result := ThreeSumClosest(v.nums, v.target)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
