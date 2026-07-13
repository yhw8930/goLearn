package main

import (
	"fmt"
	"math/rand"
)

type FullNode struct {
	Value       int
	Left, Right *FullNode
}

// IsFull1 判断二叉树是否为满二叉树：每一层的节点数都达到最大值。
// 高度为 h 的满二叉树恰有 2^h-1 个节点，分别统计高度和节点数后验证该公式。
// 时间复杂度：O(N)。空间复杂度：O(H)。
func IsFull1(head *FullNode) bool {
	if head == nil {
		return true
	}
	height := fullHeight(head)
	return (1<<height)-1 == fullNodes(head)
}

func fullHeight(head *FullNode) int {
	if head == nil {
		return 0
	}
	return max(fullHeight(head.Left), fullHeight(head.Right)) + 1
}

func fullNodes(head *FullNode) int {
	if head == nil {
		return 0
	}
	return fullNodes(head.Left) + fullNodes(head.Right) + 1
}

type fullInfo struct{ height, nodes int }

// IsFull2 在一次后序递归中同时统计整棵树的高度和节点数，再验证 2^h-1。
// 时间复杂度：O(N)。空间复杂度：O(H)。
func IsFull2(head *FullNode) bool {
	info := fullProcess(head)
	return (1<<info.height)-1 == info.nodes
}

func fullProcess(head *FullNode) fullInfo {
	if head == nil {
		return fullInfo{}
	}
	left, right := fullProcess(head.Left), fullProcess(head.Right)
	return fullInfo{max(left.height, right.height) + 1, left.nodes + right.nodes + 1}
}

func fullGenerate(level, maxLevel, maxValue int) *FullNode {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}
	head := &FullNode{Value: rand.Intn(maxValue)}
	head.Left = fullGenerate(level+1, maxLevel, maxValue)
	head.Right = fullGenerate(level+1, maxLevel, maxValue)
	return head
}

// main 随机生成二叉树，对比两种满二叉树判定方法。
func main() {
	for i := 0; i < 10000; i++ {
		head := fullGenerate(1, 5, 100)
		if IsFull1(head) != IsFull2(head) {
			fmt.Println("Oops!")
			return
		}
	}
	fmt.Println("finish!")
}
