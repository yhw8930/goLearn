package main

import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

// 题目：给定字符串数组，把所有字符串拼接起来，要求最终结果字典序最小。
// 暴力方法可以枚举所有排列，但字符串数量稍大时阶乘复杂度不可接受。
// 核心思路：定义排序规则，如果 a+b 的字典序小于 b+a，就让 a 排在 b 前面。
// 这个比较规则能保证局部相邻顺序最优，从而得到整体最小拼接结果。
// 时间复杂度：暴力为 O(N!)；贪心排序为 O(NlogN * K)，K 为拼接比较长度。
// 空间复杂度：排序和拼接结果需要 O(N*K)。

// ---------------------- 暴力解法（全排列）----------------------
// lowestString1 是暴力枚举方法。
// 它生成所有字符串排列并取字典序最小结果，适合验证贪心。
// 时间复杂度：O(N!*N*K)，K 为字符串平均长度。
// 空间复杂度：O(N!*N*K)，保存所有排列结果。
func lowestString1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ans := process(strs)
	if len(ans) == 0 {
		return ""
	}
	// 取最小字典序（Go map 无序，所以遍历找最小）
	minStr := ""
	first := true
	for s := range ans {
		if first || s < minStr {
			minStr = s
			first = false
		}
	}
	return minStr
}

// 递归生成全排列
func process(strs []string) map[string]bool {
	ans := make(map[string]bool)
	if len(strs) == 0 {
		ans[""] = true
		return ans
	}

	for i := 0; i < len(strs); i++ {
		first := strs[i]
		nexts := removeIndexString(strs, i)
		next := process(nexts)
		for cur := range next {
			ans[first+cur] = true
		}
	}
	return ans
}

// 移除指定下标的字符串
func removeIndexString(arr []string, index int) []string {
	n := len(arr)
	ans := make([]string, n-1)
	ansIdx := 0
	for i := 0; i < n; i++ {
		if i != index {
			ans[ansIdx] = arr[i]
			ansIdx++
		}
	}
	return ans
}

// ---------------------- 贪心解法（最优）----------------------
// 自定义排序器：a+b < b+a 则 a 排前面
type strSlice []string

func (s strSlice) Len() int      { return len(s) }
func (s strSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s strSlice) Less(i, j int) bool {
	return s[i]+s[j] < s[j]+s[i]
}

// lowestString2 是贪心排序方法。
// 它按照 a+b 与 b+a 的比较结果排序，排序后直接拼接。
// 时间复杂度：O(NlogN*K)，比较 a+b 和 b+a 的代价为 O(K)。
// 空间复杂度：O(N*K)，主要为排序和拼接结果。
func lowestString2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Sort(strSlice(strs))
	return strings.Join(strs, "")
}

// ---------------------- 测试工具方法 ----------------------
func generateRandomString(strLen int) string {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(strLen) + 1
	ans := make([]rune, length)
	for i := 0; i < length; i++ {
		value := rand.Intn(5)
		if rand.Float64() <= 0.5 {
			ans[i] = rune(65 + value) // 大写 A-E
		} else {
			ans[i] = rune(97 + value) // 小写 a-e
		}
	}
	return string(ans)
}

func generateRandomStringArray(arrLen, strLen int) []string {
	length := rand.Intn(arrLen) + 1
	ans := make([]string, length)
	for i := 0; i < length; i++ {
		ans[i] = generateRandomString(strLen)
	}
	return ans
}

func copyStringArray(arr []string) []string {
	ans := make([]string, len(arr))
	copy(ans, arr)
	return ans
}

// ---------------------- 主测试 ----------------------
func main() {
	arrLen := 6
	strLen := 5
	testTimes := 1000
	rand.Seed(time.Now().UnixNano())

	println("test begin")
	for i := 0; i < testTimes; i++ {
		arr1 := generateRandomStringArray(arrLen, strLen)
		arr2 := copyStringArray(arr1)

		res1 := lowestString1(arr1)
		res2 := lowestString2(arr2)

		if res1 != res2 {
			println("Oops!")
			return
		}
	}
	println("finish!")
}
