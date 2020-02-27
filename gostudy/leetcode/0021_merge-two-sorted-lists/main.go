package main

import (
	"fmt"
)

/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func fillData(nums int) *ListNode {
	head := new(ListNode)
	cur := head
	for nums > 0 {
		m := nums % 10
		cur.Next = new(ListNode)
		cur.Next.Val = m
		cur = cur.Next
		nums /= 10
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
	num1 := fillData(421)
	trans(num1)
	num2 := fillData(431)
	trans(num2)
	target := mergeTwoLists(num1, num2)
	trans(target)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	res := new(ListNode)
	if l1.Val >= l2.Val {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}
	return res
}
