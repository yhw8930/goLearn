package main

import (
	"fmt"
	"goLearn/tree"

	"golang.org/x/tools/container/intsets"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func testSparse() {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(10000))
}
func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CrateNode(2)
	fmt.Println(root)

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	root.Print()
	root.SetValue(100)

	var pRoot *tree.Node
	pRoot.SetValue(200)
	pRoot = &root
	pRoot.SetValue(300)
	pRoot.Print()
	fmt.Println()

	root.Travserse()
	fmt.Println()
	tree.Travserse(&root)
	fmt.Println()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	testSparse()
}
