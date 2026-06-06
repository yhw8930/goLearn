package main

import "fmt"

// 题目：一排纸牌分数已知，两名玩家轮流从最左或最右拿一张，双方都绝顶聪明，求最后赢家分数。
// 当前玩家拿牌后，下一轮会变成对手在剩余区间先手。
// 核心思路：定义 f(l,r) 表示当前玩家先手能拿到的最好分数，g(l,r) 表示当前玩家后手能拿到的最好分数。
// f 会选择左右两种拿法中的最大收益，g 面对对手最优选择，只能得到两种后续中的较小收益。
// 时间复杂度：暴力 O(2^N)，记忆化和动态规划 O(N^2)。
// 空间复杂度：暴力 O(N)，记忆化和动态规划 O(N^2)。

// 纸牌排成一行，数组 arr（arr = {5, 7, 4, 5}）表示每张牌的分数，两个人轮流拿，每次只能拿最左或最右，双方都绝顶聪明，求最后赢家分数。
func main() {
	arr := []int{5, 7, 4, 5, 8, 1, 6, 0, 3, 4, 6, 1, 7}
	fmt.Println(win1(arr))
	fmt.Println(win2(arr))
	fmt.Println(win3(arr))
}

// win1 暴力递归求双方都最优时最后赢家的分数。
// 先手函数 f1 和后手函数 g1 互相递归，分别表示当前范围内先拿和后拿能获得的最好分数。
//
// 时间复杂度：O(2^N)
// 空间复杂度：O(N)
func win1(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	first := f1(arr, 0, len(arr)-1)
	second := g1(arr, 0, len(arr)-1)
	return max(first, second)
}

// f1 表示在 arr[l..r] 范围内当前玩家先手拿牌能获得的最好分数。
func f1(arr []int, l, r int) int {
	if l == r {
		return arr[l]
	}
	pl := arr[l] + g1(arr, l+1, r)
	pr := arr[r] + g1(arr, l, r-1)
	return max(pl, pr)
}

// g1 表示在 arr[l..r] 范围内当前玩家后手拿牌能获得的最好分数。
// 对手会留下较差的后续局面，所以这里取两种可能中的较小值。
func g1(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	pl := f1(arr, l+1, r)
	pr := f1(arr, l, r-1)
	return min(pl, pr)
}

// win2 在 win1 的递归基础上增加缓存表，避免重复计算同一段 l/r 范围。
//
// 时间复杂度：O(N^2)
// 空间复杂度：O(N^2)
func win2(arr []int) int {
	N := len(arr)
	if N == 0 {
		return 0
	}
	fmap := make([][]int, N)
	gmap := make([][]int, N)
	for i := 0; i < N; i++ {
		fmap[i] = make([]int, N)
		gmap[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmap[i][j] = -1
			gmap[i][j] = -1
		}
	}
	first := f2(arr, 0, N-1, fmap, gmap)
	second := g2(arr, 0, N-1, fmap, gmap)
	return max(first, second)
}

// f2 是带缓存的先手函数，fmap[l][r] 记录 arr[l..r] 先手最好分数。
func f2(arr []int, l, r int, fmap, gmap [][]int) int {
	if fmap[l][r] != -1 {
		return fmap[l][r]
	}
	ans := 0
	if l == r {
		ans = arr[l]
	} else {
		pl := arr[l] + g2(arr, l+1, r, fmap, gmap)
		pr := arr[r] + g2(arr, l, r-1, fmap, gmap)
		ans = max(pl, pr)
	}
	fmap[l][r] = ans
	return ans
}

// g2 是带缓存的后手函数，gmap[l][r] 记录 arr[l..r] 后手最好分数。
func g2(arr []int, l, r int, fmap, gmap [][]int) int {
	if gmap[l][r] != -1 {
		return gmap[l][r]
	}
	ans := 0
	if l == r {
		ans = 0
	} else {
		pl := f2(arr, l+1, r, fmap, gmap)
		pr := f2(arr, l, r-1, fmap, gmap)
		ans = min(pl, pr)
	}
	gmap[l][r] = ans
	return ans
}

// win3 使用严格动态规划，从短区间推到长区间。
// fmap[l][r] 和 gmap[l][r] 分别表示 arr[l..r] 上先手、后手能获得的最好分数。
//
// 时间复杂度：O(N^2)
// 空间复杂度：O(N^2)
func win3(arr []int) int {
	N := len(arr)
	if N == 0 {
		return 0
	}
	fmap := make([][]int, N)
	gmap := make([][]int, N)
	for i := 0; i < N; i++ {
		fmap[i] = make([]int, N)
		gmap[i] = make([]int, N)
		fmap[i][i] = arr[i]
	}
	for startCol := 1; startCol < N; startCol++ {
		l := 0
		r := startCol
		for r < N {
			fmap[l][r] = max(arr[l]+gmap[l+1][r], arr[r]+gmap[l][r-1])
			gmap[l][r] = min(fmap[l+1][r], fmap[l][r-1])
			l++
			r++
		}
	}
	return max(fmap[0][N-1], gmap[0][N-1])
}
