package main

// Permutation1 使用“剩余字符列表 + 已形成路径”生成全排列。
//
// 时间复杂度：O(N*N!)，N! 个排列，每个排列长度 N。
// 空间复杂度：O(N)，不计返回结果；计结果为 O(N*N!)。
func Permutation1(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	rest := []rune(s)
	ans := make([]string, 0)
	permutationByRest(rest, "", &ans)
	return ans
}

func permutationByRest(rest []rune, path string, ans *[]string) {
	if len(rest) == 0 {
		*ans = append(*ans, path)
		return
	}
	for i := 0; i < len(rest); i++ {
		cur := rest[i]
		nextRest := append([]rune{}, rest[:i]...)
		nextRest = append(nextRest, rest[i+1:]...)
		permutationByRest(nextRest, path+string(cur), ans)
	}
}

// Permutation2 使用原地交换生成全排列。
// index 左侧已经决定，index 右侧继续枚举哪个字符放到 index 位置。
//
// 时间复杂度：O(N*N!)。
// 空间复杂度：O(N)，递归深度；计结果为 O(N*N!)。
func Permutation2(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	str := []rune(s)
	ans := make([]string, 0)
	permutationBySwap(str, 0, &ans)
	return ans
}

func permutationBySwap(str []rune, index int, ans *[]string) {
	if index == len(str) {
		*ans = append(*ans, string(str))
		return
	}
	for i := index; i < len(str); i++ {
		str[index], str[i] = str[i], str[index]
		permutationBySwap(str, index+1, ans)
		str[index], str[i] = str[i], str[index]
	}
}

// Permutation3 在每一层用 visited 去重，避免同一个字符重复放到 index 位置。
//
// 时间复杂度：O(N*N!)，有重复字符时实际结果更少。
// 空间复杂度：O(N)，不计返回结果。
func Permutation3(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	str := []rune(s)
	ans := make([]string, 0)
	permutationNoRepeat(str, 0, &ans)
	return ans
}

func permutationNoRepeat(str []rune, index int, ans *[]string) {
	if index == len(str) {
		*ans = append(*ans, string(str))
		return
	}
	visited := make(map[rune]struct{})
	for i := index; i < len(str); i++ {
		if _, ok := visited[str[i]]; ok {
			continue
		}
		visited[str[i]] = struct{}{}
		str[index], str[i] = str[i], str[index]
		permutationNoRepeat(str, index+1, ans)
		str[index], str[i] = str[i], str[index]
	}
}
