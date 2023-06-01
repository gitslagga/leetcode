package solution

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	input := []int{3, 2, 1}
	expectResult := [][]string{
		{"((()))", "(()())", "(())()", "()(())", "()()()"},
		{"(())", "()()"},
		{"()"},
	}

	for k, v := range input {
		result := GenerateParenthesis(v)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
