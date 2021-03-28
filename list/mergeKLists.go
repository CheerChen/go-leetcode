package list

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := new(IntHeap)
	heap.Init(pq)
	for _, list := range lists {
		for list != nil {
			heap.Push(pq, list.Val)
			list = list.Next
		}
	}
	p := new(ListNode)
	head := p
	for pq.Len() > 0 {
		p.Next = &ListNode{Val: heap.Pop(pq).(int)}
		p = p.Next
	}
	return head.Next
}
