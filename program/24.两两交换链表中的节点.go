package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	
	h := &ListNode{Next: head}	// 临时头结点
	head = h
	for h.Next != nil && h.Next.Next != nil {  //满足两个节点不为空的时候交换才有意义
		cur := h.Next
		pre := h.Next.Next
		tmp := pre.Next

		cur.Next = tmp
		pre.Next = cur
		h.Next = pre

		h = cur
	}
	
	return head.Next
}

// 循环迭代 -- 优化版本
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := &ListNode{Val:0, Next:head}
	head = pre

	first,sec := &ListNode{},&ListNode{}

	for pre.Next != nil && pre.Next.Next != nil {
		first,sec = pre.Next,pre.Next.Next

		pre.Next,first.Next,sec.Next,pre= sec,sec.Next,first,first

	}

	return head.Next
}
