package list

func reverseKGroup(head *ListNode, k int) *ListNode {
	a := head
	b := head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	head = reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return head
}

func reverse(a *ListNode, b *ListNode) *ListNode {
	var prev *ListNode
	cur := a
	for cur != b {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}
