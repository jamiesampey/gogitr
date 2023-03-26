package bst

import (
	"fmt"
	"log"
	"math/rand"
)

type BST struct {
	root *node
}

type node struct {
	value, balance int
	left           *node
	right          *node
}

func newLeafNode(value int) *node {
	return &node{value, 0, nil, nil}
}

func (bst *BST) PopulateTree(numNodes int) {
	numInserted := 0
	for numInserted < numNodes {
		insertSuccess := bst.insertNode(rand.Intn(100))
		if insertSuccess {
			numInserted++
		}
	}
}

func (bst *BST) insertNode(value int) bool {
	if bst.root == nil {
		log.Printf("Inserting %d at ROOT", value)
		bst.root = newLeafNode(value)
		return true
	}

	currentNode := bst.root
	for {
		if value == currentNode.value {
			logInsertOp(value, currentNode.value)
			return false
		}

		if value < currentNode.value {
			if currentNode.left != nil {
				currentNode.balance--
				currentNode = currentNode.left
				continue
			}

			logInsertOp(value, currentNode.value)
			currentNode.left = newLeafNode(value)
			return true
		}

		if currentNode.right != nil {
			currentNode.balance++
			currentNode = currentNode.right
			continue
		}

		logInsertOp(value, currentNode.value)
		currentNode.right = newLeafNode(value)
		return true
	}
}

func (bst *BST) PrintSortedTree() {
	bst.printSubTree(bst.root)
	fmt.Println()
}

func (bst *BST) printSubTree(n *node) {
	if n.left != nil {
		bst.printSubTree(n.left)
	}

	fmt.Printf("%d ", n.value)

	if n.right != nil {
		bst.printSubTree(n.right)
	}
}

func logInsertOp(newValue, lastValue int) {
	if newValue == lastValue {
		log.Printf("NO OP - BST already contains %d", newValue)
		return
	}

	pos := "LEFT"
	if newValue > lastValue {
		pos = "RIGHT"
	}
	log.Printf("Inserting %d %s of %d", newValue, pos, lastValue)
}
