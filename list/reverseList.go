package list

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		// 同时赋值
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}
