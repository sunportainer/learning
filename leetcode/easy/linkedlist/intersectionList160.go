package linkedlist

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	//这个pa != pb判断很巧妙，即使a和b完全不交叉，因为2个指针最多都会走lenth_a + length_b的长度，到最后肯定都同时走各个链表的
	//的最后，pa在b链表的最后， pb在a链表的最后。
	//如果2个link的长度一样，就同时走到最后，发现都是nil，都不交叉，就返回了
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa

}
func getIntersectionNodeWithMap(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	hsh := make(map[*ListNode]bool)
	tmp := headA
	for tmp != nil {
		hsh[tmp] = true
		tmp = tmp.Next
	}
	tmp = headB
	for tmp != nil {
		if _, ok := hsh[tmp]; ok {
			return tmp
		}
		tmp = tmp.Next
	}

	return nil

}
