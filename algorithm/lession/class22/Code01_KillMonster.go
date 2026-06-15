package main

import "math"

// KillMonsterRight 解决“砍怪兽”概率问题：
// 怪兽初始有 n 点血，英雄一共能砍 k 刀。每一刀的伤害是 0~m 中的任意整数，
// 并且每个伤害值等概率出现。只要 k 刀结束后怪兽血量小于等于 0，就认为怪兽被砍死。
// 返回怪兽被砍死的概率；如果 n、m、k 任意一个小于 1，按课程约定返回 0。
// 这个版本用暴力递归统计“砍死的伤害序列数”，再除以所有伤害序列总数 (m+1)^k。
//
// 时间复杂度：O((M+1)^K)。
// 空间复杂度：O(K)，递归调用栈深度。
func KillMonsterRight(n, m, k int) float64 {
	if n < 1 || m < 1 || k < 1 {
		return 0
	}
	all := int64(math.Pow(float64(m+1), float64(k)))
	kill := killMonsterProcess(k, m, n)
	return float64(kill) / float64(all)
}

// killMonsterProcess 返回还有 times 次攻击、怪兽剩 hp 血时，最终砍死怪兽的情况数。
func killMonsterProcess(times, m, hp int) int64 {
	if times == 0 {
		if hp <= 0 {
			return 1
		}
		return 0
	}
	if hp <= 0 {
		return int64(math.Pow(float64(m+1), float64(times)))
	}
	var ways int64
	for damage := 0; damage <= m; damage++ {
		ways += killMonsterProcess(times-1, m, hp-damage)
	}
	return ways
}

// KillMonsterDP1 返回怪兽被砍死的概率，是 KillMonsterRight 的动态规划版本。
// dp[times][hp] 表示还有 times 次攻击、怪兽剩 hp 血时砍死的情况数，枚举本次伤害转移。
//
// 时间复杂度：O(K*N*M)。
// 空间复杂度：O(K*N)。
func KillMonsterDP1(n, m, k int) float64 {
	if n < 1 || m < 1 || k < 1 {
		return 0
	}
	all := int64(math.Pow(float64(m+1), float64(k)))
	dp := make([][]int64, k+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}
	dp[0][0] = 1
	for times := 1; times <= k; times++ {
		dp[times][0] = int64(math.Pow(float64(m+1), float64(times)))
		for hp := 1; hp <= n; hp++ {
			var ways int64
			for damage := 0; damage <= m; damage++ {
				if hp-damage >= 0 {
					ways += dp[times-1][hp-damage]
				} else {
					ways += int64(math.Pow(float64(m+1), float64(times-1)))
				}
			}
			dp[times][hp] = ways
		}
	}
	return float64(dp[k][n]) / float64(all)
}

// KillMonsterDP2 返回怪兽被砍死的概率，是 KillMonsterDP1 的枚举优化版本。
// 利用前缀和关系：dp[times][hp] 由 dp[times][hp-1]、dp[times-1][hp] 和超窗口部分转移。
//
// 时间复杂度：O(K*N)。
// 空间复杂度：O(K*N)。
func KillMonsterDP2(n, m, k int) float64 {
	if n < 1 || m < 1 || k < 1 {
		return 0
	}
	all := int64(math.Pow(float64(m+1), float64(k)))
	dp := make([][]int64, k+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}
	dp[0][0] = 1
	for times := 1; times <= k; times++ {
		dp[times][0] = int64(math.Pow(float64(m+1), float64(times)))
		for hp := 1; hp <= n; hp++ {
			dp[times][hp] = dp[times][hp-1] + dp[times-1][hp]
			if hp-1-m >= 0 {
				dp[times][hp] -= dp[times-1][hp-1-m]
			} else {
				dp[times][hp] -= int64(math.Pow(float64(m+1), float64(times-1)))
			}
		}
	}
	return float64(dp[k][n]) / float64(all)
}
