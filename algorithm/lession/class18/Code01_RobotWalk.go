package main

import "fmt"

// 题目：一排有 1..N 个位置，机器人从 start 出发，每次必须走一步，走 K 步后停在 aim，求方法数。
// 机器人在 1 位置只能走到 2，在 N 位置只能走到 N-1，在中间位置可以左右二选一。
// 核心思路：递归状态是“当前在 cur，还剩 rest 步”，rest 为 0 时检查是否到达 aim。
// 同一 cur/rest 会被反复计算，所以可以改成记忆化搜索，再进一步整理成严格动态规划表。
// 时间复杂度：暴力 O(2^K)，记忆化和动态规划 O(N*K)。
// 空间复杂度：暴力 O(K)，记忆化和动态规划 O(N*K)。
func main() {
	fmt.Println(ways1(5, 2, 4, 6))
	fmt.Println(ways2(5, 2, 4, 6))
	fmt.Println(ways3(5, 2, 4, 6))
}

// ways1 暴力递归求机器人从 start 出发，刚好走 K 步到 aim 的方法数。
// 每一步根据当前位置分成只能向内走或左右两种选择。
//
// 时间复杂度：O(2^K)。
// 空间复杂度：O(K)，递归调用栈深度。
func ways1(N, start, aim, K int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}
	return process1(start, K, aim, N)
}

// process1 表示机器人当前在 cur，还剩 rest 步，最终到 aim 的方法数。
// 位置范围固定为 1~N，边界位置只能向内侧移动。
func process1(cur, rest, aim, N int) int {
	if rest == 0 {
		if cur == aim {
			return 1
		}
		return 0
	}
	if cur == 1 {
		return process1(2, rest-1, aim, N)
	}
	if cur == N {
		return process1(N-1, rest-1, aim, N)
	}
	return process1(cur-1, rest-1, aim, N) + process1(cur+1, rest-1, aim, N)
}

// ways2 在暴力递归基础上增加缓存，避免重复计算相同的 cur/rest 状态。
// 时间复杂度：O(N*K)。
// 空间复杂度：O(N*K)，缓存表。
func ways2(N, start, aim, K int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}
	dp := make([][]int, N+1)
	for i, _ := range dp {
		dp[i] = make([]int, K+1)
		for j, _ := range dp[i] {
			dp[i][j] = -1
		}
	}
	return process2(start, K, aim, N, dp)
}

// process2 表示带缓存的递归状态，dp[cur][rest] 记录已计算过的答案。
func process2(cur, rest, aim, N int, dp [][]int) int {
	if dp[cur][rest] != -1 {
		return dp[cur][rest]
	}
	ans := 0
	if rest == 0 {
		if cur == aim {
			ans = 1
		} else {
			ans = 0
		}
	} else if cur == 1 {
		ans = process2(2, rest-1, aim, N, dp)
	} else if cur == N {
		ans = process2(N-1, rest-1, aim, N, dp)
	} else {
		ans = process2(cur-1, rest-1, aim, N, dp) + process2(cur+1, rest-1, aim, N, dp)

	}
	dp[cur][rest] = ans
	return ans
}

// ways3 使用严格位置依赖的动态规划。
// dp[cur][rest] 表示从 cur 出发还剩 rest 步，最终到 aim 的方法数。
//
// 时间复杂度：O(N*K)。
// 空间复杂度：O(N*K)。
func ways3(N, start, aim, K int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}
	dp := make([][]int, N+1)
	for i, _ := range dp {
		dp[i] = make([]int, K+1)
	}
	dp[aim][0] = 1
	for rest := 1; rest <= K; rest++ {
		dp[1][rest] = dp[2][rest-1]
		for cur := 2; cur < N; cur++ {
			dp[cur][rest] = dp[cur-1][rest-1] + dp[cur+1][rest-1]
		}
		dp[N][rest] = dp[N-1][rest-1]
	}
	return dp[start][K]
}
