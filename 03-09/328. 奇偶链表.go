package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 拆分链表 然后合并即可
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 奇数头 偶数头
	// 可以用head表示奇头 此处重新定义一个变量便于查看
	oddHead, evenHead := head, head.Next
	// 双指针
	oddPtr, evenPtr := oddHead, evenHead
	for evenPtr != nil && evenPtr.Next != nil {
		// 1-2-3-4-5
		// 1-3-5
		// 1-3
		oddPtr.Next = evenPtr.Next
		// 此时oddPtr 3
		oddPtr = evenPtr.Next
		// 2-4
		evenPtr.Next = oddPtr.Next
		// 此时evenPtr 4
		evenPtr = oddPtr.Next
	}
	// 连接
	oddPtr.Next = evenHead
	return oddHead
}
