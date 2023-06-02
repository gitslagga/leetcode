package solution

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	input := []struct {
		head *ListNode
		k    int
	}{
		{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{4, &ListNode{Val: 5, Next: nil}}}}}, 2},
		{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{4, &ListNode{Val: 5, Next: nil}}}}}, 3},
	}
	expectResult := []*ListNode{
		{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}}},
		{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
	}

	for k, v := range input {
		result := ReverseKGroup(v.head, v.k)
		if reflect.DeepEqual(result, expectResult[k]) == false {
			t.Errorf("return wrong result, expect:%v, got:%v", expectResult[k], result)
		}
	}
}
