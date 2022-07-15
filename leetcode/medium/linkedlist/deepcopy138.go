package linkedlist

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }https://leetcode.cn/problems/copy-list-with-random-pointer/solution/
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cache map[*Node]*Node

func copyRandomList(head *Node) *Node {

	cache = map[*Node]*Node{}
	deepCopy(head)

	return head

}
func deepCopy(node *Node) *Node {
	if node == nil {
		return node
	}
	if v, has := cache[node]; has {
		return v
	}
	tempNode := &Node{Val: node.Val}
	cache[node] = tempNode
	tempNode.Next = deepCopy(node.Next)
	tempNode.Random = deepCopy(node.Random)
	return tempNode
}
func copyRandomList2(head *Node) *Node {

	if head == nil {
		return head
	}
	for node := head; node != nil; node = node.Next.Next {
		newNode := &Node{Val: node.Val, Next: node.Next}
		node.Next = newNode

	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	//现在把linked 分成2部分A→A′→B→B′ →C→C′
	target := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return target

}
