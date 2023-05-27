package solution

import (
	"testing"
)

func TestLongestPalindrome1(t *testing.T) {
	s := "babad"
	expectResult := "bab"

	result := LongestPalindrome(s)
	if result != expectResult {
		t.Errorf("return wrong result, expect:%v, got:%v", expectResult, result)
	}
}

func TestLongestPalindrome2(t *testing.T) {
	s := "cbbd"
	expectResult := "bb"

	result := LongestPalindrome(s)
	if result != expectResult {
		t.Errorf("return wrong result, expect:%v, got:%v", expectResult, result)
	}
}
