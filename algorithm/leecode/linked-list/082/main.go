package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 3}
	node5 := &ListNode{Val: 4}
	node6 := &ListNode{Val: 4}
	node7 := &ListNode{Val: 5}
	node8 := &ListNode{Val: 5}

	node.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node8

	var tmpNode *ListNode = deleteDuplicates(node)
	for tmpNode != nil {
		fmt.Println(tmpNode.Val)
		tmpNode = tmpNode.Next
	}

}

//x，1,2,3,3,4,4,5,5
func deleteDuplicates(head *ListNode) *ListNode {

	newHead := &ListNode{Next: head}
	pre := newHead
	cur := head

	for cur != nil {
		duplicate := false
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
			duplicate = true
		}
		cur = cur.Next
		if duplicate {
			pre.Next = cur

		} else {
			pre = pre.Next
		}
	}

	return newHead.Next
}
