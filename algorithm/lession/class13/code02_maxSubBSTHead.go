package main

import (
	"fmt"
	"math/rand"
)

type MaxSubBSTHeadNode struct {
	Value       int
	Left, Right *MaxSubBSTHeadNode
}

func maxSubBSTHeadGetSize(head *MaxSubBSTHeadNode) int {
	if head == nil {
		return 0
	}
	nodes := make([]*MaxSubBSTHeadNode, 0)
	var inOrder func(*MaxSubBSTHeadNode)
	inOrder = func(x *MaxSubBSTHeadNode) {
		if x == nil {
			return
		}
		inOrder(x.Left)
		nodes = append(nodes, x)
		inOrder(x.Right)
	}
	inOrder(head)
	for i := 1; i < len(nodes); i++ {
		if nodes[i].Value <= nodes[i-1].Value {
			return 0
		}
	}
	return len(nodes)
}

// MaxSubBSTHead1 返回二叉树中节点数最多的严格搜索二叉子树的头节点；若并列，选择左侧答案。
// 暴力判断当前整棵子树是否为 BST，否则递归比较左右答案，作为对数器。
// 时间复杂度：O(N^2)。空间复杂度：O(N)。
func MaxSubBSTHead1(head *MaxSubBSTHeadNode) *MaxSubBSTHeadNode {
	if head == nil || maxSubBSTHeadGetSize(head) != 0 {
		return head
	}
	left, right := MaxSubBSTHead1(head.Left), MaxSubBSTHead1(head.Right)
	if maxSubBSTHeadGetSize(left) >= maxSubBSTHeadGetSize(right) {
		return left
	}
	return right
}

type maxSubBSTHeadInfo struct {
	head     *MaxSubBSTHeadNode
	size     int
	min, max int
}

// MaxSubBSTHead2 使用后序递归返回最大搜索二叉子树的头节点。
// 子树带回最大 BST 的头、大小和值域；当左右整棵都是 BST 且 left.max < x < right.min 时合并为更大 BST。
// 时间复杂度：O(N)。空间复杂度：O(H)。
func MaxSubBSTHead2(head *MaxSubBSTHeadNode) *MaxSubBSTHeadNode {
	if head == nil {
		return nil
	}
	return maxSubBSTHeadProcess(head).head
}

func maxSubBSTHeadProcess(x *MaxSubBSTHeadNode) *maxSubBSTHeadInfo {
	if x == nil {
		return nil
	}
	left, right := maxSubBSTHeadProcess(x.Left), maxSubBSTHeadProcess(x.Right)
	info := &maxSubBSTHeadInfo{min: x.Value, max: x.Value}
	if left != nil {
		info.head, info.size = left.head, left.size
		info.min, info.max = min(info.min, left.min), max(info.max, left.max)
	}
	if right != nil {
		if right.size > info.size {
			info.head, info.size = right.head, right.size
		}
		info.min, info.max = min(info.min, right.min), max(info.max, right.max)
	}
	leftWholeBST := left == nil || left.head == x.Left
	rightWholeBST := right == nil || right.head == x.Right
	leftLess := left == nil || left.max < x.Value
	rightMore := right == nil || right.min > x.Value
	if leftWholeBST && rightWholeBST && leftLess && rightMore {
		info.head = x
		info.size = 1
		if left != nil {
			info.size += left.size
		}
		if right != nil {
			info.size += right.size
		}
	}
	return info
}

func maxSubBSTHeadGenerate(level, maxLevel, maxValue int) *MaxSubBSTHeadNode {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}
	head := &MaxSubBSTHeadNode{Value: rand.Intn(maxValue)}
	head.Left = maxSubBSTHeadGenerate(level+1, maxLevel, maxValue)
	head.Right = maxSubBSTHeadGenerate(level+1, maxLevel, maxValue)
	return head
}

// main 随机生成二叉树，对比暴力方法和递归信息方法返回的最大 BST 子树头节点。
func main() {
	for i := 0; i < 10000; i++ {
		head := maxSubBSTHeadGenerate(1, 4, 100)
		if MaxSubBSTHead1(head) != MaxSubBSTHead2(head) {
			fmt.Println("Oops!")
			return
		}
	}
	fmt.Println("finish!")
}
