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

	m := map[*data_structure.ListNode]int8{}

	for head != nil {
		if _, exist := m[head]; exist {
			return head
		}

		m[head] = 1
		head = head.Next
	}

	return nil
}

// 双指针
func detectCycleOfDoublePointer(head *data_structure.ListNode) *data_structure.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		if fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			for head != fast {
				head = head.Next
				fast = fast.Next
			}
			return head
		}
	}
	return nil
}
