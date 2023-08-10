package main

import (
	"fmt"
	"jamiesampey.com/gogitr/pkg/leetCode"
)

func main() {
	l1 := &leetCode.ListNode{
		Val: 2,
		Next: &leetCode.ListNode{
			Val: 4,
			Next: &leetCode.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &leetCode.ListNode{
		Val: 5,
		Next: &leetCode.ListNode{
			Val: 6,
			Next: &leetCode.ListNode{
				Val: 4,
				Next: &leetCode.ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	lt := leetCode.AddTwoNumbers(l1, l2)

	for lt != nil {
		fmt.Printf("%d ", lt.Val)
		lt = lt.Next
	}
}
