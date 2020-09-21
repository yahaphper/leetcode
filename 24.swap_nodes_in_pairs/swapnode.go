package swap_nodes_in_pairs

import (
	"leetcode/data_structure"
)

// 递归
func swapPairsOfRecursion(head *data_structure.ListNode) *data_structure.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairsOfRecursion(next.Next)
	next.Next = head
	return next
}

func swapPairsOfIteration(head *data_structure.ListNode) *data_structure.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	h := &data_structure.ListNode{Next: head}
	h.Next = head
	pre := h

	for pre.Next != nil && pre.Next.Next != nil {
		cur := pre.Next
		next := pre.Next.Next
		pre.Next = next
		cur.Next = next.Next
		next.Next = cur
		pre = cur
	}
	return h.Next
}
