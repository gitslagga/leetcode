package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	input := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
	}
	output := [][]int{{0, 1}, {1, 2}, {0, 1}}

	for k, v := range input {
		result := TwoSum(v.nums, v.target)
		if reflect.DeepEqual(result, output[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", output[k], result)
		}
	}
}
