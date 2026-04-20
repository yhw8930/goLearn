package main

import "fmt"

type Code06Node struct {
	value  int
	left   *Code06Node
	right  *Code06Node
	parent *Code06Node
}

func getSuccessorNode(node *Code06Node) *Code06Node {
	if node == nil {
		return node
	}
	if node.right != nil {
		return getLeftMost(node.right)
	} else {
		parent := node.parent
		for parent != nil && parent.right == node {
			node = parent
			parent = node.parent
		}
		return parent
	}
}

func getLeftMost(node *Code06Node) *Code06Node {
	if node == nil {
		return node
	}
	for node.left != nil {
		node = node.left
	}
	return node
}

func main() {
	head := &Code06Node{value: 6}
	head.parent = nil
	head.left = &Code06Node{value: 3}
	head.left.parent = head
	head.left.left = &Code06Node{value: 1}
	head.left.left.parent = head.left
	head.left.left.right = &Code06Node{value: 2}
	head.left.left.right.parent = head.left.left
	head.left.right = &Code06Node{value: 4}
	head.left.right.parent = head.left
	head.left.right.right = &Code06Node{value: 5}
	head.left.right.right.parent = head.left.right
	head.right = &Code06Node{value: 9}
	head.right.parent = head
	head.right.left = &Code06Node{value: 8}
	head.right.left.parent = head.right
	head.right.left.left = &Code06Node{value: 7}
	head.right.left.left.parent = head.right.left
	head.right.right = &Code06Node{value: 10}
	head.right.right.parent = head.right

	test := head.left.left
	fmt.Println(test.value, " next:", getSuccessorNode(test).value)
	test = head.left.left.right
	fmt.Println(test.value, " next:", getSuccessorNode(test).value)
	test = head.left
	fmt.Println(test.value, " next:", getSuccessorNode(test).value)
	test = head.left.right
	fmt.Println(test.value, " next:", getSuccessorNode(test).value)
	test = head.left.right.right
	fmt.Println(test.value, " next:", getSuccessorNode(test).value)
	test = head
	fmt.Println(test.value, " next:", getSuccessorNode(test).value)
	test = head.right.left.left
	fmt.Println(test.value, " next:", getSuccessorNode(test).value)
	test = head.right.left
	fmt.Println(test.value, " next:", getSuccessorNode(test).value)
	test = head.right
	fmt.Println(test.value, " next:", getSuccessorNode(test).value)
	test = head.right.right
	fmt.Println(test.value, " next:", getSuccessorNode(test))
}
