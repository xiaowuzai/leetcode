package main

type MyLinkedList struct {
	num int
	Next *ListNode
}




func Constructor() MyLinkedList {
	return MyLinkedList{}
}


// 获取链表中第 index 个节点的值。如果索引无效，则返回-1。
func (this *MyLinkedList) Get(index int) int {
	if this.num < index + 1 {
		return -1
	}

	p := this.Next
	for	i := 1; i <= index; i ++ {
		p = p.Next
	}

	return p.Val
}


// 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点
func (this *MyLinkedList) AddAtHead(val int)  {
	node := &ListNode{Val: val, Next: this.Next}
	this.Next = node
	this.num++
}

// 将值为 val 的节点追加到链表的最后一个元素
func (this *MyLinkedList) AddAtTail(val int)  {
	node := &ListNode{
		Val: val,
	}

	p := this.Next
	if p == nil {
		this.Next = node
		return
	}

	for p.Next != nil {
		p = p.Next
	}

	p.Next = node
}

// 在链表中的第 index 个节点之前添加值为 val  的节点。
// 如果 index 等于链表的长度，则该节点将附加到链表的末尾.如果 index 大于链表长度，则不会插入节点。
// 如果index小于0，则在头部插入节点。
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	node := &ListNode{
		Val: val,
	}

	if index <= 0 {
		node.Next = this.Next
		this.Next = node
		this.num++
	}else if index == this.num {
		p := this.Next
		for p.Next != nil {
			p = p.Next
		}
		p.Next = node
	}
}

// 如果索引 index 有效，则删除链表中的第 index 个节点。
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index >=this.num || index < 0 {
		return
	}
	// 至少有一个节点
	p := this.Next
	f := p.Next

	if index == 0 {
		this.Next = f
		this.num--
		return
	}

	for i := 1; i < index; i++ {
		f =f.Next
		p=p.Next
	}

	p.Next = f.Next
	f = f.Next
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */