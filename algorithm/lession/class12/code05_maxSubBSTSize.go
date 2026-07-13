package main

import (
	"fmt"
	"math/rand"
)

type MaxSubBSTSizeNode struct {
	Value       int
	Left, Right *MaxSubBSTSizeNode
}

func maxSubBSTGetSize(head *MaxSubBSTSizeNode) int {
	if head == nil {
		return 0
	}
	nodes := make([]*MaxSubBSTSizeNode, 0)
	var inOrder func(*MaxSubBSTSizeNode)
	inOrder = func(x *MaxSubBSTSizeNode) {
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

// MaxSubBSTSize1 返回二叉树中“节点数最多的搜索二叉子树”的节点数；搜索树要求值严格有序。
// 暴力检查当前整棵子树是否为 BST；若不是，再递归比较左右子树，作为对数器。
// 时间复杂度：O(N^2)。空间复杂度：O(N)。
func MaxSubBSTSize1(head *MaxSubBSTSizeNode) int {
	if head == nil {
		return 0
	}
	if size := maxSubBSTGetSize(head); size != 0 {
		return size
	}
	return max(MaxSubBSTSize1(head.Left), MaxSubBSTSize1(head.Right))
}

type maxSubBSTSizeInfo struct {
	maxBSTSize int
	allSize    int
	min, max   int
}

// MaxSubBSTSize2 使用后序递归一次求出最大搜索二叉子树大小。
// 子树返回总节点数、最大 BST 大小及值域；仅当左右整棵都是 BST 且 left.max < x < right.min 时才能合并。
// 时间复杂度：O(N)。空间复杂度：O(H)。
func MaxSubBSTSize2(head *MaxSubBSTSizeNode) int {
	if head == nil {
		return 0
	}
	return maxSubBSTSizeProcess(head).maxBSTSize
}

func maxSubBSTSizeProcess(x *MaxSubBSTSizeNode) *maxSubBSTSizeInfo {
	if x == nil {
		return nil
	}
	left, right := maxSubBSTSizeProcess(x.Left), maxSubBSTSizeProcess(x.Right)
	info := &maxSubBSTSizeInfo{maxBSTSize: 1, allSize: 1, min: x.Value, max: x.Value}
	if left != nil {
		info.allSize += left.allSize
		info.min = min(info.min, left.min)
		info.max = max(info.max, left.max)
		info.maxBSTSize = left.maxBSTSize
	}
	if right != nil {
		info.allSize += right.allSize
		info.min = min(info.min, right.min)
		info.max = max(info.max, right.max)
		info.maxBSTSize = max(info.maxBSTSize, right.maxBSTSize)
	}
	leftBST := left == nil || left.maxBSTSize == left.allSize
	rightBST := right == nil || right.maxBSTSize == right.allSize
	leftLess := left == nil || left.max < x.Value
	rightMore := right == nil || right.min > x.Value
	if leftBST && rightBST && leftLess && rightMore {
		info.maxBSTSize = info.allSize
	}
	return info
}

func maxSubBSTSizeGenerate(level, maxLevel, maxValue int) *MaxSubBSTSizeNode {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}
	head := &MaxSubBSTSizeNode{Value: rand.Intn(maxValue)}
	head.Left = maxSubBSTSizeGenerate(level+1, maxLevel, maxValue)
	head.Right = maxSubBSTSizeGenerate(level+1, maxLevel, maxValue)
	return head
}

// main 随机生成二叉树，对比暴力方法和递归信息方法求出的最大 BST 子树大小。
func main() {
	for i := 0; i < 10000; i++ {
		head := maxSubBSTSizeGenerate(1, 4, 100)
		if MaxSubBSTSize1(head) != MaxSubBSTSize2(head) {
			fmt.Println("Oops!")
			return
		}
	}
	fmt.Println("finish!")
}
