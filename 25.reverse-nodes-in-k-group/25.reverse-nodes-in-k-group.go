package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverse(a, b)
	a.Next = ReverseKGroup(b, k)
	return newHead
}

func reverse(start, end *ListNode) *ListNode {
	var pre *ListNode
	cur, next := start, start
	for cur != end {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
