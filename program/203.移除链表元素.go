package main

// removeElements
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	// 设置一个虚拟头结点
	h := &ListNode{Next: head}
	p,f  :=h,h.Next
	for f != nil {
		if f.Val == val {
			p.Next = f.Next  // 越过 f
			f = f.Next  // 让 f 指向下个节点
		}else {
			f = f.Next
			p = p.Next
		}
	}

	return h.Next
}