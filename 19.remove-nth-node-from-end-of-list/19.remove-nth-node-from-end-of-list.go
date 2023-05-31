package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	h := &ListNode{}
	h.Next = head

	right, left := h, h
	for i := 0; i < n; i++ {
		right = right.Next
	}
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	left.Next = left.Next.Next
	return h.Next
}
