package main

import (
	"fmt"
	"math/rand"
)

type CBTNode struct {
	Value       int
	Left, Right *CBTNode
}

// IsCBT1 判断二叉树是否为完全二叉树：除最后一层外都满，最后一层节点从左到右连续排列。
// 层序遍历中，节点不能只有右孩子；遇到第一个孩子不双全的节点后，后续节点必须全是叶节点。
// 时间复杂度：O(N)。空间复杂度：O(W)，W 为最大宽度。
func IsCBT1(head *CBTNode) bool {
	if head == nil {
		return true
	}
	queue := []*CBTNode{head}
	mustBeLeaf := false
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		left, right := cur.Left, cur.Right
		if (mustBeLeaf && (left != nil || right != nil)) || (left == nil && right != nil) {
			return false
		}
		if left != nil {
			queue = append(queue, left)
		}
		if right != nil {
			queue = append(queue, right)
		}
		if left == nil || right == nil {
			mustBeLeaf = true
		}
	}
	return true
}

type cbtInfo struct {
	isFull bool
	isCBT  bool
	height int
}

// IsCBT2 用递归信息判断完全二叉树。
// 每棵子树返回是否为满二叉树、是否为完全二叉树和高度，再枚举可以组成完全二叉树的四种左右形态。
// 时间复杂度：O(N)。空间复杂度：O(H)。
func IsCBT2(head *CBTNode) bool { return cbtProcess(head).isCBT }

func cbtProcess(x *CBTNode) cbtInfo {
	if x == nil {
		return cbtInfo{true, true, 0}
	}
	left, right := cbtProcess(x.Left), cbtProcess(x.Right)
	height := max(left.height, right.height) + 1
	isFull := left.isFull && right.isFull && left.height == right.height
	isCBT := isFull ||
		(left.isCBT && right.isFull && left.height == right.height+1) ||
		(left.isFull && right.isFull && left.height == right.height+1) ||
		(left.isFull && right.isCBT && left.height == right.height)
	return cbtInfo{isFull, isCBT, height}
}

func cbtGenerate(level, maxLevel, maxValue int) *CBTNode {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}
	head := &CBTNode{Value: rand.Intn(maxValue)}
	head.Left = cbtGenerate(level+1, maxLevel, maxValue)
	head.Right = cbtGenerate(level+1, maxLevel, maxValue)
	return head
}

// main 随机生成二叉树，对比层序遍历和递归信息两种判定方法。
func main() {
	for i := 0; i < 10000; i++ {
		head := cbtGenerate(1, 5, 100)
		if IsCBT1(head) != IsCBT2(head) {
			fmt.Println("Oops!")
			return
		}
	}
	fmt.Println("finish!")
}
