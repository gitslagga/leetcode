package solution

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	input := []struct {
		l1, l2 *ListNode
	}{
		{
			&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}},
			&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
		},
		{nil, nil},
		{nil, &ListNode{Val: 0, Next: nil}},
	}
	expectResult := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}}},
		nil,
		{Val: 0, Next: nil},
	}

	for k, v := range input {
		result := MergeTwoLists(v.l1, v.l2)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
