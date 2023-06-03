package solution

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	input := []struct {
		nums []int
		val  int
	}{
		{[]int{3, 2, 2, 3}, 3},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
	}
	expectResult := []int{2, 5}

	for k, v := range input {
		result := RemoveElement(v.nums, v.val)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
