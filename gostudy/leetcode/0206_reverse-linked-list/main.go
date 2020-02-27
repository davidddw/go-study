package main

import "fmt"

/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
示例 1:
输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
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
	node := reverseList(num1)
	trans(node)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		prev, curr, curr.Next = curr, curr.Next, prev
	}
	return prev
}
