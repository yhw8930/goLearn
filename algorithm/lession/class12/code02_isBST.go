package main

import (
	"fmt"
	"math/rand"
)

type BSTNode struct {
	Value       int
	Left, Right *BSTNode
}

// IsBST1 判断二叉树是否为严格二叉搜索树：任意节点的左子树值都小于它，右子树值都大于它，不允许重复值。
// 严格搜索二叉树的中序遍历结果严格递增，因此先收集节点，再检查相邻值。
// 时间复杂度：O(N)。空间复杂度：O(N)。
func IsBST1(head *BSTNode) bool {
	nodes := make([]*BSTNode, 0)
	bstInOrder(head, &nodes)
	for i := 1; i < len(nodes); i++ {
		if nodes[i].Value <= nodes[i-1].Value {
			return false
		}
	}
	return true
}

func bstInOrder(head *BSTNode, nodes *[]*BSTNode) {
	if head == nil {
		return
	}
	bstInOrder(head.Left, nodes)
	*nodes = append(*nodes, head)
	bstInOrder(head.Right, nodes)
}

type bstInfo struct {
	isBST    bool
	min, max int
}

// IsBST2 使用递归信息判断严格二叉搜索树。
// 子树返回是否为 BST 及最小、最大值；当前节点要求左右子树均为 BST，且 left.max < x < right.min。
// 时间复杂度：O(N)。空间复杂度：O(H)。
func IsBST2(head *BSTNode) bool {
	if head == nil {
		return true
	}
	return bstProcess(head).isBST
}

func bstProcess(x *BSTNode) *bstInfo {
	if x == nil {
		return nil
	}
	left, right := bstProcess(x.Left), bstProcess(x.Right)
	info := &bstInfo{isBST: true, min: x.Value, max: x.Value}
	if left != nil {
		info.min = min(info.min, left.min)
		info.max = max(info.max, left.max)
		info.isBST = info.isBST && left.isBST && left.max < x.Value
	}
	if right != nil {
		info.min = min(info.min, right.min)
		info.max = max(info.max, right.max)
		info.isBST = info.isBST && right.isBST && right.min > x.Value
	}
	return info
}

func bstGenerate(level, maxLevel, maxValue int) *BSTNode {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}
	head := &BSTNode{Value: rand.Intn(maxValue)}
	head.Left = bstGenerate(level+1, maxLevel, maxValue)
	head.Right = bstGenerate(level+1, maxLevel, maxValue)
	return head
}

// main 随机生成二叉树，对比中序遍历和递归信息两种 BST 判定方法。
func main() {
	for i := 0; i < 10000; i++ {
		head := bstGenerate(1, 4, 100)
		if IsBST1(head) != IsBST2(head) {
			fmt.Println("Oops!")
			return
		}
	}
	fmt.Println("finish!")
}
