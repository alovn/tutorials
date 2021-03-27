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
	node4 := &ListNode{Val: 4}
	node.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 //环形

	// node4.Next = &ListNode{Val: 5}
	fmt.Println(hasCycle(node))
}
func hasCycle(head *ListNode) bool {
	cur := head
	next := head
	for cur != nil && next != nil && next.Next != nil {
		cur = cur.Next
		next = next.Next.Next
		if cur == next {
			return true
		}
	}
	return false
}
