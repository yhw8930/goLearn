package main

import (
	"fmt"
	"math/rand"
)

// CoinWays 返回使用 arr 中纸币正好组成 aim 的方法数。
// 每张纸币都被认为是不同的，即使面值相同也按不同纸币处理；递归枚举每张纸币要或不要。0/1背包求方案数，494、416
//
// 时间复杂度：O(2^N)。
// 空间复杂度：O(N)，递归调用栈深度。
func CoinWays(arr []int, aim int) int {
	return coinsWayEveryPaperDifferentProcess(arr, 0, aim)
}

func coinsWayEveryPaperDifferentProcess(arr []int, index, rest int) int {
	if rest < 0 {
		return 0
	}
	if index == len(arr) {
		if rest == 0 {
			return 1
		}
		return 0
	}
	return coinsWayEveryPaperDifferentProcess(arr, index+1, rest) +
		coinsWayEveryPaperDifferentProcess(arr, index+1, rest-arr[index])
}

// CoinsWayEveryPaperDifferentDP 返回使用 arr 中纸币正好组成 aim 的方法数。
// dp[index][rest] 表示 arr[index..] 组成 rest 的方法数，由“不用当前纸币”和“用当前纸币”转移。
//
// 时间复杂度：O(N*aim)。
// 空间复杂度：O(N*aim)。
func CoinsWayEveryPaperDifferentDP(arr []int, aim int) int {
	if aim == 0 {
		return 1
	}
	n := len(arr)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[n][0] = 1
	for index := n - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			dp[index][rest] = dp[index+1][rest]
			if rest-arr[index] >= 0 {
				dp[index][rest] += dp[index+1][rest-arr[index]]
			}
		}
	}
	return dp[0][aim]
}

func coinsWayEveryPaperDifferentRandomArray(maxLen, maxValue int) []int {
	n := rand.Intn(maxLen)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(maxValue) + 1
	}
	return arr
}

func coinsWayEveryPaperDifferentPrintArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
