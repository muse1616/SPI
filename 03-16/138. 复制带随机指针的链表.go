package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
// hash表实现映射
var hashTable = make(map[*Node]*Node, 0)

func copyRandomList(head *Node) *Node {
	// 递归终点
	if head == nil {
		return nil
	}
	// 存在映射
	if v, ok := hashTable[head]; ok {
		return v
	}
	// 新建节点
	node := &Node{Val: head.Val}
	// 加入映射
	hashTable[head] = node
	// 递归
	node.Next = copyRandomList(head.Next)
	node.Random = copyRandomList(head.Random)
	return node
}
