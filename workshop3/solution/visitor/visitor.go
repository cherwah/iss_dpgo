package visitor

import (
	"fmt"
	"solution/tree"
)

type Find_visitor struct {
	Found bool
	Key   int
}

func (v *Find_visitor) Accept(node *tree.Node) {
	if node.Key == v.Key {
		v.Found = true
	}
}

func (v *Find_visitor) Print() {
	if v.Found {
		fmt.Printf("The value %d is found.", v.Key)
	} else {
		fmt.Printf("The value %d is not found.", v.Key)
	}
}
