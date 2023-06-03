package solution

import (
	"reflect"
	"testing"
)

func TestStrStr(t *testing.T) {
	input := []struct {
		haystack string
		needle   string
	}{
		{"sadbutsad", "sad"},
		{"sadbutsad", "adb"},
		{"leetcode", "leeto"},
	}
	expectResult := []int{0, 1, -1}

	for k, v := range input {
		result := StrStr(v.haystack, v.needle)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
