package main

import (
	"math/rand"
	"time"
)

// 题目：求二叉树所有层中节点数量最多的一层宽度。
// 宽度指同一层实际存在的节点个数，不包含空位置。
// 核心思路：层序遍历时统计每层节点数，可以用 map 记录每个节点层号。
// 也可以不用 map，用当前层最后节点和下一层最后节点两个标记判断何时层结束。
// 时间复杂度：O(N)。
// 空间复杂度：O(W) 或 O(N)，取决于是否使用层号 map。

type Code05_Node struct {
	value int
	left  *Code05_Node
	right *Code05_Node
}

// Code05_maxWidthUseMap 用层号 map 统计最大宽度。
// 每个节点入队时记录层号，遍历时统计当前层节点数。
func Code05_maxWidthUseMap(head *Code05_Node) int {
	if head == nil {
		return 0
	}
	queue := []*Code05_Node{head}
	levelMap := make(map[*Code05_Node]int)
	levelMap[head] = 1
	curLevel := 1
	curLevelNodes := 0
	maxWidth := 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		curNodeLevel := levelMap[cur]

		if cur.left != nil {
			levelMap[cur.left] = curNodeLevel + 1
			queue = append(queue, cur.left)
		}
		if cur.right != nil {
			levelMap[cur.right] = curNodeLevel + 1
			queue = append(queue, cur.right)
		}

		if curNodeLevel == curLevel {
			curLevelNodes++
		} else {
			if curLevelNodes > maxWidth {
				maxWidth = curLevelNodes
			}
			curLevel++
			curLevelNodes = 1
		}
	}
	// 处理最后一层
	if curLevelNodes > maxWidth {
		maxWidth = curLevelNodes
	}
	return maxWidth
}

// Code05_maxWidthNoMap 不使用层号 map。
// 它用当前层结束节点和下一层结束节点判断何时完成一层统计。
func Code05_maxWidthNoMap(head *Code05_Node) int {
	if head == nil {
		return 0
	}
	queue := []*Code05_Node{head}
	curEnd := head
	var nextEnd *Code05_Node
	maxWidth := 0
	curLevelNodes := 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.left != nil {
			queue = append(queue, cur.left)
			nextEnd = cur.left
		}
		if cur.right != nil {
			queue = append(queue, cur.right)
			nextEnd = cur.right
		}

		curLevelNodes++

		if cur == curEnd {
			if curLevelNodes > maxWidth {
				maxWidth = curLevelNodes
			}
			curLevelNodes = 0
			curEnd = nextEnd
		}
	}
	return maxWidth
}

func Code05_generateRandomBST(maxLevel int, maxValue int) *Code05_Node {
	rand.Seed(time.Now().UnixNano())
	return Code05_generate(1, maxLevel, maxValue)
}

func Code05_generate(level int, maxLevel int, maxValue int) *Code05_Node {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}
	head := &Code05_Node{value: rand.Intn(maxValue)}
	head.left = Code05_generate(level+1, maxLevel, maxValue)
	head.right = Code05_generate(level+1, maxLevel, maxValue)
	return head
}

func main() {
	maxLevel := 10
	maxValue := 100
	testTimes := 100000

	for i := 0; i < testTimes; i++ {
		head := Code05_generateRandomBST(maxLevel, maxValue)
		if Code05_maxWidthUseMap(head) != Code05_maxWidthNoMap(head) {
			println("Oops!")
		}
	}
	println("finish!")
}
