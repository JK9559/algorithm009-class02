package main

func mergeList(l1, l2 *ListNode) *ListNode {
	preNode := &ListNode{
		Val:  -1,
		Next: nil,
	}
	p := preNode
	// 注意这里的条件 应该是 && 不应该是 ||
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return preNode.Next
}
