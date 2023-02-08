package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	cur := &ListNode{Next: head}
	head = cur

	arr := make([]*ListNode, 0, n)
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}

	if len(arr) <= n {
		return head
	}

	if len(arr)-n+1 >= len(arr){
		arr[len(arr)-n-1].Next = nil
	}else{
		arr[len(arr)-n-1].Next = arr[len(arr)-n+1]
	}

	return head.Next
}

// 双指针
// 一前一后指针，n是几，前面的指针就先走几步
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	h := &ListNode{Next: head}
	head = h

	fast,slow := head,head
	// 走 n+1 步。 这样 slow 会指向被删除节点的前一个节点
	for n >= 0 {
		fast = fast.Next
		n--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head.Next
}
