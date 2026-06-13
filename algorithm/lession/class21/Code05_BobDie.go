package main

import "math"

// LivePosibility1 返回 Bob 从 (row,col) 出发走 k 步后仍在 N*M 区域内的概率。
// 每一步等概率向上下左右走，递归统计所有能留在棋盘内的路径数，再除以总路径数 4^k。
//
// 时间复杂度：O(4^K)。
// 空间复杂度：O(K)，递归调用栈深度。
func LivePosibility1(row, col, k, n, m int) float64 {
	return float64(bobDieProcess(row, col, k, n, m)) / math.Pow(4, float64(k))
}

func bobDieProcess(row, col, rest, n, m int) int64 {
	if row < 0 || row == n || col < 0 || col == m {
		return 0
	}
	if rest == 0 {
		return 1
	}
	up := bobDieProcess(row-1, col, rest-1, n, m)
	down := bobDieProcess(row+1, col, rest-1, n, m)
	left := bobDieProcess(row, col-1, rest-1, n, m)
	right := bobDieProcess(row, col+1, rest-1, n, m)
	return up + down + left + right
}

// LivePosibility2 返回 Bob 从 (row,col) 出发走 k 步后仍在 N*M 区域内的概率。
// dp[r][c][rest] 表示从 (r,c) 出发还剩 rest 步时的存活路径数，按步数从小到大填表。
//
// 时间复杂度：O(N*M*K)。
// 空间复杂度：O(N*M*K)。
func LivePosibility2(row, col, k, n, m int) float64 {
	dp := make([][][]int64, n)
	for i := range dp {
		dp[i] = make([][]int64, m)
		for j := range dp[i] {
			dp[i][j] = make([]int64, k+1)
			dp[i][j][0] = 1
		}
	}
	for rest := 1; rest <= k; rest++ {
		for r := 0; r < n; r++ {
			for c := 0; c < m; c++ {
				dp[r][c][rest] = bobDiePick(dp, n, m, r-1, c, rest-1)
				dp[r][c][rest] += bobDiePick(dp, n, m, r+1, c, rest-1)
				dp[r][c][rest] += bobDiePick(dp, n, m, r, c-1, rest-1)
				dp[r][c][rest] += bobDiePick(dp, n, m, r, c+1, rest-1)
			}
		}
	}
	return float64(dp[row][col][k]) / math.Pow(4, float64(k))
}

func bobDiePick(dp [][][]int64, n, m, r, c, rest int) int64 {
	if r < 0 || r == n || c < 0 || c == m {
		return 0
	}
	return dp[r][c][rest]
}
