package main

import "fmt"

// 题目：给定升序数组 arr 和目标值 num，返回小于等于 num 的最右位置。
// 如果不存在这样的数，返回 -1。这个问题找的是满足条件的右边界。
// 二分时只要 arr[mid] <= num，mid 就是候选答案，但右侧可能还有更靠后的可行位置。
// 核心思路：记录答案后继续向右收缩，最终得到最右的可行下标。
// 前提：数组必须整体有序。
// 时间复杂度：O(logN)。
// 空间复杂度：O(1)。

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7
	fmt.Println(mostRightNoMoreNumIndex(arr, target))
}

// arr有序， 找出<=num最右的位置
func mostRightNoMoreNumIndex(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	left, right := 0, len(arr)-1
	index := -1
	for left <= right {
		mid := left + (right-left)>>1
		if arr[mid] <= target {
			left = mid + 1
			index = mid
		} else {
			right = mid - 1
		}
	}
	return index
}
