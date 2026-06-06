package main

import "fmt"

// 题目：给定一个相邻元素不相等的无序数组，返回任意一个局部最小值位置。
// 局部最小指该位置的值小于它相邻位置的值；边界位置只需要和唯一邻居比较。
// 如果开头或结尾已经是局部最小，直接返回；否则数组两端形成向内下降趋势。
// 核心思路：看中点和左右邻居的大小，局部最小一定存在于下降的一侧，因此可以二分。
// 前提：任意相邻元素不相等。
// 时间复杂度：O(logN)。
// 空间复杂度：O(1)。

func main() {
	arr1 := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	arr2 := []int{19, 17}
	arr3 := []int{19, 17, 19, 15, 19}
	arr4 := []int{19, 17, 19, 15, 19, 23, 2, 45, 3, 4}
	fmt.Println(oneMinIndex(arr1))
	fmt.Println(oneMinIndex(arr2))
	fmt.Println(oneMinIndex(arr3))
	fmt.Println(oneMinIndex(arr4))
}

func oneMinIndex(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 || arr[0] < arr[1] {
		return 0
	}
	n := len(arr)
	if arr[n-1] < arr[n-2] {
		return n - 1
	}
	left, right := 1, n-2
	for left < right {
		mid := left + (right-left)>>1
		if arr[mid] > arr[mid-1] {
			right = mid - 1
		} else if arr[mid] > arr[mid+1] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
