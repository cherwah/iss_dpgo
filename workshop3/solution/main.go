package main

import (
	"solution/tree"
	"solution/visitor"
)

func main() {
	bt := create_tree()

	visitor := visitor.Find_visitor{Key: 3}

	bt.Traverse(&visitor)

	visitor.Print()
}

// create a binary-search tree in the memory
func create_tree() tree.Binary_Tree {

	bt := tree.Binary_Tree{}

	keys := []int{9, 3, 6, 2, 8, 15, 13}

	for _, key := range keys {
		node := tree.Node{
			Key: key,
		}
		bt.Add(&node)
	}

	return bt
}
