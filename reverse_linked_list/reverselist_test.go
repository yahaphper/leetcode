package reverse_linked_list

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	list := ListNode{Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	l := reverseList(&list)

	fmt.Print(l)
}
