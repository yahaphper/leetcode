package linked_list_cycle

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	l1 := ListNode{
		Val: 1,
	}

	fmt.Print(hasCycleOfSpeedAndPointer(&l1))
}
