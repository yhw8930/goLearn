package main

import "fmt"

// 题目：数字字符串只含 0..9，按 1->A、2->B、...、26->Z 转成字母，求有多少种合法转法。
// 字符 0 不能单独转化，只能作为 10 或 20 的一部分出现。
// 核心思路：递归来到 i 位置时，如果 str[i] 是 0 直接返回 0。
// 否则可以让 str[i] 单独转，也可以在 i+1 存在且两位数不超过 26 时把 str[i..i+1] 一起转。
// 时间复杂度：暴力 O(2^N)，动态规划 O(N)。
// 空间复杂度：暴力 O(N)，动态规划 O(N)。

// main 演示数字字符串转字母字符串的方案数问题：
// 字符串只包含数字字符 '0'~'9'，其中 1->A、2->B、...、26->Z。
// 要求返回整个数字字符串有多少种合法转化方案；单独的 '0' 不能转化。
func main() {
	fmt.Println(number("7210231231232031203123"))
	fmt.Println(convertToLetterStringDP("7210231231232031203123"))
}

// number 返回数字字符串 str 有多少种转成字母字符串的方案。
// 从左到右递归尝试：当前位置可以单独转，也可以和下一位组成 10~26 一起转。
//
// 时间复杂度：O(2^N)。
// 空间复杂度：O(N)，递归调用栈深度。
func number(str string) int {
	if len(str) == 0 {
		return 0
	}
	return convertToLetterStringProcess([]byte(str), 0)
}

// convertToLetterStringProcess 表示 str[0..i-1] 已经完成转化，返回 str[i...] 的转化方案数。
// 如果当前位置是 '0'，说明之前的决定无法让这个 0 合法归属，返回 0。
//
// 时间复杂度：O(2^(N-i))。
// 空间复杂度：O(N-i)。
func convertToLetterStringProcess(str []byte, i int) int {
	if i == len(str) {
		return 1
	}
	if str[i] == '0' {
		return 0
	}
	ways := convertToLetterStringProcess(str, i+1)
	if i+1 < len(str) && (str[i]-'0')*10+str[i+1]-'0' < 27 {
		ways += convertToLetterStringProcess(str, i+2)
	}
	return ways
}

// convertToLetterStringDP 返回数字字符串 s 有多少种转成字母字符串的方案，是 number 的动态规划版本。
// dp[i] 表示 s[i...] 的转化方案数，从右往左填表。
//
// 时间复杂度：O(N)。
// 空间复杂度：O(N)。
func convertToLetterStringDP(s string) int {
	if len(s) == 0 {
		return 0
	}
	str := []byte(s)
	n := len(str)
	dp := make([]int, n+1)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		if str[i] != '0' {
			ways := dp[i+1]
			if i+1 < n && (str[i]-'0')*10+str[i+1]-'0' < 27 {
				ways += dp[i+2]
			}
			dp[i] = ways
		}
	}
	return dp[0]
}
