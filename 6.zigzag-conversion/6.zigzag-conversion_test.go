package solution

import (
	"testing"
)

func TestConvert(t *testing.T) {
	input := []struct {
		s       string
		numRows int
	}{
		{"PAYPALISHIRING", 3},
		{"PAYPALISHIRING", 4},
		{"A", 1},
	}
	expectResult := []string{"PAHNAPLSIIGYIR", "PINALSIGYAHRPI", "A"}

	for k, v := range input {
		result := Convert(v.s, v.numRows)
		if result != expectResult[k] {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
