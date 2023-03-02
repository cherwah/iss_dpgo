package tree

type Acceptable interface {
	Accept(Node *Node)
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type Binary_Tree struct {
	Root *Node
}

func (bt *Binary_Tree) Add(node *Node) {
	if bt.Root == nil {
		bt.Root = node
		return
	}

	bt.add_recursive(bt.Root, node)
}

func (bt *Binary_Tree) add_recursive(curr_node *Node, new_node *Node) {
	if curr_node.Key < new_node.Key {
		if curr_node.Right == nil {
			curr_node.Right = new_node
		} else {
			bt.add_recursive(curr_node.Right, new_node)
		}
	} else {
		if curr_node.Left == nil {
			curr_node.Left = new_node
		} else {
			bt.add_recursive(curr_node.Left, new_node)
		}
	}
}

func (bt *Binary_Tree) Traverse(visitor Acceptable) {
	bt.depth_first(bt.Root, visitor)
}

func (bt *Binary_Tree) depth_first(curr_node *Node, visitor Acceptable) {
	if curr_node != nil {
		visitor.Accept(curr_node)
	}

	if curr_node.Left != nil {
		bt.depth_first(curr_node.Left, visitor)
	}

	if curr_node.Right != nil {
		bt.depth_first(curr_node.Right, visitor)
	}
}
