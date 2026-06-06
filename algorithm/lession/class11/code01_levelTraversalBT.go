package main

import "fmt"

// 题目：按层从上到下、从左到右遍历二叉树。
// 层序遍历需要先处理离根更近的节点，符合队列先进先出的特点。
// 核心思路：根节点先入队，每弹出一个节点就打印它，并把它的左、右孩子依次入队。
// 队列为空时说明所有可达节点都已经按层处理完。
// 时间复杂度：O(N)。
// 空间复杂度：O(W)，W 为树的最大宽度。

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
