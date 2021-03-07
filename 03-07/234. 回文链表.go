package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间O(n),空间O(n)
// func isPalindrome(head *ListNode) bool {
//     if head == nil{
//         return true
//     }
//     // 先计算链表长度
//     cur:=head
//     store:=make([]int,0)
//     for cur!=nil{
//         store = append(store,cur.Val)
//         cur = cur.Next
//     }
//     // 判断数组是否回文数组
//     for i:=0;i<len(store)/2;i++{
//         if store[i]!=store[len(store)-i-1]{
//             return false
//         }
//     }
//     return true
// }

// 时间O(n),空间O(1)
// 直接反转后半部分链表 比较前一部分和后一部分
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 要求空间O(1)所以在遍历链表的时候就必须判断出结果 不能使用辅助数组
	// 先用快慢指针找到后半部分链表的起点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		// 慢指针一次走一步
		slow = slow.Next
		// 快指针一次走两步
		fast = fast.Next.Next
	}
	// 循环结束时 慢指针指向后半部分链表开头 此时前后链表长度有可能不同(原链表长度为奇数时),之后再处理
	// 此时翻转后部分链表
	slow = reverse(slow)
	// 比较slow和head
	for head != nil && slow != nil {
		if head.Val != slow.Val {
			return false
		}
		head = head.Next
		slow = slow.Next
	}
	// slow有可能多出来的那个节点不需要管

	return true
}

// 翻转链表
func reverse(head *ListNode) *ListNode {
	var dummy *ListNode
	cur := head
	for cur != nil {
		// 保存下一个节点
		next := cur.Next
		cur.Next = dummy
		dummy = cur
		cur = next
	}
	return dummy
}
