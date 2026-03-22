package main

import (
	"fmt"
)

// 冒泡排序核心逻辑（完全对齐Java版）
func bubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 外层循环：控制比较的右边界，从len(arr)-1递减到1
	for e := len(arr) - 1; e > 0; e-- {
		// 内层循环：0~e-1 两两比较，大的元素往后冒泡
		for i := 0; i < e; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
			}
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
