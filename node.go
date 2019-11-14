package main

import (
	"fmt"
)

type treenode struct {
	value       int
	left, right *treenode
}

func (node treenode) print() {
	fmt.Print(node.value, " ")
}

func createNode(value int) *treenode {
	return &treenode{value: value}
}

func (node *treenode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.value = value
}

func (node *treenode) travserse() {
	if node == nil {
		return
	}
	node.left.travserse()
	node.print()
	node.right.travserse()
}

func traverse(node *treenode) {
	if node != nil {
		traverse(node.left)
		node.print()
		traverse(node.right)
	}
}
func main() {
	var root treenode
	root = treenode{value: 3}
	root.left = &treenode{}
	root.right = &treenode{5, nil, nil}
	root.right.left = new(treenode)
	root.left.right = createNode(2)
	fmt.Println(root)

	root.right.left.setValue(4)
	root.right.left.print()
	fmt.Println()

	root.print()
	root.setValue(100)

	var pRoot *treenode
	pRoot.setValue(200)
	pRoot = &root
	pRoot.setValue(300)
	pRoot.print()
	fmt.Println()

	root.travserse()

	traverse(&root)
	/*nodes := []treenode{
		{value: 3},
		{},
		{6, nil, nil},
	}
	fmt.Println(nodes)*/
}
