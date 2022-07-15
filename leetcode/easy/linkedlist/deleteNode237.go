package linkedlist

func deleteNode(node *ListNode) {
	next := node.Next

	for next != nil {
		t := next.Val
		node.Val = t

		next = next.Next
		node = node.Next
	}
	node.Next = nil

}
