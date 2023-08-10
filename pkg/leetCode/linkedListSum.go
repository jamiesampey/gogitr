package leetCode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	retPtr := new(ListNode)
	carryover := 0

	for sumNode := retPtr; !done(l1, l2, carryover); sumNode = sumNode.Next {
		sum := carryover

		if l1 != nil {
			fmt.Printf("l1.Val=%d ", l1.Val)
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			fmt.Printf("l2.Val=%d ", l2.Val)
			sum += l2.Val
			l2 = l2.Next
		}

		if sum >= 10 {
			sum = sum % 10
			carryover = 1
		} else {
			carryover = 0
		}

		sumNode.Val = sum
		fmt.Printf("sumNode.Val=%d\n", sumNode.Val)

		if !done(l1, l2, carryover) {
			sumNode.Next = new(ListNode)
		}
	}

	return retPtr
}

func done(l1, l2 *ListNode, carryover int) bool {
	return l1 == nil && l2 == nil && carryover == 0
}
