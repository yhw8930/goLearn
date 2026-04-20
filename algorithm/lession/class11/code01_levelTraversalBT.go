package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func levelTraversalBT(head *Node) {
	if head == nil {
		return
	}
	queue := []*Node{head}
	for len(queue) > 0 {
		node := queue[0]
		queue[0] = nil
		queue = queue[1:]
		fmt.Println(node.Value)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func main() {
	head := &Node{Value: 1}
	head.Left = &Node{Value: 2}
	head.Right = &Node{Value: 3}
	head.Left.Left = &Node{Value: 4}
	head.Left.Right = &Node{Value: 5}
	head.Right.Left = &Node{Value: 6}
	head.Right.Right = &Node{Value: 7}

	levelTraversalBT(head)
	fmt.Println("========")
}
