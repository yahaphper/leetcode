package reverse_linked_list

import (
	"leetcode/data_structure"
)

/*
	反转一个单链表。

	示例:

	输入: 1->2->3->4->5->NULL
	输出: 5->4->3->2->1->NULL
*/

func reverseList(head *data_structure.ListNode) *data_structure.ListNode {
	var pre *data_structure.ListNode = nil

	for head != nil {
		head.Next, head, pre = pre, head.Next, head
	}

	return pre
}
