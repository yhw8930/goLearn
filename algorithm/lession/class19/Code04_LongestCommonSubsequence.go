package main

// 这个问题leetcode上可以直接测
// 链接：https://leetcode.com/problems/longest-common-subsequence/

// LongestCommonSubsequence1 返回 s1 和 s2 的最长公共子序列长度。
// 递归含义：只考虑 str1[0..i] 和 str2[0..j]，答案来自跳过 str1[i]、跳过 str2[j]、
// 或在两字符相等时同时保留它们这三种可能。
//
// 时间复杂度：O(3^(N+M))，存在大量重复子问题。
// 空间复杂度：O(N+M)，递归调用栈深度。
func LongestCommonSubsequence1(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	str1 := []byte(s1)
	str2 := []byte(s2)
	return longestCommonSubsequenceProcess1(str1, str2, len(str1)-1, len(str2)-1)
}

func longestCommonSubsequenceProcess1(str1, str2 []byte, i, j int) int {
	if i == 0 && j == 0 {
		if str1[i] == str2[j] {
			return 1
		}
		return 0
	}
	if i == 0 {
		if str1[i] == str2[j] {
			return 1
		}
		return longestCommonSubsequenceProcess1(str1, str2, i, j-1)
	}
	if j == 0 {
		if str1[i] == str2[j] {
			return 1
		}
		return longestCommonSubsequenceProcess1(str1, str2, i-1, j)
	}
	p1 := longestCommonSubsequenceProcess1(str1, str2, i-1, j)
	p2 := longestCommonSubsequenceProcess1(str1, str2, i, j-1)
	p3 := 0
	if str1[i] == str2[j] {
		p3 = 1 + longestCommonSubsequenceProcess1(str1, str2, i-1, j-1)
	}
	return max(p1, max(p2, p3))
}

// LongestCommonSubsequence2 返回 s1 和 s2 的最长公共子序列长度，是递归尝试的动态规划版本。
// dp[i][j] 表示 str1[0..i] 和 str2[0..j] 的最长公共子序列长度，按依赖从左上到右下填表。
//
// 时间复杂度：O(N*M)。
// 空间复杂度：O(N*M)。
func LongestCommonSubsequence2(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	str1 := []byte(s1)
	str2 := []byte(s2)
	n := len(str1)
	m := len(str2)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	if str1[0] == str2[0] {
		dp[0][0] = 1
	}
	for j := 1; j < m; j++ {
		if str1[0] == str2[j] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < n; i++ {
		if str1[i] == str2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			p1 := dp[i-1][j]
			p2 := dp[i][j-1]
			p3 := 0
			if str1[i] == str2[j] {
				p3 = 1 + dp[i-1][j-1]
			}
			dp[i][j] = max(p1, max(p2, p3))
		}
	}
	return dp[n-1][m-1]
}
