package main

// 直方图中的最大矩形：给定数组 height 表示一排相邻柱子的高度（宽度都为 1），
// 求在这些柱子之间能勾勒出的面积最大的矩形，返回该面积。
// 测试链接：https://leetcode.com/problems/largest-rectangle-in-histogram
//
// 核心思想：用单调栈（栈底到栈顶高度递增）。当 height[i] 不大于栈顶高度时弹出 j，
// 以 height[j] 为高度的矩形向右扩到 i（i 是右边第一个更矮的），
// 向左扩到弹出后的新栈顶 k（k 是左边第一个更矮的），宽度为 i-k-1。

// LargestRectangleArea1 使用 Go 切片当作单调栈。
//
// 时间复杂度：O(N)。
// 空间复杂度：O(N)。
func LargestRectangleArea1(height []int) int {
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

// LargestRectangleArea2 用一个定长数组 + 栈顶指针 si 模拟栈，常数更优。
//
// 时间复杂度：O(N)。
// 空间复杂度：O(N)。
func LargestRectangleArea2(height []int) int {
	if len(height) == 0 {
		return 0
	}
	n := len(height)
	stack := make([]int, n)
	si := -1
	maxArea := 0
	for i := 0; i < n; i++ {
		for si != -1 && height[i] <= height[stack[si]] {
			j := stack[si]
			si--
			k := -1
			if si != -1 {
				k = stack[si]
			}
			if curArea := (i - k - 1) * height[j]; curArea > maxArea {
				maxArea = curArea
			}
		}
		si++
		stack[si] = i
	}
	for si != -1 {
		j := stack[si]
		si--
		k := -1
		if si != -1 {
			k = stack[si]
		}
		if curArea := (n - k - 1) * height[j]; curArea > maxArea {
			maxArea = curArea
		}
	}
	return maxArea
}
