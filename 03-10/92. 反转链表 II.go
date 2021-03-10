package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 如果left==1 则相当于翻转前right个节点
	if left == 1 {
		return reverseN(head, right)
	}
	// left不等于1时直接往后快进到left==1
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

// 记录后驱节点
var successor *ListNode

//翻转前n个节点链表
func reverseN(head *ListNode, N int) *ListNode {
	if N == 1 {
		successor = head.Next
		return head
	}
	last := reverseN(head.Next, N-1)
	head.Next.Next = head
	head.Next = successor
	return last
}
