package main

import (
	"jamiesampey.com/gogitr/bst"
)

func main() {
	tree := bst.BST{}
	tree.PopulateTree(10)
	tree.PrintSortedTree()
}
