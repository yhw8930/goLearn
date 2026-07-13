package main

import (
	"fmt"
	"math/rand"
)

type LowestAncestorNode struct {
	Value       int
	Left, Right *LowestAncestorNode
}

// LowestAncestor1 返回二叉树中节点 a、b 的最低公共祖先，即同时包含二者的最深节点。
// 前提：a、b 都属于以 head 为根的树。先建立父节点表，再记录 a 到根的祖先集合，沿 b 向上找到首个交点。
// 时间复杂度：O(N)。空间复杂度：O(N)。
func LowestAncestor1(head, a, b *LowestAncestorNode) *LowestAncestorNode {
	if head == nil {
		return nil
	}
	parents := map[*LowestAncestorNode]*LowestAncestorNode{head: nil}
	lowestAncestorFillParents(head, parents)
	ancestors := make(map[*LowestAncestorNode]struct{})
	for cur := a; cur != nil; cur = parents[cur] {
		ancestors[cur] = struct{}{}
	}
	for cur := b; cur != nil; cur = parents[cur] {
		if _, ok := ancestors[cur]; ok {
			return cur
		}
	}
	return nil
}

func lowestAncestorFillParents(head *LowestAncestorNode, parents map[*LowestAncestorNode]*LowestAncestorNode) {
	if head.Left != nil {
		parents[head.Left] = head
		lowestAncestorFillParents(head.Left, parents)
	}
	if head.Right != nil {
		parents[head.Right] = head
		lowestAncestorFillParents(head.Right, parents)
	}
}

type lowestAncestorInfo struct {
	findA, findB bool
	ans          *LowestAncestorNode
}

// LowestAncestor2 使用后序递归寻找最低公共祖先。
// 每棵子树返回是否找到 a、b 以及已经确定的答案；首个同时找到二者的节点就是最低公共祖先。
// 时间复杂度：O(N)。空间复杂度：O(H)。
func LowestAncestor2(head, a, b *LowestAncestorNode) *LowestAncestorNode {
	return lowestAncestorProcess(head, a, b).ans
}

func lowestAncestorProcess(x, a, b *LowestAncestorNode) lowestAncestorInfo {
	if x == nil {
		return lowestAncestorInfo{}
	}
	left := lowestAncestorProcess(x.Left, a, b)
	right := lowestAncestorProcess(x.Right, a, b)
	findA := x == a || left.findA || right.findA
	findB := x == b || left.findB || right.findB
	ans := left.ans
	if ans == nil {
		ans = right.ans
	}
	if ans == nil && findA && findB {
		ans = x
	}
	return lowestAncestorInfo{findA, findB, ans}
}

func lowestAncestorGenerate(level, maxLevel, maxValue int) *LowestAncestorNode {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}
	head := &LowestAncestorNode{Value: rand.Intn(maxValue)}
	head.Left = lowestAncestorGenerate(level+1, maxLevel, maxValue)
	head.Right = lowestAncestorGenerate(level+1, maxLevel, maxValue)
	return head
}

func lowestAncestorPick(head *LowestAncestorNode) *LowestAncestorNode {
	if head == nil {
		return nil
	}
	nodes := make([]*LowestAncestorNode, 0)
	var collect func(*LowestAncestorNode)
	collect = func(x *LowestAncestorNode) {
		if x == nil {
			return
		}
		nodes = append(nodes, x)
		collect(x.Left)
		collect(x.Right)
	}
	collect(head)
	return nodes[rand.Intn(len(nodes))]
}

// main 随机生成二叉树及其中两个节点，对比父节点表和递归信息两种最低公共祖先算法。
func main() {
	for i := 0; i < 10000; i++ {
		head := lowestAncestorGenerate(1, 4, 100)
		a, b := lowestAncestorPick(head), lowestAncestorPick(head)
		if LowestAncestor1(head, a, b) != LowestAncestor2(head, a, b) {
			fmt.Println("Oops!")
			return
		}
	}
	fmt.Println("finish!")
}
