package main

import "fmt"

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序
的方式存储的，并且它们的每个节点只能存储 一位 数字。
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

// ListNode 链表
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
	num1 := fillData(352)
	trans(num1)
	num2 := fillData(653)
	trans(num2)
	target := addTwoNumbers(num1, num2)
	trans(target)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c1 := l1
	c2 := l2
	head := new(ListNode)
	d1 := head
	var sum int
	for c1 != nil || c2 != nil {
		sum /= 10
		if c1 != nil {
			sum += c1.Val
			c1 = c1.Next
		}
		if c2 != nil {
			sum += c2.Val
			c2 = c2.Next
		}
		d1.Next = new(ListNode)
		d1.Next.Val = sum % 10
		d1 = d1.Next
	}
	if sum/10 == 1 {
		d1.Next = new(ListNode)
		d1.Next.Val = 1
	}
	return head.Next
}
