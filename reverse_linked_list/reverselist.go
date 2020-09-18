package reverse_linked_list

/*
	反转一个单链表。

	示例:

	输入: 1->2->3->4->5->NULL
	输出: 5->4->3->2->1->NULL
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil

	for head != nil {
		// next := head.Next
		//
		// head.Next = pre
		// pre = head
		// head = next

		head.Next, head, pre = pre, head.Next, head
	}

	return pre
}
