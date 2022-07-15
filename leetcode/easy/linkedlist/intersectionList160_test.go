package linkedlist

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}

	heada := &ListNode{1, nil}
	node1 := &ListNode{2, nil}
	node2 := &ListNode{3, nil}
	node3 := &ListNode{4, nil}
	heada.Next = node1
	node1.Next = node2
	node2.Next = node3

	headb := &ListNode{1, nil}
	node1 = &ListNode{2, nil}
	node2 = &ListNode{3, nil}
	node3 = &ListNode{4, nil}
	node4 := &ListNode{5, nil}
	headb.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"leetcode160", args{heada, headb}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
