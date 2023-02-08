package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	// 通过 map 来存储一条链表
	m := make(map[*ListNode]struct{}, 0)
	cur := headA
	for cur !=nil {
		m[cur]	= struct{}{}
		cur = cur.Next
	}

	cur = headB
	for cur != nil {
		if _,has := m[cur]; has {
			return cur
		}
		cur = cur.Next
	}

	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 获取 A 链表长度链表长度
	lenA,lenB :=0,0
	n := 0
	cur := headA
	for cur !=nil {
		lenA++
		cur = cur.Next
	}

	cur = headB
	for cur != nil {
		lenB++
		cur = cur.Next
	}

	long,short := headA, headB
	if lenA > lenB{
		n = lenA-lenB
	}else {
		long,short =  headB,headA
		n = lenB-lenA
	}

	for long != nil {
		n--
		if n >= 0 {
			long = long.Next
			continue
		}
		if long == short {
			return long
		}
		long=long.Next
		short = short.Next
	}

	return nil
}