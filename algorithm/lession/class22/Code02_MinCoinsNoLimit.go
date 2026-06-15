package main

import (
	"fmt"
	"math"
	"math/rand"
)

// MinCoins 解决“无限张货币组成目标金额的最少张数”问题：
// arr 中每个数代表一种正数面值，认为每种面值的货币都有无限张。
// 给定目标金额 aim，要求正好组成 aim，并返回需要的最少货币张数。
// 如果无论如何都无法组成 aim，返回 math.MaxInt 作为不可达标记；aim 为 0 时答案是 0 张。
// 这个版本从左到右考虑面值，递归枚举当前面值使用 0 张、1 张、2 张等所有可能。
//
// 时间复杂度：O(aim^N) 量级，存在大量重复子问题。
// 空间复杂度：O(N)，递归调用栈深度。
func MinCoins(arr []int, aim int) int {
	return minCoinsNoLimitProcess(arr, 0, aim)
}

func minCoinsNoLimitProcess(arr []int, index, rest int) int {
	if index == len(arr) {
		if rest == 0 {
			return 0
		}
		return math.MaxInt
	}
	ans := math.MaxInt
	for zhang := 0; zhang*arr[index] <= rest; zhang++ {
		next := minCoinsNoLimitProcess(arr, index+1, rest-zhang*arr[index])
		if next != math.MaxInt {
			ans = minCoinsNoLimitMin(ans, zhang+next)
		}
	}
	return ans
}

// MinCoinsDP1 返回用 arr 中面值组成 aim 的最少张数，是 MinCoins 的动态规划版本。
// dp[index][rest] 表示 arr[index..] 组成 rest 的最少张数，仍显式枚举当前面值张数。
//
// 时间复杂度：O(N*aim^2)。
// 空间复杂度：O(N*aim)。
func MinCoinsDP1(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	n := len(arr)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	for rest := 1; rest <= aim; rest++ {
		dp[n][rest] = math.MaxInt
	}
	for index := n - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			ans := math.MaxInt
			for zhang := 0; zhang*arr[index] <= rest; zhang++ {
				next := dp[index+1][rest-zhang*arr[index]]
				if next != math.MaxInt {
					ans = minCoinsNoLimitMin(ans, zhang+next)
				}
			}
			dp[index][rest] = ans
		}
	}
	return dp[0][aim]
}

// MinCoinsDP2 返回用 arr 中面值组成 aim 的最少张数，是斜率优化后的动态规划版本。
// dp[index][rest] 可由“不用当前面值”和“继续使用一张当前面值”两种选择转移。
//
// 时间复杂度：O(N*aim)。
// 空间复杂度：O(N*aim)。
func MinCoinsDP2(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	n := len(arr)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	for rest := 1; rest <= aim; rest++ {
		dp[n][rest] = math.MaxInt
	}
	for index := n - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			dp[index][rest] = dp[index+1][rest]
			if rest-arr[index] >= 0 && dp[index][rest-arr[index]] != math.MaxInt {
				dp[index][rest] = minCoinsNoLimitMin(dp[index][rest], dp[index][rest-arr[index]]+1)
			}
		}
	}
	return dp[0][aim]
}

func minCoinsNoLimitMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCoinsNoLimitRandomArray(maxLen, maxValue int) []int {
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

func minCoinsNoLimitPrintArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
