package main

import "fmt"

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
