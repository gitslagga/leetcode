package solution

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	input := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	}
	expectResult := []int{2, 5}

	for k, v := range input {
		result := RemoveDuplicates(v)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
