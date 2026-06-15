package main

import (
	"fmt"
	"math/rand"
)

// SplitSumClosedRight 解决“数组分成两组，累加和尽量接近”问题：
// 给定一个非负数组 arr，要把所有数分成两个集合。希望两个集合的累加和尽量接近，
// 返回较小集合在最优划分下的累加和。因为两个集合总和固定，这等价于从 arr 中挑一些数，
// 让挑出的累加和不超过总和的一半，并且尽量接近这一半。
// 这个版本用暴力递归尝试每个数选或不选。
//
// 时间复杂度：O(2^N)。
// 空间复杂度：O(N)，递归调用栈深度。
func SplitSumClosedRight(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return splitSumClosedProcess(arr, 0, sum/2)
}

// splitSumClosedProcess 返回 arr[i..] 可自由选择，累加和不超过 rest 时最接近 rest 的值。
func splitSumClosedProcess(arr []int, i, rest int) int {
	if i == len(arr) {
		return 0
	}
	p1 := splitSumClosedProcess(arr, i+1, rest)
	p2 := 0
	if arr[i] <= rest {
		p2 = arr[i] + splitSumClosedProcess(arr, i+1, rest-arr[i])
	}
	return splitSumClosedMax(p1, p2)
}

// SplitSumClosedDP 返回把数组分成两部分时，较小部分最接近总和一半的累加和。
// dp[i][rest] 表示 arr[i..] 中累加和不超过 rest 的最优值，是递归尝试的动态规划版本。
//
// 时间复杂度：O(N*sum)。
// 空间复杂度：O(N*sum)。
func SplitSumClosedDP(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	sum /= 2
	n := len(arr)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, sum+1)
	}
	for i := n - 1; i >= 0; i-- {
		for rest := 0; rest <= sum; rest++ {
			p1 := dp[i+1][rest]
			p2 := 0
			if arr[i] <= rest {
				p2 = arr[i] + dp[i+1][rest-arr[i]]
			}
			dp[i][rest] = splitSumClosedMax(p1, p2)
		}
	}
	return dp[0][sum]
}

func splitSumClosedMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func splitSumClosedRandomArray(length, value int) []int {
	arr := make([]int, length)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(value)
	}
	return arr
}

func splitSumClosedPrintArray(arr []int) {
	for _, num := range arr {
		fmt.Print(num, " ")
	}
	fmt.Println()
}
