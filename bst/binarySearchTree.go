package bst

import (
	"fmt"
	"math/rand"
)

type BST struct {
	root *Node
}

func (bst *BST) PopulateTree(numNodes int) {
	nodeCount := 0
	var err error
	for nodeCount < numNodes {
		value := rand.Intn(100)
		bst.root, err = bst.root.Insert(value)
		if err == nil {
			nodeCount++
			fmt.Println()
			fmt.Printf("Insert #%d: %d, Tree:\n", nodeCount, value)
			bst.root.Print(0, -1)
		} else {
			fmt.Println(err)
		}
	}
}

func (bst *BST) PrintSortedTree() {
	fmt.Print("\nRecursive Tree Sort: ")
	bst.printSubTree(bst.root)
	fmt.Println()
}

func (bst *BST) printSubTree(n *Node) {
	if n.left != nil {
		bst.printSubTree(n.left)
	}

	fmt.Printf("%d ", n.value)

	if n.right != nil {
		bst.printSubTree(n.right)
	}
}
