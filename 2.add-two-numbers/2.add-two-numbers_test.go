package solution

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbersExample1(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	expectResult := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}
	result := AddTwoNumbers(l1, l2)
	if reflect.DeepEqual(result, expectResult) == false {
		t.Errorf("return wrong result, expect:%v, got:%v", expectResult, result)
	}
}

func TestAddTwoNumbersExample2(t *testing.T) {
	l1 := &ListNode{
		Val:  0,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  0,
		Next: nil,
	}
	expectResult := &ListNode{
		Val:  0,
		Next: nil,
	}
	result := AddTwoNumbers(l1, l2)
	if reflect.DeepEqual(result, expectResult) == false {
		t.Errorf("return wrong result, expect:%v, got:%v", expectResult, result)
	}
}

func TestAddTwoNumbersExample3(t *testing.T) {
	//Input: tmp1 = [9,9,9,9,9,9,9], tmp2 = [9,9,9,9]
	//Output: [8,9,9,9,0,0,0,1]
	tmp1, tmp2 := &ListNode{}, &ListNode{}
	l1, l2 := tmp1, tmp2
	for i := 0; i < 7; i++ {
		tmp1.Val = 9
		tmp1.Next = &ListNode{}
		tmp1 = tmp1.Next
	}
	for i := 0; i < 4; i++ {
		tmp2.Val = 9
		tmp2.Next = &ListNode{}
		tmp2 = tmp2.Next
	}
	expectResult := []int{8, 9, 9, 9, 0, 0, 0, 1}
	result := AddTwoNumbers(l1, l2)

	for i := 0; result != nil; i++ {
		if i >= len(expectResult) {
			t.Errorf("runtime error: index out of range")
		}
		if expectResult[i] != result.Val {
			t.Errorf("return wrong result, expect:%v, got:%+v", expectResult[i], result)
		}
		result = result.Next
	}
}
