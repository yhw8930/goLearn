package main

import "fmt"

// 题目：有 N 件物品，每件有重量 w[i] 和价值 v[i]，背包容量为 bag，求不超重时最大价值。
// 每件物品只能选择放入或不放入，不能切割，也不能重复放入。
// 核心思路：递归状态是“来到 index 物品，背包剩余容量 rest”，选择跳过当前物品或放入当前物品。
// 如果放入导致 rest 小于 0，则该方案无效；递归依赖 index/rest 两个变量，可以改成二维动态规划表。
// 时间复杂度：暴力 O(2^N)，动态规划 O(N*bag)。
// 空间复杂度：暴力 O(N)，动态规划 O(N*bag)。

// main 演示 0/1 背包问题：
// 有一个背包，最大承重为 bag。给定 N 件物品，第 i 件物品重量为 w[i]，价值为 v[i]。
// 每件物品只能选择放入或不放入，不能切割，也不能重复放入。
// 要求在总重量不超过 bag 的前提下，返回能够获得的最大总价值。
func main() {
	weights := []int{3, 2, 4, 7, 3, 1, 7}
	values := []int{5, 6, 3, 19, 12, 4, 2}
	bag := 15
	fmt.Println(maxValue(weights, values, bag))
	fmt.Println(knapsackDP(weights, values, bag))
}

// maxValue 求 0/1 背包在不超过 bag 容量时能获得的最大价值。
// 每件物品只能选择要或不要，递归从左到右尝试所有可能。
//
// 时间复杂度：O(2^N)。
// 空间复杂度：O(N)，递归调用栈深度。
func maxValue(w, v []int, bag int) int {
	if len(w) == 0 || len(v) == 0 || len(w) != len(v) || bag < 0 {
		return 0
	}
	return knapsackProcess(w, v, 0, bag)
}

// knapsackProcess 表示从 index 位置开始自由选择，剩余容量为 rest 时的最大价值。
// rest 小于 0 代表之前的选择已经超重，用 -1 表示无效方案。
//
// 时间复杂度：O(2^(N-index))。
// 空间复杂度：O(N-index)。
func knapsackProcess(w, v []int, index, rest int) int {
	if rest < 0 {
		return -1
	}
	if index == len(w) {
		return 0
	}
	p1 := knapsackProcess(w, v, index+1, rest)
	p2 := 0
	next := knapsackProcess(w, v, index+1, rest-w[index])
	if next != -1 {
		p2 = v[index] + next
	}
	return max(p1, p2)
}

// knapsackDP 求 0/1 背包最大价值，是 maxValue 递归尝试的动态规划版本。
// dp[index][rest] 表示从 index 位置开始选，剩余容量为 rest 时的最大价值。
//
// 时间复杂度：O(N*bag)。
// 空间复杂度：O(N*bag)。
func knapsackDP(w, v []int, bag int) int {
	if len(w) == 0 || len(v) == 0 || len(w) != len(v) || bag < 0 {
		return 0
	}
	n := len(w)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, bag+1)
	}
	for index := n - 1; index >= 0; index-- {
		for rest := 0; rest <= bag; rest++ {
			p1 := dp[index+1][rest]
			p2 := 0
			if rest-w[index] >= 0 {
				p2 = v[index] + dp[index+1][rest-w[index]]
			}
			dp[index][rest] = max(p1, p2)
		}
	}
	return dp[0][bag]
}
