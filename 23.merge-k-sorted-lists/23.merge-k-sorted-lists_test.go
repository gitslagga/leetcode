package solution

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	input := [][]*ListNode{
		{
			&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}},
			&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
			&ListNode{Val: 2, Next: &ListNode{Val: 6, Next: nil}},
		},
		{},
		{nil},
	}
	expectResult := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}}},
		nil,
		nil,
	}

	for k, v := range input {
		result := MergeKLists(v)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
