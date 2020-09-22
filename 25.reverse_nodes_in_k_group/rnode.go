package reverse_nodes_in_k_group

import (
	"leetcode/data_structure"
)

func reverseKGroup(head *data_structure.ListNode, k int) *data_structure.ListNode {
	hair := &data_structure.ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = reverseList(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func reverseList(head, tail *data_structure.ListNode) (*data_structure.ListNode, *data_structure.ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}
