package main

import "fmt"

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
*/

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func initData(nums ...int) *ListNode {
	head := new(ListNode)
	cur := head
	for i := 0; i < len(nums); i++ {
		cur.Next = new(ListNode)
		cur.Next.Val = nums[i]
		cur = cur.Next
	}
	return head.Next
}

func trans(header *ListNode) {
	for header != nil {
		fmt.Printf("%d -> ", header.Val)
		header = header.Next
	}
	fmt.Println()
}

func main() {
	num1 := initData(1, 2, 3, 4, 5)
	trans(num1)
	node := removeNthFromEnd(num1, 3)
	trans(node)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := new(ListNode)
	newHead.Next = head
	prev, curr := newHead, head
	i := 1
	for fast := head; fast.Next != nil; fast = fast.Next {
		if i < n {
			i++
		} else {
			prev, curr = prev.Next, curr.Next
		}
	}
	prev.Next = curr.Next
	return newHead.Next
}
