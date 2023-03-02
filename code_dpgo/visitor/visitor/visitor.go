package visitor

import (
	"fmt"
	"strconv"
	"visitor/tree"
)

type Sum_visitor struct {
	Sum int
}

func (v *Sum_visitor) Accept(node *tree.Node) {
	v.Sum += node.Key
}

func (v *Sum_visitor) Print() {
	fmt.Println("Sum is " + strconv.Itoa(v.Sum))
}
