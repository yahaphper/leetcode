package reverse_linked_list

import (
	"fmt"
	"leetcode/data_structure"
	"testing"
)

func TestReverseList(t *testing.T) {
	list := data_structure.ListNode{
		Val: 1,
		Next: &data_structure.ListNode{
			Val: 2,
			Next: &data_structure.ListNode{
				Val: 3,
				Next: &data_structure.ListNode{
					Val: 4,
					Next: &data_structure.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Print(reverseList(&list))
}
