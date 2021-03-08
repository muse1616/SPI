package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 求长度
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	// 按照题目 k 应该不会大于 length
	cur = head
	for i := 0; i < length-k; i++ {
		cur = cur.Next
	}
	return cur
}
