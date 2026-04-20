package main

import "fmt"

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
