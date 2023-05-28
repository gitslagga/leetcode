package solution

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	input := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
	}
	expectResult := []int{49, 1}

	for k, v := range input {
		result := MaxArea(v)
		if result != expectResult[k] {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
