package solution

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	input := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{4, nil}}}},
		{},
		{Val: 1, Next: nil},
	}
	expectResult := []*ListNode{
		{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}},
		{},
		{Val: 1, Next: nil},
	}

	for k, v := range input {
		result := SwapPairs(v)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
