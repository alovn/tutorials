package main

import "fmt"

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

	node.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	tmpNode := reverse(node)
	for tmpNode != nil {
		fmt.Println(tmpNode.Val)
		tmpNode = tmpNode.Next
	}
}

func reverse(head *ListNode) *ListNode {
	var newHead, tmp *ListNode
	for head != nil {
		tmp = head
		head = head.Next
		tmp.Next = newHead
		newHead = tmp
	}
	return newHead
}
