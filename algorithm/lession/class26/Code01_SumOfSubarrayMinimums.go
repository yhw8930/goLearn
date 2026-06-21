package main

import "math/rand"

// 子数组最小值之和：给定数组 arr，对它的每一个连续子数组求出其中的最小值，
// 再把所有这些最小值累加起来，返回总和（结果可能很大，最优解对 1e9+7 取模）。
// 测试链接：https://leetcode.com/problems/sum-of-subarray-minimums/
//
// 最优解思路：对每个 arr[i]，统计有多少个子数组以 arr[i] 作为最小值。
// 设 left[i] 为左边最近的“小于等于 arr[i]”的位置，right[i] 为右边最近的“小于 arr[i]”的位置
// （一边取等、一边严格，避免相等元素重复计数）。则以 arr[i] 为最小值的子数组个数为
// (i-left[i]) * (right[i]-i)，贡献 (i-left[i])*(right[i]-i)*arr[i]。

// SubArrayMinSum1 是对数器用的暴力解：枚举所有子数组并扫描求最小值。
//
// 时间复杂度：O(N^3)。
// 空间复杂度：O(1)。
func SubArrayMinSum1(arr []int) int {
	ans := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			min := arr[i]
			for k := i + 1; k <= j; k++ {
				if arr[k] < min {
					min = arr[k]
				}
			}
			ans += min
		}
	}
	return ans
}

// SubArrayMinSum2 是最优解思路，但 left/right 用暴力 O(N^2) 求，便于对照理解。
//
// 时间复杂度：O(N^2)。
// 空间复杂度：O(N)。
func SubArrayMinSum2(arr []int) int {
	left := leftNearLessEqualBrute(arr)
	right := rightNearLessBrute(arr)
	ans := 0
	for i := 0; i < len(arr); i++ {
		start := i - left[i]
		end := right[i] - i
		ans += start * end * arr[i]
	}
	return ans
}

// leftNearLessEqualBrute 暴力求每个位置左边最近的“<=”位置，找不到记 -1。
func leftNearLessEqualBrute(arr []int) []int {
	n := len(arr)
	left := make([]int, n)
	for i := 0; i < n; i++ {
		ans := -1
		for j := i - 1; j >= 0; j-- {
			if arr[j] <= arr[i] {
				ans = j
				break
			}
		}
		left[i] = ans
	}
	return left
}

// rightNearLessBrute 暴力求每个位置右边最近的“<”位置，找不到记 N。
func rightNearLessBrute(arr []int) []int {
	n := len(arr)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		ans := n
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				ans = j
				break
			}
		}
		right[i] = ans
	}
	return right
}

// SumSubarrayMins 是最优解：用单调栈在 O(N) 内求出 left/right，并对 1e9+7 取模。
//
// 时间复杂度：O(N)。
// 空间复杂度：O(N)。
func SumSubarrayMins(arr []int) int {
	left := nearLessEqualLeft(arr)
	right := nearLessRight(arr)
	var ans int64
	const mod = 1000000007
	for i := 0; i < len(arr); i++ {
		start := int64(i - left[i])
		end := int64(right[i] - i)
		ans += start * end * int64(arr[i])
		ans %= mod
	}
	return int(ans)
}

// nearLessEqualLeft 单调栈求每个位置左边最近的“<=”位置（找不到记 -1）。
func nearLessEqualLeft(arr []int) []int {
	n := len(arr)
	left := make([]int, n)
	stack := make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[i] <= arr[stack[len(stack)-1]] {
			left[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		left[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}
	return left
}

// nearLessRight 单调栈求每个位置右边最近的“<”位置（找不到记 N）。
func nearLessRight(arr []int) []int {
	n := len(arr)
	right := make([]int, n)
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		right[stack[len(stack)-1]] = n
		stack = stack[:len(stack)-1]
	}
	return right
}

// subArrayMinRandomArray 生成长度 len、元素 1~maxValue 的随机数组。
func subArrayMinRandomArray(length, maxValue int) []int {
	ans := make([]int, length)
	for i := range ans {
		ans[i] = rand.Intn(maxValue) + 1
	}
	return ans
}
