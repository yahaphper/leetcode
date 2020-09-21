package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	lMap := map[*ListNode]bool{}

	for head != nil {
		if _, ok := lMap[head]; ok {
			return true
		}

		lMap[head] = true
		head = head.Next
	}

	return false
}
