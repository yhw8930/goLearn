package main

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// NQueensNum1 解决 N 皇后问题：
// 在 n*n 的棋盘上摆放 n 个皇后，要求任意两个皇后都不能共行、共列、共斜线。
// 因为每一行必须且只能放一个皇后，所以按行从上到下尝试每一列。
// record[row] 记录 row 行皇后放在哪一列，放新皇后前检查它是否和之前行的皇后冲突。
// 返回所有合法摆法数量；n < 1 时返回 0。
//
// 时间复杂度：O(N!) 量级。
// 空间复杂度：O(N)，递归调用栈和记录数组。
func NQueensNum1(n int) int {
	if n < 1 {
		return 0
	}
	record := make([]int, n)
	return nQueensProcess1(0, record, n)
}

func nQueensProcess1(i int, record []int, n int) int {
	if i == n {
		return 1
	}
	res := 0
	for j := 0; j < n; j++ {
		if nQueensIsValid(record, i, j) {
			record[i] = j
			res += nQueensProcess1(i+1, record, n)
		}
	}
	return res
}

func nQueensIsValid(record []int, i, j int) bool {
	for k := 0; k < i; k++ {
		if j == record[k] || absInt(record[k]-j) == absInt(i-k) {
			return false
		}
	}
	return true
}

// NQueensNum2 返回 n 皇后问题的合法摆法数量，使用位运算加速。
// limit 标记棋盘可用列，colLim、leftDiaLim、rightDiaLim 分别记录列和两类对角线限制。
//
// 时间复杂度：O(N!) 量级，但常数远小于普通递归。
// 空间复杂度：O(N)，递归调用栈深度。
func NQueensNum2(n int) int {
	if n < 1 || n > 32 {
		return 0
	}
	var limit uint32
	if n == 32 {
		limit = ^uint32(0)
	} else {
		limit = (uint32(1) << n) - 1
	}
	return nQueensProcess2(limit, 0, 0, 0)
}

func nQueensProcess2(limit, colLim, leftDiaLim, rightDiaLim uint32) int {
	if colLim == limit {
		return 1
	}
	pos := limit & ^(colLim | leftDiaLim | rightDiaLim)
	res := 0
	for pos != 0 {
		mostRightOne := pos & -pos
		pos -= mostRightOne
		res += nQueensProcess2(limit, colLim|mostRightOne, (leftDiaLim|mostRightOne)<<1, (rightDiaLim|mostRightOne)>>1)
	}
	return res
}
