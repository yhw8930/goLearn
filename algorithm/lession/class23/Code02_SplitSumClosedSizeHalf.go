package main

import (
	"fmt"
	"math"
	"math/rand"
)

// SplitSumClosedSizeHalfRight 解决“数组分成两组，个数接近且累加和接近”问题：
// 给定一个非负数组 arr，要把所有数分成两个集合。除了要求两个集合的累加和尽量接近，
// 还要求两个集合的元素个数也尽量接近：如果数组长度为偶数，两边必须各有 N/2 个数；
// 如果数组长度为奇数，两边个数必须分别是 N/2 和 N/2+1。
// 返回满足个数要求时，较小累加和在最优划分下能达到的最大值。
// 这个版本递归枚举每个位置的数选或不选，并用 picks 限制必须挑够的数量。
//
// 时间复杂度：O(2^N)。
// 空间复杂度：O(N)，递归调用栈深度。
func SplitSumClosedSizeHalfRight(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	if len(arr)&1 == 0 {
		return splitSumClosedSizeHalfProcess(arr, 0, len(arr)/2, sum/2)
	}
	return splitSumClosedSizeHalfMax(
		splitSumClosedSizeHalfProcess(arr, 0, len(arr)/2, sum/2),
		splitSumClosedSizeHalfProcess(arr, 0, len(arr)/2+1, sum/2),
	)
}

// splitSumClosedSizeHalfProcess 返回 arr[i..] 中必须挑 picks 个数、累加和不超过 rest 的最优值。
// 返回 -1 表示无法完成指定挑选数量。
func splitSumClosedSizeHalfProcess(arr []int, i, picks, rest int) int {
	if i == len(arr) {
		if picks == 0 {
			return 0
		}
		return -1
	}
	p1 := splitSumClosedSizeHalfProcess(arr, i+1, picks, rest)
	p2 := -1
	next := -1
	if picks-1 >= 0 && arr[i] <= rest {
		next = splitSumClosedSizeHalfProcess(arr, i+1, picks-1, rest-arr[i])
	}
	if next != -1 {
		p2 = arr[i] + next
	}
	return splitSumClosedSizeHalfMax(p1, p2)
}

// SplitSumClosedSizeHalfDP 返回个数尽量相等时较小累加和最接近总和一半的值。
// dp[i][picks][rest] 表示 arr[i..] 必须挑 picks 个数且不超过 rest 的最优值。
//
// 时间复杂度：O(N*M*sum)，M=(N+1)/2。
// 空间复杂度：O(N*M*sum)。
func SplitSumClosedSizeHalfDP(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	sum /= 2
	n := len(arr)
	m := (n + 1) / 2
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, sum+1)
			for rest := range dp[i][j] {
				dp[i][j][rest] = -1
			}
		}
	}
	for rest := 0; rest <= sum; rest++ {
		dp[n][0][rest] = 0
	}
	for i := n - 1; i >= 0; i-- {
		for picks := 0; picks <= m; picks++ {
			for rest := 0; rest <= sum; rest++ {
				p1 := dp[i+1][picks][rest]
				p2 := -1
				next := -1
				if picks-1 >= 0 && arr[i] <= rest {
					next = dp[i+1][picks-1][rest-arr[i]]
				}
				if next != -1 {
					p2 = arr[i] + next
				}
				dp[i][picks][rest] = splitSumClosedSizeHalfMax(p1, p2)
			}
		}
	}
	if len(arr)&1 == 0 {
		return dp[0][len(arr)/2][sum]
	}
	return splitSumClosedSizeHalfMax(dp[0][len(arr)/2][sum], dp[0][len(arr)/2+1][sum])
}

// SplitSumClosedSizeHalfDP2 返回个数尽量相等时较小累加和最接近总和一半的值。
// 该版本按前 i 个数建表，dp[i][j][k] 表示前 i+1 个数中挑 j 个且不超过 k 的最优值。
//
// 时间复杂度：O(N*M*sum)。
// 空间复杂度：O(N*M*sum)。
func SplitSumClosedSizeHalfDP2(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	sum >>= 1
	n := len(arr)
	m := (n + 1) >> 1
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, sum+1)
			for k := range dp[i][j] {
				dp[i][j][k] = math.MinInt
			}
		}
	}
	for i := 0; i < n; i++ {
		for k := 0; k <= sum; k++ {
			dp[i][0][k] = 0
		}
	}
	for k := 0; k <= sum; k++ {
		if arr[0] <= k {
			dp[0][1][k] = arr[0]
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= splitSumClosedSizeHalfMin(i+1, m); j++ {
			for k := 0; k <= sum; k++ {
				dp[i][j][k] = dp[i-1][j][k]
				if k-arr[i] >= 0 && dp[i-1][j-1][k-arr[i]] != math.MinInt {
					dp[i][j][k] = splitSumClosedSizeHalfMax(dp[i][j][k], dp[i-1][j-1][k-arr[i]]+arr[i])
				}
			}
		}
	}
	return splitSumClosedSizeHalfMax(dp[n-1][m][sum], dp[n-1][n-m][sum])
}

func splitSumClosedSizeHalfMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func splitSumClosedSizeHalfMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func splitSumClosedSizeHalfRandomArray(length, value int) []int {
	arr := make([]int, length)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(value)
	}
	return arr
}

func splitSumClosedSizeHalfPrintArray(arr []int) {
	for _, num := range arr {
		fmt.Print(num, " ")
	}
	fmt.Println()
}
