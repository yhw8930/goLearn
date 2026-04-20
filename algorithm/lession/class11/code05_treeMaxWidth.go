package main

import (
	"math/rand"
	"time"
)

type Code05_Node struct {
	value int
	left  *Code05_Node
	right *Code05_Node
}

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
