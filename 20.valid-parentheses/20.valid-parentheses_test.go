package solution

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	input := []string{"()", "()[]{}", "(]"}
	expectResult := []bool{true, true, false}

	for k, v := range input {
		result := IsValid(v)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
