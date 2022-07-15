package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func loopNodes(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

//简单反转[1,2,3,4,5,6]
func Reverse(head *ListNode) *ListNode {
	if nil == head || head.Next == nil {
		return head
	}
	tail := head
	next := head.Next
	for next.Next != nil {
		tmp := next.Next
		next.Next = tail
		tail = next
		next = tmp
	}
	next.Next = tail
	head.Next = nil
	return next
}

//递归反转 图解https://segmentfault.com/a/1190000021364359
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := reverseListRecursive(head.Next)
	fmt.Printf("cuurent result node is %v \n", result)
	head.Next.Next = head
	head.Next = nil
	return result

}

func TestReverse() {
	head := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 4, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}
	node6 := &ListNode{Val: 6, Next: nil}
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	loopNodes(Reverse(head))

	head = &ListNode{Val: 1, Next: nil}
	node2 = &ListNode{Val: 2, Next: nil}
	node3 = &ListNode{Val: 3, Next: nil}
	head.Next = node2
	node2.Next = node3

	loopNodes(reverseListRecursive(head))
}

func TestReverseBetween() {
	head := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 4, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}
	node6 := &ListNode{Val: 6, Next: nil}
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	loopNodes(reverseBetween(head, 2, 4))

}

//Leetcode #92
/*给你单链表的头指针 head 和两个整left 和 right ，其left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

head = [5], left = 1, right = 1
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next

	}
	return dummyNode.Next

}

/*

	作者：LeetCode-Solution
	链接：https://leetcode.cn/problems/linked-list-cycle/solution/huan-xing-lian-biao-by-leetcode-solution/
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		fmt.Printf("slow is %v and fast is %v \n", slow, fast)
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true

}
func TestHasCircle() {
	head := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 4, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}
	node6 := &ListNode{Val: 6, Next: nil}
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node3
	fmt.Println(hasCycle(head))

}
