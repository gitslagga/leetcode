package solution

import (
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	input := []struct {
		s     string
		words []string
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}},
		{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}},
	}
	expectResult := [][]int{
		{0, 9},
		{},
		{6, 9, 12},
	}

	for k, v := range input {
		result := FindSubstring(v.s, v.words)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
