package main

import (
	"jamiesampey.com/gogitr/pkg/bst"
)

func main() {
	tree := bst.BST{}
	tree.PopulateTree(10)
	tree.PrintSortedTree()
}
