package main

// 测试链接：https://leetcode.com/problems/longest-palindromic-subsequence/

// Lpsl1 返回字符串 s 的最长回文子序列长度。
// 递归含义：palindromeSubsequenceF(str, l, r) 表示 str[l..r] 上的答案，
// 枚举左右端点是否参与最终回文子序列。
//
// 时间复杂度：O(4^N)，存在大量重复子问题。
// 空间复杂度：O(N)，递归调用栈深度。
func Lpsl1(s string) int {
	if len(s) == 0 {
		return 0
	}
	str := []byte(s)
	return palindromeSubsequenceF(str, 0, len(str)-1)
}

func palindromeSubsequenceF(str []byte, l, r int) int {
	if l == r {
		return 1
	}
	if l == r-1 {
		if str[l] == str[r] {
			return 2
		}
		return 1
	}
	p1 := palindromeSubsequenceF(str, l+1, r-1)
	p2 := palindromeSubsequenceF(str, l, r-1)
	p3 := palindromeSubsequenceF(str, l+1, r)
	p4 := 0
	if str[l] == str[r] {
		p4 = 2 + palindromeSubsequenceF(str, l+1, r-1)
	}
	return max(max(p1, p2), max(p3, p4))
}

// Lpsl2 返回字符串 s 的最长回文子序列长度，是 Lpsl1 的动态规划版本。
// dp[l][r] 表示 str[l..r] 的最长回文子序列长度，按区间长度从小到大填表。
//
// 时间复杂度：O(N^2)。
// 空间复杂度：O(N^2)。
func Lpsl2(s string) int {
	if len(s) == 0 {
		return 0
	}
	str := []byte(s)
	n := len(str)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[n-1][n-1] = 1
	for i := 0; i < n-1; i++ {
		dp[i][i] = 1
		if str[i] == str[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i+1] = 1
		}
	}
	for l := n - 3; l >= 0; l-- {
		for r := l + 2; r < n; r++ {
			dp[l][r] = max(dp[l][r-1], dp[l+1][r])
			if str[l] == str[r] {
				dp[l][r] = max(dp[l][r], 2+dp[l+1][r-1])
			}
		}
	}
	return dp[0][n-1]
}

// LongestPalindromeSubseq1 返回字符串 s 的最长回文子序列长度。
// 一个字符串的最长回文子序列，等价于它和反转串的最长公共子序列。
//
// 时间复杂度：O(N^2)。
// 空间复杂度：O(N^2)。
func LongestPalindromeSubseq1(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	str := []byte(s)
	reverse := palindromeSubsequenceReverse(str)
	return palindromeSubsequenceLCS(str, reverse)
}

func palindromeSubsequenceReverse(str []byte) []byte {
	n := len(str)
	reverse := make([]byte, n)
	for i := 0; i < len(str); i++ {
		n--
		reverse[n] = str[i]
	}
	return reverse
}

func palindromeSubsequenceLCS(str1, str2 []byte) int {
	n := len(str1)
	m := len(str2)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	if str1[0] == str2[0] {
		dp[0][0] = 1
	}
	for i := 1; i < n; i++ {
		if str1[i] == str2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for j := 1; j < m; j++ {
		if str1[0] == str2[j] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if str1[i] == str2[j] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return dp[n-1][m-1]
}

// LongestPalindromeSubseq2 返回字符串 s 的最长回文子序列长度。
// 直接在原串上做区间动态规划，dp[i][j] 由左侧、下侧和左下内层区间转移而来。
//
// 时间复杂度：O(N^2)。
// 空间复杂度：O(N^2)。
func LongestPalindromeSubseq2(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	str := []byte(s)
	n := len(str)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[n-1][n-1] = 1
	for i := 0; i < n-1; i++ {
		dp[i][i] = 1
		if str[i] == str[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i+1] = 1
		}
	}
	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			if str[i] == str[j] {
				dp[i][j] = max(dp[i][j], dp[i+1][j-1]+2)
			}
		}
	}
	return dp[0][n-1]
}
