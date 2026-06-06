package main

import "fmt"

// 题目：给定一个整数数组，要求把数组原地调整为升序。
// 插入排序假设左侧 0..i-1 已经有序，处理 i 位置时让它向左移动到合适位置。
// 当新数不再小于左边数字时停止，左侧 0..i 继续保持有序。
// 核心思路：像整理手牌一样逐个插入，适合小规模或基本有序的数据。
// 时间复杂度：最坏 O(N^2)，数组接近有序时可接近 O(N)。
// 空间复杂度：O(1)。

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	insertSort(arr)
	fmt.Println(arr)
}

// insertSort 插入排序
// 构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
//
// 时间复杂度: O(n²)。在最好的情况下（数组已经排序），时间复杂度为 O(n)。
// 空间复杂度: O(1)。这是一个原地排序算法。
// 稳定性: 稳定。相等的元素不会改变它们的相对顺序。
// insertSort 是标准插入排序实现。
// 它从左到右扩展有序区，每次让当前数向左交换到合适位置。
func insertSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr)-1; i++ {
		cur := arr[i]
		preIndex := i - 1
		for preIndex >= 0 && arr[preIndex] > cur {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = cur
	}
}

// insertSortInvalid 保留一个容易出错的写法示例。
// 它用于对比标准插入排序边界和循环条件，不作为推荐实现。
func insertSortInvalid(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
