package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// 哈希表
	// hashTable:=make(map[int]bool,0)
	// dummy:=&ListNode{Next:head}
	// cur:=dummy
	// for cur.Next!=nil{
	//     if _,ok:=hashTable[cur.Next.Val];ok{
	//         cur.Next = cur.Next.Next
	//     }else{
	//         hashTable[cur.Next.Val] = true
	//         cur = cur.Next
	//     }
	// }
	// return dummy.Next
	// 利用有序特点 双指针 prev cur
	if head == nil || head.Next == nil {
		return head
	}
	prev, cur := head, head.Next

	for cur != nil {
		if cur.Val == prev.Val {
			prev.Next = cur.Next
			cur = cur.Next
		} else {
			prev = cur
			cur = cur.Next
		}
	}
	return head

}
