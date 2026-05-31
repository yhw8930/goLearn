package main

import "fmt"

// 机器人走路，核心是：
// 在 1 ~ N 这些位置上，机器人从 start 出发，每次只能走一步：
// 在中间位置：可以往左或往右
// 在 1：只能走到 2
// 在 N：只能走到 N-1
// 问：刚好走 K 步，最终停在 aim 的方法数有多少？
func main() {
	fmt.Println(ways1(5, 2, 4, 6))
	fmt.Println(ways2(5, 2, 4, 6))
	fmt.Println(ways3(5, 2, 4, 6))
}

func ways1(N, start, aim, K int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}
	return process1(start, K, aim, N)
}

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

func process2(cur, rest, aim, N int, dp [][]int) int {
	if dp[cur][rest] != -1 {
		return dp[cur][rest]
	}
	ans := 0
	if rest == 0 {
		if cur == aim {
			ans = 1
		}
		ans = 0
	} else if cur == 1 {
		ans = process1(2, rest-1, aim, N)
	} else if cur == N {
		ans = process1(N-1, rest-1, aim, N)
	} else {
		ans = process1(cur-1, rest-1, aim, N) + process1(cur+1, rest-1, aim, N)

	}
	dp[cur][rest] = ans
	return ans
}

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
