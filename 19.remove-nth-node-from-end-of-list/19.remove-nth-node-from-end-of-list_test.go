package solution

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	input := []struct {
		head *ListNode
		n    int
	}{
		{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}, 2},
		{&ListNode{Val: 1, Next: nil}, 1},
		{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}, 1},
	}
	expectResult := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}},
		nil,
		{1, nil},
	}

	for k, v := range input {
		result := RemoveNthFromEnd(v.head, v.n)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
