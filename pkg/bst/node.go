package bst

import (
	"fmt"
	"strings"
)

type Node struct {
	value, height int
	left          *Node
	right         *Node
}

func (node *Node) Insert(val int) (*Node, error) {

	if node == nil {
		return &Node{value: val, height: 1}, nil
	}

	if node.value == val {
		return node, fmt.Errorf("node %d already exists", val)
	}

	var err error
	if val < node.value {
		node.left, err = node.left.Insert(val)
	} else {
		node.right, err = node.right.Insert(val)
	}

	if err == nil {
		node.height = node.maxChildHeight() + 1
		return node.rebalance(), nil
	}

	return node, err
}

func (node *Node) Balance() int {
	if node == nil {
		return 0
	}
	return node.rightHeight() - node.leftHeight()
}

func (node *Node) Print(i, parentValue int) {
	if node == nil {
		return
	}

	var position string
	if parentValue >= 0 {
		if node.value < parentValue {
			position = fmt.Sprintf("left of %d", parentValue)
		} else {
			position = fmt.Sprintf("right of %d", parentValue)
		}
	}

	fmt.Printf("%s%d[%d,%d] %s\n", strings.Repeat(" ", i*4), node.value,
		node.Balance(), node.height, position)

	node.left.Print(i+1, node.value)
	node.right.Print(i+1, node.value)
}

func (node *Node) rebalance() *Node {
	switch {

	// shift LEFT cases
	case node.Balance() > 1 && node.right.Balance() == 1:
		// node is too heavy on the right, and it's right child is weighted right
		return node.shiftLeft()
	case node.Balance() > 1 && node.right.Balance() == -1:
		// node is too heavy on the right, and it's right child is weighted left
		node.right = node.right.shiftRight()
		node = node.shiftLeft()
		node.height = node.maxChildHeight() + 1
		return node

	// shift RIGHT cases
	case node.Balance() < -1 && node.left.Balance() == -1:
		// node is too heavy on the left, and it's left child is weighted left
		return node.shiftRight()
	case node.Balance() < -1 && node.left.Balance() == 1:
		// node is too heavy on the left, and it's left child is weighted right
		node.left = node.left.shiftLeft()
		node = node.shiftRight()
		node.height = node.maxChildHeight() + 1
		return node
	}

	return node // no rebalance needed
}

// The right child SHIFTS UP AND LEFT to become the parent node
// The orig parent node SHIFTS DOWN AND LEFT to become the new parent's left node
// The orig right child's left must move over to the new left child's right
func (node *Node) shiftLeft() *Node {
	fmt.Printf("Shift: %d[%d,%d] moves down and left\n", node.value, node.Balance(), node.height)

	origRight := node.right
	node.right = origRight.left
	origRight.left = node

	node.height = node.maxChildHeight() + 1
	origRight.height = origRight.maxChildHeight() + 1

	return origRight // origRight becomes the new parent
}

// The left child shifts UP and RIGHT to become the parent node
// The orig parent node shifts DOWN and RIGHT to become the new parent's right node
// The orig left child's right must move over to the new right child's left
func (node *Node) shiftRight() *Node {
	fmt.Printf("Shift: %d[%d,%d] moves down and right\n", node.value, node.Balance(), node.height)

	origLeft := node.left
	node.left = origLeft.right
	origLeft.right = node

	node.height = node.maxChildHeight() + 1
	origLeft.height = origLeft.maxChildHeight() + 1

	return origLeft // origLeft becomes the new parent
}

func (node *Node) maxChildHeight() int {
	leftHeight := node.leftHeight()
	rightHeight := node.rightHeight()

	if leftHeight > rightHeight {
		return leftHeight
	}
	return rightHeight
}

func (node *Node) leftHeight() int {
	if node.left == nil {
		return 0
	}
	return node.left.height
}

func (node *Node) rightHeight() int {
	if node.right == nil {
		return 0
	}
	return node.right.height
}
