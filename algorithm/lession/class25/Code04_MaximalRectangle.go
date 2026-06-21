package main

// 最大全 1 子矩形：给定一个只含字符 '0' 和 '1' 的二维矩阵 matrix，
// 找出其中只由 '1' 组成的、面积最大的矩形，返回该面积。
// 测试链接：https://leetcode.com/problems/maximal-rectangle/
//
// 核心思想：逐行向下压缩成直方图——height[j] 表示以当前行为底、第 j 列上方连续 '1' 的个数
// （遇到 '0' 清零）。对每一行得到的直方图调用“直方图最大矩形”即可。

// MaximalRectangle 逐行构造直方图并取所有行结果的最大值。
//
// 时间复杂度：O(R*C)，R、C 为矩阵行列数。
// 空间复杂度：O(C)。
func MaximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxArea := 0
	height := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				height[j] = 0
			} else {
				height[j]++
			}
		}
		if area := maxRecFromBottom(height); area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// maxRecFromBottom 求以 height 为直方图时的最大矩形面积（单调栈）。
//
// 时间复杂度：O(C)。
// 空间复杂度：O(C)。
func maxRecFromBottom(height []int) int {
	if len(height) == 0 {
		return 0
	}
	maxArea := 0
	stack := make([]int, 0, len(height))
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] <= height[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k := -1
			if len(stack) > 0 {
				k = stack[len(stack)-1]
			}
			if curArea := (i - k - 1) * height[j]; curArea > maxArea {
				maxArea = curArea
			}
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		j := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k := -1
		if len(stack) > 0 {
			k = stack[len(stack)-1]
		}
		if curArea := (len(height) - k - 1) * height[j]; curArea > maxArea {
			maxArea = curArea
		}
	}
	return maxArea
}
