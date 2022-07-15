package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	array := []int{}
	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}
	i, j := 0, len(array)-1
	for i < j {
		if array[i] != array[j] {
			return false
		}
		i++
		j--
	}
	return true

}
func reverse(head *ListNode) *ListNode {
	tmp := head
	if head.Next != nil {
		head = reverse(head.Next)
		tmp.Next.Next = tmp
		tmp.Next = nil
	}

	return head
}
