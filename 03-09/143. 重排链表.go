package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 一分为二
// 后半部分逆序
// 前后组合
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// 找到中点
	middle := findMiddle(head)
	l1, l2 := head, middle.Next
	// 前后断开
	middle.Next = nil
	// 反转
	l2 = reverse(l2)
	// 合并
	mergeList(l1, l2)
}

// 寻找中点
func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 组合链表
func mergeList(l1, l2 *ListNode) {
	var cur1, cur2 *ListNode
	for l1 != nil && l2 != nil {
		cur1 = l1.Next
		cur2 = l2.Next
		l1.Next = l2
		l1 = cur1
		l2.Next = l1
		l2 = cur2
	}
}

// 反转链表
func reverse(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}
