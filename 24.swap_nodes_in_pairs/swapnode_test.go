package swap_nodes_in_pairs

import (
	"fmt"
	"leetcode/data_structure"
	"testing"
)

func Test(t *testing.T) {
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

	fmt.Print(swapPairsOfIteration(&list))
}
