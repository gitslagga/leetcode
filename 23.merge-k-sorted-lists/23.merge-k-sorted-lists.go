package solution

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}
type IntHeap []*ListNode

func (h IntHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Len() int            { return len(h) }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	ans := (*h)[n-1]
	*h = (*h)[:n-1]
	return ans
}

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	h := IntHeap{}
	heap.Init(&h)
	for _, list := range lists {
		if list != nil {
			heap.Push(&h, list)
		}
	}
	dummy := &ListNode{}
	p := dummy
	for h.Len() > 0 {
		min := heap.Pop(&h).(*ListNode)
		p.Next = min
		p = p.Next
		if min.Next != nil {
			heap.Push(&h, min.Next)
		}
	}
	return dummy.Next
}
