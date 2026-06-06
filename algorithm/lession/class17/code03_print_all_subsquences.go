package main

// 题目：给定字符串，打印它的所有子序列；子序列可以不连续但必须保持原字符相对顺序。
// 每个字符都有两种选择：要当前字符，或者不要当前字符。
// 核心思路：递归来到 index 位置时，分别走“不加入 path”和“加入 path”两条分支。
// 如果字符串有重复字符，可能生成重复子序列，去重版本用 set 收集最终结果。
// 时间复杂度：O(2^N)。
// 空间复杂度：O(N) 递归深度，不计结果集；保存结果需要 O(2^N)。

// Subs 返回字符串 s 的所有子序列。
// 每个字符都只有两种选择：不要它，或要它。
//
// 时间复杂度：O(N*2^N)，共有 2^N 个结果，拷贝字符串也有代价。
// 空间复杂度：O(N)，不计返回结果时递归深度为 N；计结果为 O(N*2^N)。
func Subs(s string) []string {
	ans := make([]string, 0)
	processSubsequence([]rune(s), 0, "", &ans)
	return ans
}

func processSubsequence(str []rune, index int, path string, ans *[]string) {
	if index == len(str) {
		*ans = append(*ans, path)
		return
	}
	processSubsequence(str, index+1, path, ans)
	processSubsequence(str, index+1, path+string(str[index]), ans)
}

// SubsNoRepeat 返回去重后的所有子序列。
// 当原字符串有重复字符时，不同选择路径可能得到相同字符串，用 set 去重。
//
// 时间复杂度：O(N*2^N)。
// 空间复杂度：O(N*2^N)，set 和返回结果保存所有不同子序列。
func SubsNoRepeat(s string) []string {
	set := make(map[string]struct{})
	processSubsequenceNoRepeat([]rune(s), 0, "", set)
	ans := make([]string, 0, len(set))
	for cur := range set {
		ans = append(ans, cur)
	}
	return ans
}

func processSubsequenceNoRepeat(str []rune, index int, path string, set map[string]struct{}) {
	if index == len(str) {
		set[path] = struct{}{}
		return
	}
	processSubsequenceNoRepeat(str, index+1, path, set)
	processSubsequenceNoRepeat(str, index+1, path+string(str[index]), set)
}
