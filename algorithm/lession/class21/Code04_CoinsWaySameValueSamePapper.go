package main

import (
	"fmt"
	"math/rand"
)

type coinsWaySameValueSamePapperInfo struct {
	coins  []int
	zhangs []int
}

func coinsWaySameValueSamePapperGetInfo(arr []int) coinsWaySameValueSamePapperInfo {
	counts := make(map[int]int)
	for _, value := range arr {
		counts[value]++
	}
	coins := make([]int, 0, len(counts))
	zhangs := make([]int, 0, len(counts))
	for value, count := range counts {
		coins = append(coins, value)
		zhangs = append(zhangs, count)
	}
	return coinsWaySameValueSamePapperInfo{coins: coins, zhangs: zhangs}
}

// CoinsWaySameValueSamePapper 返回使用 arr 中纸币正好组成 aim 的方法数。
// 相同面值的纸币认为没有区别，先统计每种面值的张数，再递归枚举每种面值使用几张。多重背包求方案数，2585
//
// 时间复杂度：O(aim^K) 量级，K 是不同面值数量。
// 空间复杂度：O(K)，递归调用栈深度。
func CoinsWaySameValueSamePapper(arr []int, aim int) int {
	if len(arr) == 0 || aim < 0 {
		return 0
	}
	info := coinsWaySameValueSamePapperGetInfo(arr)
	return coinsWaySameValueSamePapperProcess(info.coins, info.zhangs, 0, aim)
}

func coinsWaySameValueSamePapperProcess(coins, zhangs []int, index, rest int) int {
	if index == len(coins) {
		if rest == 0 {
			return 1
		}
		return 0
	}
	ways := 0
	for zhang := 0; zhang*coins[index] <= rest && zhang <= zhangs[index]; zhang++ {
		ways += coinsWaySameValueSamePapperProcess(coins, zhangs, index+1, rest-zhang*coins[index])
	}
	return ways
}

// CoinsWaySameValueSamePapperDP1 返回有限张同值无差别纸币组成 aim 的方法数。
// dp[index][rest] 表示 coins[index..] 组成 rest 的方法数，显式枚举当前面值使用张数。
//
// 时间复杂度：O(K*aim*每种最大张数)。
// 空间复杂度：O(K*aim)。
func CoinsWaySameValueSamePapperDP1(arr []int, aim int) int {
	if len(arr) == 0 || aim < 0 {
		return 0
	}
	info := coinsWaySameValueSamePapperGetInfo(arr)
	coins := info.coins
	zhangs := info.zhangs
	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[n][0] = 1
	for index := n - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			ways := 0
			for zhang := 0; zhang*coins[index] <= rest && zhang <= zhangs[index]; zhang++ {
				ways += dp[index+1][rest-zhang*coins[index]]
			}
			dp[index][rest] = ways
		}
	}
	return dp[0][aim]
}

// CoinsWaySameValueSamePapperDP2 返回有限张同值无差别纸币组成 aim 的方法数。
// 在 DP1 基础上用窗口关系优化：加上同层 rest-coin 的值，再扣掉超过张数限制的部分。
//
// 时间复杂度：O(K*aim)。
// 空间复杂度：O(K*aim)。
func CoinsWaySameValueSamePapperDP2(arr []int, aim int) int {
	if len(arr) == 0 || aim < 0 {
		return 0
	}
	info := coinsWaySameValueSamePapperGetInfo(arr)
	coins := info.coins
	zhangs := info.zhangs
	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[n][0] = 1
	for index := n - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			dp[index][rest] = dp[index+1][rest]
			if rest-coins[index] >= 0 {
				dp[index][rest] += dp[index][rest-coins[index]]
			}
			overRest := rest - coins[index]*(zhangs[index]+1)
			if overRest >= 0 {
				dp[index][rest] -= dp[index+1][overRest]
			}
		}
	}
	return dp[0][aim]
}

func coinsWaySameValueSamePapperRandomArray(maxLen, maxValue int) []int {
	n := rand.Intn(maxLen)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(maxValue) + 1
	}
	return arr
}

func coinsWaySameValueSamePapperPrintArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
