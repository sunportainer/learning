package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := head.Next
	//1-->2-->3
	//1-->2-->3-->4-->5
	for even != nil || even.Next != nil {
		odd.Next = even.Next
		odd = even.Next
		even.Next = odd.Next
		even = odd.Next

	}
	odd.Next = evenHead
	return head

}
