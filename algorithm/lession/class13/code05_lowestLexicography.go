package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// LowestString1 给定字符串数组，要求把所有字符串各使用一次并拼接，返回字典序最小的结果；空数组返回空串。
// 暴力枚举所有排列，把完整拼接结果放入集合，再从中选出字典序最小者，适合作为对数器。
// 时间复杂度：O(N!*L)，L 为全部字符串的总长度。空间复杂度：O(N!*L)。
func LowestString1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	all := lowestStringPermutations(strs)
	ans := ""
	first := true
	for s := range all {
		if first || s < ans {
			ans, first = s, false
		}
	}
	return ans
}

func lowestStringPermutations(strs []string) map[string]struct{} {
	ans := make(map[string]struct{})
	if len(strs) == 0 {
		ans[""] = struct{}{}
		return ans
	}
	for i, first := range strs {
		for rest := range lowestStringPermutations(lowestStringRemove(strs, i)) {
			ans[first+rest] = struct{}{}
		}
	}
	return ans
}

func lowestStringRemove(strs []string, index int) []string {
	ans := make([]string, 0, len(strs)-1)
	ans = append(ans, strs[:index]...)
	ans = append(ans, strs[index+1:]...)
	return ans
}

// LowestString2 用贪心排序得到字典序最小的拼接结果。
// 若 a+b < b+a，则 a 必须排在 b 前；按此比较器排序后直接拼接即可得到全局最优解。函数会修改 strs 的顺序。
// 时间复杂度：O(N*logN*L)，L 表示一次拼接比较的最大字符数。空间复杂度：O(L)。
func LowestString2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Slice(strs, func(i, j int) bool { return strs[i]+strs[j] < strs[j]+strs[i] })
	return strings.Join(strs, "")
}

func lowestStringRandom(strLen int) string {
	length := rand.Intn(strLen) + 1
	ans := make([]byte, length)
	for i := range ans {
		value := byte(rand.Intn(5))
		if rand.Float64() < 0.5 {
			ans[i] = 'A' + value
		} else {
			ans[i] = 'a' + value
		}
	}
	return string(ans)
}

func lowestStringRandomArray(arrLen, strLen int) []string {
	ans := make([]string, rand.Intn(arrLen)+1)
	for i := range ans {
		ans[i] = lowestStringRandom(strLen)
	}
	return ans
}

// main 随机生成字符串数组，对比全排列和贪心排序得到的最小字典序拼接结果。
func main() {
	fmt.Println("test begin")
	for i := 0; i < 10000; i++ {
		arr1 := lowestStringRandomArray(6, 5)
		arr2 := append([]string(nil), arr1...)
		if LowestString1(arr1) != LowestString2(arr2) {
			fmt.Println(arr1)
			fmt.Println("Oops!")
			return
		}
	}
	fmt.Println("finish!")
}
