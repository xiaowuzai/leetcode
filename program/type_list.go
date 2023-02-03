package main

type ListNode struct {
      Val int
      Next *ListNode
}

func NewHeadList(data []int) *ListNode{
	head := &ListNode{}
	p := head
	for k, v := range data {
		if k== 0 {
			p.Val=v
			continue
		}
		if p != nil {
			p.AppendListNode(v)
			p = p.Next
		}
	}
	return head
}

// 在当前节点后面追加节点。需要传入最后面的节点,eg : ln.Next == nil
func (ln *ListNode)AppendListNode(data int){
	ln.Next = &ListNode{Val: data}
}

// 在当前节点后面追加节点。需要传入最后面的节点,eg : ln.Next == nil
func (ln *ListNode)Print(){
	p := ln
	for p != nil {
		println(p.Val)
		p = p.Next
	}
}
