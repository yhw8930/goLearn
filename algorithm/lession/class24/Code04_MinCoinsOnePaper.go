package main

import (
	"fmt"
	"math"
	"math/rand"
)

// MinCoinsOnePaper 解决“每张纸币只能使用一次，组成目标金额的最少张数”问题：
// arr 中的每个元素代表一张纸币的面值，即使面值相同，也先把它们当作不同的纸币。
// 给定目标金额 aim，要求正好组成 aim，返回使用的最少纸币张数。
// 如果无法组成 aim，返回 math.MaxInt 作为不可达标记；aim 为 0 时答案是 0。
// 这个版本递归处理每张纸币：当前位置的纸币要么不用，要么使用一次。
//
// 时间复杂度：O(2^N)。
// 空间复杂度：O(N)，递归调用栈深度。
func MinCoinsOnePaper(arr []int, aim int) int {
	return minCoinsOnePaperProcess(arr, 0, aim)
}

func minCoinsOnePaperProcess(arr []int, index, rest int) int {
	if rest < 0 {
		return math.MaxInt
	}
	if index == len(arr) {
		if rest == 0 {
			return 0
		}
		return math.MaxInt
	}
	p1 := minCoinsOnePaperProcess(arr, index+1, rest)
	p2 := minCoinsOnePaperProcess(arr, index+1, rest-arr[index])
	if p2 != math.MaxInt {
		p2++
	}
	return minCoinsOnePaperMin(p1, p2)
}

// MinCoinsOnePaperDP1 返回每张纸币最多用一次时组成 aim 的最少张数。
// dp[index][rest] 表示 arr[index..] 自由选择，正好组成 rest 的最少张数。
//
// 时间复杂度：O(N*aim)。
// 空间复杂度：O(N*aim)。
func MinCoinsOnePaperDP1(arr []int, aim int) int {
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
			p1 := dp[index+1][rest]
			p2 := math.MaxInt
			if rest-arr[index] >= 0 {
				p2 = dp[index+1][rest-arr[index]]
			}
			if p2 != math.MaxInt {
				p2++
			}
			dp[index][rest] = minCoinsOnePaperMin(p1, p2)
		}
	}
	return dp[0][aim]
}

type minCoinsOnePaperInfo struct {
	coins  []int
	zhangs []int
}

func minCoinsOnePaperGetInfo(arr []int) minCoinsOnePaperInfo {
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
	return minCoinsOnePaperInfo{coins: coins, zhangs: zhangs}
}

// MinCoinsOnePaperDP2 返回每张纸币最多用一次时组成 aim 的最少张数。
// 先把相同面值的纸币合并成“面值 coins[i] 有 zhangs[i] 张”，再枚举每种面值使用几张。
//
// 时间复杂度：O(arr长度 + 货币种数*aim*每种货币平均张数)。
// 空间复杂度：O(货币种数*aim)。
func MinCoinsOnePaperDP2(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	info := minCoinsOnePaperGetInfo(arr)
	coins := info.coins
	zhangs := info.zhangs
	n := len(coins)
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
			for zhang := 1; zhang*coins[index] <= aim && zhang <= zhangs[index]; zhang++ {
				if rest-zhang*coins[index] >= 0 && dp[index+1][rest-zhang*coins[index]] != math.MaxInt {
					dp[index][rest] = minCoinsOnePaperMin(dp[index][rest], zhang+dp[index+1][rest-zhang*coins[index]])
				}
			}
		}
	}
	return dp[0][aim]
}

// MinCoinsOnePaperDP3 返回每张纸币最多用一次时组成 aim 的最少张数。
// 在 DP2 的合并面值基础上，对每个面值按余数分组：
// mod, mod+coin, mod+2*coin... 这些位置之间只差使用当前面值的张数。
// 用窗口最小值维护 dp[index+1][pre] + 使用张数 的最优候选，同时保证使用张数不超过 zhangs[index]。
//
// 时间复杂度：O(arr长度 + 货币种数*aim)。
// 空间复杂度：O(货币种数*aim)。
func MinCoinsOnePaperDP3(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	info := minCoinsOnePaperGetInfo(arr)
	c := info.coins
	z := info.zhangs
	n := len(c)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	for rest := 1; rest <= aim; rest++ {
		dp[n][rest] = math.MaxInt
	}
	for i := n - 1; i >= 0; i-- {
		for mod := 0; mod < minCoinsOnePaperMin(aim+1, c[i]); mod++ {
			w := make([]int, 0)
			w = append(w, mod)
			dp[i][mod] = dp[i+1][mod]
			for r := mod + c[i]; r <= aim; r += c[i] {
				for len(w) > 0 && minCoinsOnePaperWindowValue(dp[i+1][w[len(w)-1]], w[len(w)-1], r, c[i]) >= dp[i+1][r] {
					w = w[:len(w)-1]
				}
				w = append(w, r)
				overdue := r - c[i]*(z[i]+1)
				if w[0] == overdue {
					w = w[1:]
				}
				dp[i][r] = minCoinsOnePaperWindowValue(dp[i+1][w[0]], w[0], r, c[i])
			}
		}
	}
	return dp[0][aim]
}

func minCoinsOnePaperWindowValue(base, pre, cur, coin int) int {
	if base == math.MaxInt {
		return math.MaxInt
	}
	return base + minCoinsOnePaperCompensate(pre, cur, coin)
}

func minCoinsOnePaperCompensate(pre, cur, coin int) int {
	return (cur - pre) / coin
}

func minCoinsOnePaperMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCoinsOnePaperRandomArray(n, maxValue int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(maxValue) + 1
	}
	return arr
}

func minCoinsOnePaperPrintArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
