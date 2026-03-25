package main

import (
	"fmt"
)

// BubbleSort 冒泡排序
//
// 思路:
// 重复地遍历要排序的数列，一次比较两个元素，如果它们的顺序错误就把它们交换过来。
// 遍历数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。
//
// 复杂度分析:
// - 时间复杂度: O(n²)。在最好的情况下（数组已经排序），时间复杂度为 O(n)。
// - 空间复杂度: O(1)。这是一个原地排序算法。
// - 稳定性: 稳定。相等的元素不会改变它们的相对顺序。
func bubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for e := len(arr) - 1; e > 0; e-- {
		swapped := false // 交换标志位
		for i := 0; i < e; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// 主测试函数
func main() {
	testTime := 500000 // 测试次数
	maxSize := 100     // 数组最大长度
	maxValue := 100    // 元素绝对值最大值
	succeed := true

	// 批量测试：验证冒泡排序正确性
	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := copyArray(arr1)

		bubbleSort(arr1)
		comparator(arr2)

		if !isEqual(arr1, arr2) {
			succeed = false
			fmt.Println("排序结果不一致：")
			printArray(arr1)
			printArray(arr2)
			break
		}
	}

	// 输出测试结果
	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fucking fucked!")
	}

	// 单独测试一个示例
	fmt.Println("\n示例测试：")
	arr := generateRandomArray(maxSize, maxValue)
	fmt.Print("排序前：")
	printArray(arr)
	bubbleSort(arr)
	fmt.Print("排序后：")
	printArray(arr)
}
