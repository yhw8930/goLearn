package main

import (
	"fmt"
	"math/rand"
)

type MaxDistanceNode struct {
	Value       int
	Left, Right *MaxDistanceNode
}

// MaxDistance1 返回二叉树任意两节点之间路径所包含的最大节点数；空树返回 0。
// 暴力枚举所有节点对，借助父节点表找到最近公共祖先并计算距离，作为对数器。
// 时间复杂度：O(N^2*H)，最坏 O(N^3)。空间复杂度：O(N)。
func MaxDistance1(head *MaxDistanceNode) int {
	if head == nil {
		return 0
	}
	nodes := make([]*MaxDistanceNode, 0)
	parents := map[*MaxDistanceNode]*MaxDistanceNode{head: nil}
	var collect func(*MaxDistanceNode)
	collect = func(x *MaxDistanceNode) {
		if x == nil {
			return
		}
		nodes = append(nodes, x)
		if x.Left != nil {
			parents[x.Left] = x
		}
		if x.Right != nil {
			parents[x.Right] = x
		}
		collect(x.Left)
		collect(x.Right)
	}
	collect(head)
	ans := 0
	for i := range nodes {
		for j := i; j < len(nodes); j++ {
			ans = max(ans, maxDistanceBetween(parents, nodes[i], nodes[j]))
		}
	}
	return ans
}

func maxDistanceBetween(parents map[*MaxDistanceNode]*MaxDistanceNode, a, b *MaxDistanceNode) int {
	ancestors := make(map[*MaxDistanceNode]struct{})
	for cur := a; cur != nil; cur = parents[cur] {
		ancestors[cur] = struct{}{}
	}
	lca := b
	for {
		if _, ok := ancestors[lca]; ok {
			break
		}
		lca = parents[lca]
	}
	distance := 1
	for cur := a; cur != lca; cur = parents[cur] {
		distance++
	}
	for cur := b; cur != lca; cur = parents[cur] {
		distance++
	}
	return distance
}

type maxDistanceInfo struct{ maxDistance, height int }

// MaxDistance2 使用后序递归求二叉树最大距离。
// 最长路径要么完全位于左树、完全位于右树，要么经过当前节点，取三者最大值。
// 时间复杂度：O(N)。空间复杂度：O(H)。
func MaxDistance2(head *MaxDistanceNode) int { return maxDistanceProcess(head).maxDistance }

func maxDistanceProcess(x *MaxDistanceNode) maxDistanceInfo {
	if x == nil {
		return maxDistanceInfo{}
	}
	left, right := maxDistanceProcess(x.Left), maxDistanceProcess(x.Right)
	return maxDistanceInfo{
		maxDistance: max(max(left.maxDistance, right.maxDistance), left.height+right.height+1),
		height:      max(left.height, right.height) + 1,
	}
}

func maxDistanceGenerate(level, maxLevel, maxValue int) *MaxDistanceNode {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}
	head := &MaxDistanceNode{Value: rand.Intn(maxValue)}
	head.Left = maxDistanceGenerate(level+1, maxLevel, maxValue)
	head.Right = maxDistanceGenerate(level+1, maxLevel, maxValue)
	return head
}

// main 随机生成二叉树，对比节点对暴力枚举和树形 DP 两种最大距离算法。
func main() {
	for i := 0; i < 10000; i++ {
		head := maxDistanceGenerate(1, 4, 100)
		if MaxDistance1(head) != MaxDistance2(head) {
			fmt.Println("Oops!")
			return
		}
	}
	fmt.Println("finish!")
}
