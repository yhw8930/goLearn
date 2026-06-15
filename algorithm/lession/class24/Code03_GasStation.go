package main

// 测试链接：https://leetcode.com/problems/gas-station

// CanCompleteCircuit 解决“加油站绕一圈”问题：
// gas[i] 表示第 i 个加油站能加到的油量，cost[i] 表示从第 i 站开到下一站要消耗的油量。
// 车从某个站出发时油箱为空，沿环形路线依次经过所有站，途中油量不能为负。
// 返回任意一个可以绕完整圈的出发站下标；如果不存在这样的站，返回 -1。
// 这个实现先用 GoodArray 找出所有可行出发点，再返回第一个可行位置。
//
// 时间复杂度：O(N)。
// 空间复杂度：O(N)。
func CanCompleteCircuit(gas, cost []int) int {
	good := GoodArray(gas, cost)
	for i := 0; i < len(gas); i++ {
		if good[i] {
			return i
		}
	}
	return -1
}

// GoodArray 返回每个加油站是否可以作为绕完整圈的出发点。
// 把 gas[i]-cost[i] 复制两份并做前缀和。对每个长度为 N 的区间，
// 只要区间内前缀和最小值减去出发点前的前缀和不小于 0，就说明从该点出发全程油量不为负。
// 用滑动窗口最小值结构在 O(N) 时间内检查所有起点。
//
// 时间复杂度：O(N)。
// 空间复杂度：O(N)。
func GoodArray(g, c []int) []bool {
	n := len(g)
	m := n << 1
	arr := make([]int, m)
	for i := 0; i < n; i++ {
		arr[i] = g[i] - c[i]
		arr[i+n] = g[i] - c[i]
	}
	for i := 1; i < m; i++ {
		arr[i] += arr[i-1]
	}
	w := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(w) > 0 && arr[w[len(w)-1]] >= arr[i] {
			w = w[:len(w)-1]
		}
		w = append(w, i)
	}
	ans := make([]bool, n)
	for offset, i, j := 0, 0, n; j < m; j++ {
		if arr[w[0]]-offset >= 0 {
			ans[i] = true
		}
		if w[0] == i {
			w = w[1:]
		}
		for len(w) > 0 && arr[w[len(w)-1]] >= arr[j] {
			w = w[:len(w)-1]
		}
		w = append(w, j)
		offset = arr[i]
		i++
	}
	return ans
}
