package main

// 斐波那契类问题与矩阵快速幂：本文件演示三类满足线性递推的数列，
// 每类都给出三种解法——暴力递归、迭代、以及 O(logN) 的矩阵快速幂。
// 关键结论：任何严格线性递推（如 F(n)=F(n-1)+F(n-2)）都能写成“状态向量 = 基矩阵^k × 初始向量”，
// 用快速幂把矩阵幂降到 O(logN)（这里不考虑大整数溢出与取模）。

// ---------- 第一类：标准斐波那契 F(1)=F(2)=1, F(n)=F(n-1)+F(n-2) ----------

// Fib1 暴力递归。时间复杂度：O(2^N)。空间复杂度：O(N)（递归栈）。
func Fib1(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return Fib1(n-1) + Fib1(n-2)
}

// Fib2 自底向上迭代。时间复杂度：O(N)。空间复杂度：O(1)。
func Fib2(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	res, pre := 1, 1
	for i := 3; i <= n; i++ {
		res, pre = res+pre, res
	}
	return res
}

// Fib3 矩阵快速幂。时间复杂度：O(logN)。空间复杂度：O(1)。
func Fib3(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	base := [][]int{{1, 1}, {1, 0}}
	res := fibMatrixPower(base, n-2)
	return res[0][0] + res[1][0]
}

// ---------- 第二类：S(1)=1, S(2)=2, S(n)=S(n-1)+S(n-2)（与斐波那契同递推、不同初值）----------

// Step1 暴力递归。时间复杂度：O(2^N)。空间复杂度：O(N)。
func Step1(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	return Step1(n-1) + Step1(n-2)
}

// Step2 迭代。时间复杂度：O(N)。空间复杂度：O(1)。
func Step2(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	res, pre := 2, 1
	for i := 3; i <= n; i++ {
		res, pre = res+pre, res
	}
	return res
}

// Step3 矩阵快速幂。时间复杂度：O(logN)。空间复杂度：O(1)。
func Step3(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	base := [][]int{{1, 1}, {1, 0}}
	res := fibMatrixPower(base, n-2)
	return 2*res[0][0] + res[1][0]
}

// ---------- 第三类：母牛问题 C(1)=1,C(2)=2,C(3)=3, C(n)=C(n-1)+C(n-3) ----------
// 每头母牛长到第 3 年开始每年生一头小母牛，牛不会死，求第 n 年牛的总数。

// Cow1 暴力递归。时间复杂度：O(2^N)。空间复杂度：O(N)。
func Cow1(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 || n == 3 {
		return n
	}
	return Cow1(n-1) + Cow1(n-3)
}

// Cow2 迭代。时间复杂度：O(N)。空间复杂度：O(1)。
func Cow2(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 || n == 3 {
		return n
	}
	res, pre, prepre := 3, 2, 1
	for i := 4; i <= n; i++ {
		res, pre, prepre = res+prepre, res, pre
	}
	return res
}

// Cow3 矩阵快速幂（3×3 基矩阵）。时间复杂度：O(logN)。空间复杂度：O(1)。
func Cow3(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 || n == 3 {
		return n
	}
	base := [][]int{
		{1, 1, 0},
		{0, 0, 1},
		{1, 0, 0},
	}
	res := fibMatrixPower(base, n-3)
	return 3*res[0][0] + 2*res[1][0] + res[2][0]
}

// fibMatrixPower 用快速幂求方阵 m 的 p 次幂，初始为单位矩阵。
//
// 时间复杂度：O(K^3 * logP)，K 为矩阵阶数。
// 空间复杂度：O(K^2)。
func fibMatrixPower(m [][]int, p int) [][]int {
	res := make([][]int, len(m))
	for i := range res {
		res[i] = make([]int, len(m[0]))
		res[i][i] = 1 // 单位矩阵
	}
	t := m // 矩阵的 1 次方
	for ; p != 0; p >>= 1 {
		if p&1 != 0 {
			res = fibMuliMatrix(res, t)
		}
		t = fibMuliMatrix(t, t)
	}
	return res
}

// fibMuliMatrix 返回两个矩阵相乘的结果。
func fibMuliMatrix(m1, m2 [][]int) [][]int {
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
