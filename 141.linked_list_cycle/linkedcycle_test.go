package linked_list_cycle

import (
	"fmt"
	"leetcode/data_structure"
	"testing"
)

func Test(t *testing.T) {
	l1 := data_structure.ListNode{
		Val: 1,
	}

	fmt.Print(hasCycleOfSpeedAndPointer(&l1))
}
