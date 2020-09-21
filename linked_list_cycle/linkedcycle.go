package linked_list_cycle

/*
给定一个链表，判断链表中是否有环。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用map进行判断
func hasCycleOfMap(head *ListNode) bool {
	if head == nil {
		return false
	}

	m := map[*ListNode]bool{}

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
func hasCycleOfSpeedAndPointer(head *ListNode) bool {
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
