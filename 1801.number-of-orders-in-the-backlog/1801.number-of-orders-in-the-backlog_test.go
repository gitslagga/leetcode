package solution

import (
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	input := [][][]int{
		{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}},
		{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}},
	}
	expectResult := []int{6, 999999984}

	for k, v := range input {
		result := getNumberOfBacklogOrders(v)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
