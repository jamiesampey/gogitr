package bst

import (
	"fmt"
	"log"
	"math/rand"
)

type BST struct {
	root *Node
}

func (bst *BST) PopulateTree(numNodes int) {
	numInserted := 0
	var err error
	for numInserted < numNodes {
		num := rand.Intn(100)
		bst.root, err = bst.root.Insert(num)
		if err == nil {
			numInserted++
			log.Printf("Inserted %d | Tree height %d | Tree balance %d",
				num, bst.root.height, bst.root.Balance())
		} else {
			fmt.Println(err)
		}
	}
}

func (bst *BST) PrintSortedTree() {
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

func logInsertOp(newValue, lastValue int) {
	if newValue == lastValue {

		return
	}

	pos := "LEFT"
	if newValue > lastValue {
		pos = "RIGHT"
	}
	log.Printf("Inserting %d %s of %d", newValue, pos, lastValue)
}
