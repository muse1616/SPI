package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 只有一个节点 反转后就是自身
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	// 1-><-2<-3
	head.Next.Next = head
	head.Next = nil
	return last
}
