package main

// SplitNumberWays 解决“整数裂开”问题：
// 给定一个正数 n，把它拆成若干个正数相加，要求拆出来的数从左到右不能下降。
// 例如 4 可以拆成 1+1+1+1、1+1+2、1+3、2+2、4；而 3+1 不算新方法，
// 因为它和 1+3 只是顺序不同，并且不满足不下降要求。返回合法拆法数量。
// 递归参数 pre 表示下一段至少要拆出多大的数，rest 表示还剩多少需要继续拆。
//
// 时间复杂度：O(n^n) 量级，存在大量重复子问题。
// 空间复杂度：O(n)，递归调用栈深度。
func SplitNumberWays(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return splitNumberProcess(1, n)
}

func splitNumberProcess(pre, rest int) int {
	if rest == 0 {
		return 1
	}
	if pre > rest {
		return 0
	}
	ways := 0
	for first := pre; first <= rest; first++ {
		ways += splitNumberProcess(first, rest-first)
	}
	return ways
}

// SplitNumberDP1 返回把正数 n 拆成不下降序列的方法数，是 SplitNumberWays 的动态规划版本。
// dp[pre][rest] 表示上一个数为 pre、还剩 rest 要拆时的方法数，显式枚举下一个 first。
//
// 时间复杂度：O(n^3)。
// 空间复杂度：O(n^2)。
func SplitNumberDP1(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for pre := 1; pre <= n; pre++ {
		dp[pre][0] = 1
		dp[pre][pre] = 1
	}
	for pre := n - 1; pre >= 1; pre-- {
		for rest := pre + 1; rest <= n; rest++ {
			ways := 0
			for first := pre; first <= rest; first++ {
				ways += dp[first][rest-first]
			}
			dp[pre][rest] = ways
		}
	}
	return dp[1][n]
}

// SplitNumberDP2 返回把正数 n 拆成不下降序列的方法数，是 SplitNumberDP1 的优化版本。
// dp[pre][rest] = dp[pre+1][rest] + dp[pre][rest-pre]，分别代表不用 pre 开头和用 pre 开头。
//
// 时间复杂度：O(n^2)。
// 空间复杂度：O(n^2)。
func SplitNumberDP2(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for pre := 1; pre <= n; pre++ {
		dp[pre][0] = 1
		dp[pre][pre] = 1
	}
	for pre := n - 1; pre >= 1; pre-- {
		for rest := pre + 1; rest <= n; rest++ {
			dp[pre][rest] = dp[pre+1][rest] + dp[pre][rest-pre]
		}
	}
	return dp[1][n]
}
