package solution

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	input := []struct {
		s string
		p string
	}{
		{"aa", "a"},
		{"aa", "a*"},
		{"ab", ".*"},
	}
	expectResult := []bool{false, true, true}

	for k, v := range input {
		result := IsMatch(v.s, v.p)
		if result != expectResult[k] {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
