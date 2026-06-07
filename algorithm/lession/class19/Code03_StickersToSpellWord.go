package main

import "math"

// 本题测试链接：https://leetcode.com/problems/stickers-to-spell-word

// MinStickers1 返回拼出 target 至少需要多少张贴纸，每种贴纸可以使用无限张。
// 暴力递归枚举第一张贴纸，扣掉它能贡献的字符后继续解决剩余字符串。
//
// 时间复杂度：指数级，取决于贴纸数量和 target 长度。
// 空间复杂度：O(N)，N 为 target 长度，主要来自递归深度和剩余字符串。
func MinStickers1(stickers []string, target string) int {
	ans := stickersToSpellWordProcess1(stickers, target)
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

// stickersToSpellWordProcess1 返回拼出 target 的最少贴纸数。
// 如果某张贴纸不能让 target 变短，说明它作为第一张贴纸没有意义，直接跳过。
func stickersToSpellWordProcess1(stickers []string, target string) int {
	if len(target) == 0 {
		return 0
	}
	min := math.MaxInt
	for _, first := range stickers {
		rest := stickersToSpellWordMinus(target, first)
		if len(rest) != len(target) {
			next := stickersToSpellWordProcess1(stickers, rest)
			if next < min {
				min = next
			}
		}
	}
	if min == math.MaxInt {
		return min
	}
	return min + 1
}

// stickersToSpellWordMinus 返回 s1 扣掉 s2 中字符后剩余的有序字符串。
func stickersToSpellWordMinus(s1, s2 string) string {
	count := [26]int{}
	for _, ch := range s1 {
		count[ch-'a']++
	}
	for _, ch := range s2 {
		count[ch-'a']--
	}
	var rest []byte
	for i := 0; i < 26; i++ {
		for count[i] > 0 {
			rest = append(rest, byte('a'+i))
			count[i]--
		}
	}
	return string(rest)
}

// MinStickers2 返回拼出 target 至少需要多少张贴纸，是 MinStickers1 的词频优化版本。
// 先把每张贴纸转成字符词频表，再递归计算剩余词频；同时只尝试包含 target 首字符的贴纸。
//
// 时间复杂度：指数级，但词频表和首字符剪枝显著减少常数与无效分支。
// 空间复杂度：O(M*26+N)，M 为贴纸数量，N 为 target 长度。
func MinStickers2(stickers []string, target string) int {
	counts := stickersToSpellWordStickerCounts(stickers)
	ans := stickersToSpellWordProcess2(counts, target)
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

// stickersToSpellWordProcess2 使用贴纸词频表递归求解 target 的最少贴纸数。
func stickersToSpellWordProcess2(stickers [][26]int, target string) int {
	if len(target) == 0 {
		return 0
	}
	tcounts := stickersToSpellWordTargetCounts(target)
	first := target[0] - 'a'
	min := math.MaxInt
	for _, sticker := range stickers {
		if sticker[first] == 0 {
			continue
		}
		rest := stickersToSpellWordRest(tcounts, sticker)
		next := stickersToSpellWordProcess2(stickers, rest)
		if next < min {
			min = next
		}
	}
	if min == math.MaxInt {
		return min
	}
	return min + 1
}

// MinStickers3 返回拼出 target 至少需要多少张贴纸，是带记忆化搜索的最优版本。
// dp 记录每个剩余字符串的答案，避免重复计算同一个子问题。
//
// 时间复杂度：O(S*M*26+N)，S 为不同剩余字符串状态数，M 为贴纸数量，N 为 target 长度。
// 空间复杂度：O(M*26+S*N)，用于贴纸词频表、记忆化表和递归栈。
func MinStickers3(stickers []string, target string) int {
	counts := stickersToSpellWordStickerCounts(stickers)
	dp := map[string]int{"": 0}
	ans := stickersToSpellWordProcess3(counts, target, dp)
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

// stickersToSpellWordProcess3 在 process2 的基础上增加记忆化缓存。
func stickersToSpellWordProcess3(stickers [][26]int, target string, dp map[string]int) int {
	if ans, ok := dp[target]; ok {
		return ans
	}
	tcounts := stickersToSpellWordTargetCounts(target)
	first := target[0] - 'a'
	min := math.MaxInt
	for _, sticker := range stickers {
		if sticker[first] == 0 {
			continue
		}
		rest := stickersToSpellWordRest(tcounts, sticker)
		next := stickersToSpellWordProcess3(stickers, rest, dp)
		if next < min {
			min = next
		}
	}
	ans := min
	if min != math.MaxInt {
		ans = min + 1
	}
	dp[target] = ans
	return ans
}

func stickersToSpellWordStickerCounts(stickers []string) [][26]int {
	counts := make([][26]int, len(stickers))
	for i, sticker := range stickers {
		for _, ch := range sticker {
			counts[i][ch-'a']++
		}
	}
	return counts
}

func stickersToSpellWordTargetCounts(target string) [26]int {
	tcounts := [26]int{}
	for _, ch := range target {
		tcounts[ch-'a']++
	}
	return tcounts
}

func stickersToSpellWordRest(tcounts, sticker [26]int) string {
	var rest []byte
	for i := 0; i < 26; i++ {
		for count := tcounts[i] - sticker[i]; count > 0; count-- {
			rest = append(rest, byte('a'+i))
		}
	}
	return string(rest)
}
