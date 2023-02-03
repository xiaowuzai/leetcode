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
	
	cur := head
	pre := head.Next
	for pre != nil {
		tmp := pre.Next

		pre.Next = cur
		cur.Next = tmp

		cur = tmp

	}
}
