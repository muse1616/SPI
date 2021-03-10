package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// special condition
	if head == nil || head.Next == nil{
		return head
	}
	// 三指针改变即可
	dummy:=&ListNode{Next:head}
	pre:=dummy
	for pre.Next!=nil && pre.Next.Next!=nil{
		// pre->1->2->3
		// pre->2->1->3
		pre.Next,pre.Next.Next,pre.Next.Next.Next = pre.Next.Next,pre.Next.Next.Next,pre.Next
		pre = pre.Next.Next
	}

	return dummy.Next
}