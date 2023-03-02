package main

import (
	"visitor/tree"
	"visitor/visitor"
)

func main() {
	bt := create_tree()

	visitor := visitor.Sum_visitor{}
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
