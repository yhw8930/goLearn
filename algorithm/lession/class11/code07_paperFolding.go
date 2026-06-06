package main

import "fmt"

// 题目：一张纸连续对折 N 次后，从上到下打印每条折痕是凹还是凸。
// 第 1 次折痕一定是凹；每个折痕继续展开时，左子折痕是凹，右子折痕是凸。
// 核心思路：折痕展开顺序等价于一棵深度为 N 的二叉树中序遍历。
// 递归参数记录当前层数和当前折痕方向，先打印左凹，再打印当前，再打印右凸。
// 时间复杂度：O(2^N)，需要打印所有折痕。
// 空间复杂度：O(N)，递归深度。

// 折纸问题（中序遍历二叉树）
func Code07_printAllFolds(N int) {
	Code07_process(1, N, true)
	fmt.Println()
}

// 当前你来了一个节点，脑海中想象的！
// 这个节点在第i层，一共有N层，N固定不变的
// 这个节点如果是凹的话，down = T
// 这个节点如果是凸的话，down = F
// 函数的功能：中序打印以你想象的节点为头的整棵树！
func Code07_process(i int, N int, down bool) {
	if i > N {
		return
	}
	Code07_process(i+1, N, true)
	// 打印凹 or 凸
	if down {
		fmt.Print("凹 ")
	} else {
		fmt.Print("凸 ")
	}
	Code07_process(i+1, N, false)
}

func main() {
	N := 4
	Code07_printAllFolds(N)
}
