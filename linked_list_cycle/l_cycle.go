package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	l_map := map[*ListNode]bool{}

	for head != nil {
		if _, ok := l_map[head]; ok {
			return true
		}

		l_map[head] = true
		head = head.Next
	}

	return false
}
