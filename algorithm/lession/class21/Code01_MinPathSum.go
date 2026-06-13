package main

import (
	"fmt"
	"math/rand"
)

// MinPathSum1 返回矩阵从左上角走到右下角的最小路径和，每次只能向右或向下。
// dp[i][j] 表示走到 (i,j) 的最小路径和，由上方或左方的最优值转移而来。
//
// 时间复杂度：O(row*col)。
// 空间复杂度：O(row*col)。
func MinPathSum1(m [][]int) int {
	if len(m) == 0 || len(m[0]) == 0 {
		return 0
	}
	row := len(m)
	col := len(m[0])
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}
	dp[0][0] = m[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + m[i][0]
	}
	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + m[0][j]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + m[i][j]
		}
	}
	return dp[row-1][col-1]
}

// MinPathSum2 返回矩阵从左上角走到右下角的最小路径和，是 MinPathSum1 的空间压缩版本。
// 一维 dp[j] 表示当前行走到第 j 列的最小路径和，更新时同时利用左侧和上一行的值。
//
// 时间复杂度：O(row*col)。
// 空间复杂度：O(col)。
func MinPathSum2(m [][]int) int {
	if len(m) == 0 || len(m[0]) == 0 {
		return 0
	}
	row := len(m)
	col := len(m[0])
	dp := make([]int, col)
	dp[0] = m[0][0]
	for j := 1; j < col; j++ {
		dp[j] = dp[j-1] + m[0][j]
	}
	for i := 1; i < row; i++ {
		dp[0] += m[i][0]
		for j := 1; j < col; j++ {
			dp[j] = min(dp[j-1], dp[j]) + m[i][j]
		}
	}
	return dp[col-1]
}

func minPathSumGenerateRandomMatrix(rowSize, colSize int) [][]int {
	if rowSize < 0 || colSize < 0 {
		return nil
	}
	result := make([][]int, rowSize)
	for i := range result {
		result[i] = make([]int, colSize)
		for j := range result[i] {
			result[i][j] = rand.Intn(100)
		}
	}
	return result
}

func minPathSumPrintMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
