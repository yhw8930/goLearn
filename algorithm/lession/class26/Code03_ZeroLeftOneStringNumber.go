package main

// 0 左 1 字符串计数：规定一个由 0 和 1 组成的字符串中，每个 0 的左边必须紧跟着一个 1
// （即 0 不能在开头，也不能有两个连续的 0；1 可以随意出现）。
// 给定长度 n，问长度恰好为 n 的合法字符串有多少个。
//
// 分析：设 f(i) 表示从第 i 位开始往后填到第 n 位的合法方案数。第 i 位放 1，则后面从 i+1 继续；
// 第 i 位放 1 且第 i+1 位放 0（“10”整体占两位），则后面从 i+2 继续。于是 f(i)=f(i+1)+f(i+2)，
// 本质就是斐波那契数列，答案满足 getNum(n)=Fib(n+1)。

// GetNum1 暴力递归。时间复杂度：O(2^N)。空间复杂度：O(N)。
func GetNum1(n int) int {
	if n < 1 {
		return 0
	}
	return zeroOneProcess(1, n)
}

// zeroOneProcess 表示当前来到第 i 位、总长 n 时，后续的合法填法数量。
func zeroOneProcess(i, n int) int {
	if i == n-1 {
		return 2
	}
	if i == n {
		return 1
	}
	return zeroOneProcess(i+1, n) + zeroOneProcess(i+2, n)
}

// GetNum2 迭代。时间复杂度：O(N)。空间复杂度：O(1)。
func GetNum2(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	pre, cur := 1, 1
	for i := 2; i < n+1; i++ {
		cur, pre = cur+pre, cur
	}
	return cur
}

// GetNum3 矩阵快速幂。时间复杂度：O(logN)。空间复杂度：O(1)。
func GetNum3(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	base := [][]int{{1, 1}, {1, 0}}
	res := zeroOneMatrixPower(base, n-2)
	return 2*res[0][0] + res[1][0]
}

// zeroOneMatrixPower 用快速幂求方阵 m 的 p 次幂，初始为单位矩阵。
//
// 时间复杂度：O(K^3 * logP)。
// 空间复杂度：O(K^2)。
func zeroOneMatrixPower(m [][]int, p int) [][]int {
	res := make([][]int, len(m))
	for i := range res {
		res[i] = make([]int, len(m[0]))
		res[i][i] = 1
	}
	tmp := m
	for ; p != 0; p >>= 1 {
		if p&1 != 0 {
			res = zeroOneMuliMatrix(res, tmp)
		}
		tmp = zeroOneMuliMatrix(tmp, tmp)
	}
	return res
}

// zeroOneMuliMatrix 返回两个矩阵相乘的结果。
func zeroOneMuliMatrix(m1, m2 [][]int) [][]int {
	res := make([][]int, len(m1))
	for i := range res {
		res[i] = make([]int, len(m2[0]))
	}
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return res
}
