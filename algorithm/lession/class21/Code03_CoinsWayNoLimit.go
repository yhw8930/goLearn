package main

import (
	"fmt"
	"math/rand"
)

// CoinsWayNoLimit 返回使用 arr 中面值正好组成 aim 的方法数。完全背包求方案数，518
// 每种面值可以使用任意张，递归枚举当前面值使用 0 张、1 张、2 张等所有可能。
//
// 时间复杂度：O(aim^N) 量级，存在大量重复子问题。
// 空间复杂度：O(N)，递归调用栈深度。
func CoinsWayNoLimit(arr []int, aim int) int {
	if len(arr) == 0 || aim < 0 {
		return 0
	}
	return coinsWayNoLimitProcess(arr, 0, aim)
}

func coinsWayNoLimitProcess(arr []int, index, rest int) int {
	if index == len(arr) {
		if rest == 0 {
			return 1
		}
		return 0
	}
	ways := 0
	for zhang := 0; zhang*arr[index] <= rest; zhang++ {
		ways += coinsWayNoLimitProcess(arr, index+1, rest-zhang*arr[index])
	}
	return ways
}

// CoinsWayNoLimitDP1 返回无限张货币组成 aim 的方法数，是枚举张数递归的动态规划版本。
// dp[index][rest] 表示 arr[index..] 组成 rest 的方法数，仍显式枚举当前面值使用张数。
//
// 时间复杂度：O(N*aim^2)。
// 空间复杂度：O(N*aim)。
func CoinsWayNoLimitDP1(arr []int, aim int) int {
	if len(arr) == 0 || aim < 0 {
		return 0
	}
	n := len(arr)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[n][0] = 1
	for index := n - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			ways := 0
			for zhang := 0; zhang*arr[index] <= rest; zhang++ {
				ways += dp[index+1][rest-zhang*arr[index]]
			}
			dp[index][rest] = ways
		}
	}
	return dp[0][aim]
}

// CoinsWayNoLimitDP2 返回无限张货币组成 aim 的方法数，是斜率优化后的动态规划版本。
// dp[index][rest] = dp[index+1][rest] + dp[index][rest-arr[index]]，省掉枚举张数。
//
// 时间复杂度：O(N*aim)。
// 空间复杂度：O(N*aim)。
func CoinsWayNoLimitDP2(arr []int, aim int) int {
	if len(arr) == 0 || aim < 0 {
		return 0
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
				dp[index][rest] += dp[index][rest-arr[index]]
			}
		}
	}
	return dp[0][aim]
}

func coinsWayNoLimitRandomArray(maxLen, maxValue int) []int {
	n := rand.Intn(maxLen)
	arr := make([]int, n)
	has := make([]bool, maxValue+1)
	for i := 0; i < n; i++ {
		for {
			arr[i] = rand.Intn(maxValue) + 1
			if !has[arr[i]] {
				break
			}
		}
		has[arr[i]] = true
	}
	return arr
}

func coinsWayNoLimitPrintArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
