package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//
func mergeKLists(lists []*ListNode) (ans *ListNode) {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	for _, v := range lists {
		ans = mergeTwoLists(ans, v)
	}
	return
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				cur.Next = l1
				l1 = l1.Next
				cur = cur.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
				cur = cur.Next
			}
		} else if l1 != nil {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next
		} else if l2 != nil {
			cur.Next = l2
			l2 = l2.Next
			cur = cur.Next
		}
	}

	return dummy.Next
}
