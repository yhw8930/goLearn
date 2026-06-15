package main

import "math/rand"

// SlidingWindowMaxRight 解决“固定大小窗口最大值数组”问题：
// 给定数组 arr 和窗口大小 w，窗口从最左侧开始，每次向右滑动一个位置。
// 每个窗口都要收集窗口内的最大值，最终返回所有窗口最大值组成的数组。
// 如果 arr 为空、w 小于 1，或数组长度小于 w，说明无法形成合法窗口，返回 nil。
// 这个版本是对数器用的暴力方法：每个窗口都重新扫描一遍求最大值。
//
// 时间复杂度：O(N*W)。
// 空间复杂度：O(1)，不计算返回数组。
func SlidingWindowMaxRight(arr []int, w int) []int {
	if len(arr) == 0 || w < 1 || len(arr) < w {
		return nil
	}
	n := len(arr)
	res := make([]int, n-w+1)
	index := 0
	l := 0
	r := w - 1
	for r < n {
		maxValue := arr[l]
		for i := l + 1; i <= r; i++ {
			if arr[i] > maxValue {
				maxValue = arr[i]
			}
		}
		res[index] = maxValue
		index++
		l++
		r++
	}
	return res
}

// GetMaxWindow 返回 arr 中每个长度为 w 的滑动窗口最大值。
// 用一个双端队列保存下标，队列中对应的值从头到尾严格递减；队头始终是当前窗口最大值下标。
//
// 时间复杂度：O(N)，每个下标最多进队、出队一次。
// 空间复杂度：O(W)，双端队列空间。
func GetMaxWindow(arr []int, w int) []int {
	if len(arr) == 0 || w < 1 || len(arr) < w {
		return nil
	}
	qmax := make([]int, 0, w)
	res := make([]int, len(arr)-w+1)
	index := 0
	for r := 0; r < len(arr); r++ {
		for len(qmax) > 0 && arr[qmax[len(qmax)-1]] <= arr[r] {
			qmax = qmax[:len(qmax)-1]
		}
		qmax = append(qmax, r)
		if qmax[0] == r-w {
			qmax = qmax[1:]
		}
		if r >= w-1 {
			res[index] = arr[qmax[0]]
			index++
		}
	}
	return res
}

func slidingWindowMaxGenerateRandomArray(maxSize, maxValue int) []int {
	arr := make([]int, rand.Intn(maxSize+1))
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue + 1)
	}
	return arr
}

func slidingWindowMaxIsEqual(arr1, arr2 []int) bool {
	if (arr1 == nil) != (arr2 == nil) {
		return false
	}
	if arr1 == nil && arr2 == nil {
		return true
	}
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
