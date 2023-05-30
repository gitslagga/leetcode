package solution

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	input := []string{"23", "", "2"}
	expectResult := [][]string{
		{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		{},
		{"a", "b", "c"},
	}

	for k, v := range input {
		result := LetterCombinations(v)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
