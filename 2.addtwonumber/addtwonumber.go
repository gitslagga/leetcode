package addtwonumber

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var m, v1, v2, v int
	cur1, cur2 := l1, l2
	dummyHead := &ListNode{}
	node := dummyHead
	for cur1 != nil || cur2 != nil {
		v1, v2 = 0, 0
		if cur1 != nil {
			v1 = cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			v2 = cur2.Val
			cur2 = cur2.Next
		}
		v = v1 + v2 + m
		if v >= 10 {
			v = v - 10
			m = 1
		} else {
			m = 0
		}

		node.Next = &ListNode{Val: v}
		node = node.Next
	}
	if m == 1 {
		node.Next = &ListNode{Val: 1}
	}
	return dummyHead.Next
}
