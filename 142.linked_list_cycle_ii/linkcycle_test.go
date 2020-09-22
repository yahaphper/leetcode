package linked_list_cycle_ii

import (
	"fmt"
	"leetcode/data_structure"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	l1 := data_structure.ListNode{
		Val: 3,
		Next: &data_structure.ListNode{
			Val: 2,
			Next: &data_structure.ListNode{
				Val: 0,
				Next: &data_structure.ListNode{
					Val: -4,
				},
			},
		},
	}
	fmt.Print(detectCycleOfDoublePointer(&l1))
}
