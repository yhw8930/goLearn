package main

import (
	"fmt"
	"math/rand"
)

type BalancedNode struct {
	Value       int
	Left, Right *BalancedNode
}

// IsBalanced1 判断一棵二叉树是否为平衡二叉树：每个节点的左右子树高度差都不超过 1。
// 递归求高度，并通过共享标记记录是否已经发现不平衡节点。
// 时间复杂度：O(N)。空间复杂度：O(H)。
func IsBalanced1(head *BalancedNode) bool {
	ans := true
	balancedHeight1(head, &ans)
	return ans
}

func balancedHeight1(head *BalancedNode, ans *bool) int {
	if !*ans || head == nil {
		return 0
	}
	left := balancedHeight1(head.Left, ans)
	right := balancedHeight1(head.Right, ans)
	if left-right > 1 || right-left > 1 {
		*ans = false
	}
	return max(left, right) + 1
}

type balancedInfo struct {
	isBalanced bool
	height     int
}

// IsBalanced2 使用二叉树递归套路判断平衡性。
// 每棵子树同时返回高度和平衡状态，当前节点合并左右信息即可。
// 时间复杂度：O(N)。空间复杂度：O(H)。
func IsBalanced2(head *BalancedNode) bool { return balancedProcess(head).isBalanced }

func balancedProcess(x *BalancedNode) balancedInfo {
	if x == nil {
		return balancedInfo{true, 0}
	}
	left, right := balancedProcess(x.Left), balancedProcess(x.Right)
	diff := left.height - right.height
	return balancedInfo{
		isBalanced: left.isBalanced && right.isBalanced && diff >= -1 && diff <= 1,
		height:     max(left.height, right.height) + 1,
	}
}

func balancedGenerate(level, maxLevel, maxValue int) *BalancedNode {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}
	head := &BalancedNode{Value: rand.Intn(maxValue)}
	head.Left = balancedGenerate(level+1, maxLevel, maxValue)
	head.Right = balancedGenerate(level+1, maxLevel, maxValue)
	return head
}

// main 随机生成二叉树，对比共享标记和递归信息两种平衡性判定方法。
func main() {
	for i := 0; i < 10000; i++ {
		head := balancedGenerate(1, 5, 100)
		if IsBalanced1(head) != IsBalanced2(head) {
			fmt.Println("Oops!")
			return
		}
	}
	fmt.Println("finish!")
}
