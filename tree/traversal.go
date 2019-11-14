package tree

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
