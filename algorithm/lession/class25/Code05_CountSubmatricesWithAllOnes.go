package main

// 统计全 1 子矩形的数量：给定一个只含 0 和 1 的二维矩阵 mat，
// 统计其中只由 1 组成的子矩形（子矩阵）一共有多少个，返回这个数量。
// 测试链接：https://leetcode.com/problems/count-submatrices-with-all-ones
//
// 核心思想：逐行向下压缩成直方图 height（遇到 0 清零）。对每个以当前行为底的直方图，
// 用单调栈统计“必须以这一行为底边”的全 1 子矩形数量并累加。
// 弹出 cur 时，它高出左右两侧（左侧 left、右侧 i）的部分共 (height[cur]-down) 层，
// 每一层宽度为 n=i-left-1 的范围内可产生 num(n)=n*(n+1)/2 个连续子段。

// NumSubmat 逐行构造直方图并累加每行底边贡献的全 1 子矩形数量。
//
// 时间复杂度：O(R*C)。
// 空间复杂度：O(C)。
func NumSubmat(mat [][]int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}
	nums := 0
	height := make([]int, len(mat[0]))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				height[j] = 0
			} else {
				height[j]++
			}
		}
		nums += countFromBottom(height)
	}
	return nums
}

// countFromBottom 统计以 height 为直方图时，必须贴着底边的全 1 子矩形数量。
//
// 时间复杂度：O(C)。
// 空间复杂度：O(C)。
func countFromBottom(height []int) int {
	if len(height) == 0 {
		return 0
	}
	nums := 0
	stack := make([]int, len(height))
	si := -1
	for i := 0; i < len(height); i++ {
		for si != -1 && height[stack[si]] >= height[i] {
			cur := stack[si]
			si--
			if height[cur] > height[i] {
				left := -1
				if si != -1 {
					left = stack[si]
				}
				n := i - left - 1
				down := height[i]
				if left != -1 && height[left] > down {
					down = height[left]
				}
				nums += (height[cur] - down) * submatNum(n)
			}
		}
		si++
		stack[si] = i
	}
	for si != -1 {
		cur := stack[si]
		si--
		left := -1
		if si != -1 {
			left = stack[si]
		}
		n := len(height) - left - 1
		down := 0
		if left != -1 {
			down = height[left]
		}
		nums += (height[cur] - down) * submatNum(n)
	}
	return nums
}

// submatNum 返回长度为 n 的范围内连续子段的数量，即 n*(n+1)/2。
func submatNum(n int) int {
	return (n * (1 + n)) >> 1
}
