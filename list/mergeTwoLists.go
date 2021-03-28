package list

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	merge := &ListNode{}
	mergeh := merge
	for l1 != nil {
		if l2 == nil {
			break
		}
		if l1.Val >= l2.Val {
			merge.Next = l2
			merge = merge.Next
			l2 = l2.Next
		} else {
			merge.Next = l1
			merge = merge.Next
			l1 = l1.Next
		}
	}

	for l2 != nil {
		if l1 == nil {
			break
		}
		if l1.Val >= l2.Val {
			merge.Next = l2
			merge = merge.Next
			l2 = l2.Next
		} else {
			merge.Next = l1
			merge = merge.Next
			l1 = l1.Next
		}
	}

	for l1 != nil {
		merge.Next = l1
		merge = merge.Next
		l1 = l1.Next
	}

	for l2 != nil {
		merge.Next = l2
		merge = merge.Next
		l2 = l2.Next
	}

	return mergeh.Next

}
