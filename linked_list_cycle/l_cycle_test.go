package linked_list_cycle

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	fmt.Print(hasCycle(&l1))
}
