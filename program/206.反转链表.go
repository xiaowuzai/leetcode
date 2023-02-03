package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	var cur *ListNode
	pre := head
	for pre != nil {
		tmp := pre.Next
		pre.Next = cur

		cur = pre
		pre = tmp
	}

	return cur
}