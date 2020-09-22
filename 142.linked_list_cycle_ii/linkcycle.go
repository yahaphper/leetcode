package linked_list_cycle_ii

import (
	"leetcode/data_structure"
)

/*
	给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
*/

// map
func detectCycleOfMap(head *data_structure.ListNode) *data_structure.ListNode {
	if head == nil {
		return nil
	}

	m := map[*data_structure.ListNode]bool{}

	for head != nil {
		if _, exist := m[head]; exist {
			return head
		}

		m[head] = true
		head = head.Next
	}

	return nil
}
