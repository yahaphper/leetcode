package linked_list_cycle

import (
	"leetcode/data_structure"
)

/*
给定一个链表，判断链表中是否有环。
*/


// 使用map进行判断
func hasCycleOfMap(head *data_structure.ListNode) bool {
	if head == nil {
		return false
	}

	m := map[*data_structure.ListNode]bool{}

	for head != nil {
		if _, exist := m[head]; exist {
			return true
		}

		m[head] = true
		head = head.Next
	}

	return false
}

// 使用快慢指针
func hasCycleOfDoublePointer(head *data_structure.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != slow {

		if fast == nil || fast.Next == nil || fast.Next.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
