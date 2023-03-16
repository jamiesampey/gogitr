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
	value float32
	left  *node
	right *node
}

func (bst *BST) PopulateTree(numNodes int) {
	for i := 0; i < numNodes; i++ {
		bst.insertNode(rand.Float32() * 100)
	}
}

func (bst *BST) insertNode(value float32) {
	if bst.root == nil {
		log.Printf("Inserting %.2f at ROOT", value)
		bst.root = &node{value, nil, nil}
		return
	}

	currentNode := bst.root
	for {
		if value == currentNode.value {
			logInsertOp(value, currentNode.value)
			return
		}

		if value < currentNode.value {
			if currentNode.left != nil {
				currentNode = currentNode.left
				continue
			}

			logInsertOp(value, currentNode.value)
			currentNode.left = &node{value, nil, nil}
			return
		}

		if currentNode.right != nil {
			currentNode = currentNode.right
			continue
		}

		logInsertOp(value, currentNode.value)
		currentNode.right = &node{value, nil, nil}
		return
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

	fmt.Printf("%.2f ", n.value)

	if n.right != nil {
		bst.printSubTree(n.right)
	}
}

func logInsertOp(newValue, lastValue float32) {
	if newValue == lastValue {
		log.Printf("NO OP - BST already contains %.2f", newValue)
		return
	}

	pos := "LEFT"
	if newValue > lastValue {
		pos = "RIGHT"
	}
	log.Printf("Inserting %.2f %s of %.2f", newValue, pos, lastValue)
}
