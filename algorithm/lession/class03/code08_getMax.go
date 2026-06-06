package main

import "fmt"

// 题目：用递归方式求数组中的最大值。
// 把数组范围 L..R 分成左右两半，分别求出两边最大值。
// 当范围只有一个数时，它就是该范围的最大值。
// 核心思路：递归分解问题，回溯时用 max 合并左右子问题答案。
// 时间复杂度：O(N)。
// 空间复杂度：O(logN)，递归深度来自二分。

func GetMax(arr []int) int {
	return process(arr, 0, len(arr)-1)
}

func process(arr []int, L, R int) int {
	if L == R {
		return arr[L]
	}
	mid := L + ((R - L) >> 1)
	leftMax := process(arr, L, mid)
	rightMax := process(arr, mid+1, R)

	if leftMax > rightMax {
		return leftMax
	}
	return rightMax
}

func main() {
	arr := []int{3, 1, 5, 0, 2, 9, 7}
	fmt.Println(GetMax(arr)) // 输出 9
}
