package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var dummy *ListNode = &ListNode{} // dummy节点，应对首节点需要删除的情况
	dummy.Next = head
	pre := dummy     // pre指针
	head = head.Next // pre.Next == head , 为了判断是否存在重复节点，head需=head.Next

	for head != nil {
		if head.Val == pre.Next.Val {
			head = head.Next
		} else {
			if pre.Next.Next == head { // 如果出现pre -> 2 -> 3(head) , 说明2可取
				pre = pre.Next
			} else { // 如果pre -> 2 -> 2 -> 3(head) -> 3，head = 3不一定可取，但是2需要删除，因此pre.Next = 3
				pre.Next = head
			}
			head = head.Next
		}
	}
	if pre.Next.Next != nil { // 如果1 -> 2(pre) -> 3 -> nil(head) , 则直接返回； 如果1 -> 2(pre) -> 3 -> 3 -> nil(head), 说明pre后面不可取
		pre.Next = nil
	}

	return dummy.Next
}
