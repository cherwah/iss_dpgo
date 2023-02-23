package main

import (
	"fmt"
	"strconv"
	"visitor/tree"
)

type sum_visitor struct {
	sum int
}

func (v *sum_visitor) Accept(node *tree.Node) {
	v.sum += node.Key
}

func (v sum_visitor) print() {
	fmt.Println("Sum is " + strconv.Itoa(v.sum))
}

func main() {
	bt := create_tree()

	visitor := sum_visitor{}
	bt.Traverse(&visitor)

	visitor.print()
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
