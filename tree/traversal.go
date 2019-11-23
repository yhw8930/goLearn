package tree

import "fmt"

func (node *Node) Travserse() {
	if node == nil {
		return
	}
	node.Left.Travserse()
	node.Print()
	node.Right.Travserse()
}

func Travserse(node *Node) {
	if node != nil {
		Travserse(node.Left)
		node.Print()
		Travserse(node.Right)
	}
}

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
