package main

import (
	"jamiesampey.com/gogitr/bst"
)

func main() {
	tree := bst.BST{}
	tree.PopulateTree(50)
	tree.PrintSortedTree()
}
