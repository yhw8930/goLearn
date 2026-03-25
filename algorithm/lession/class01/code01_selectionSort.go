package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// SelectionSort 选择排序
//
// 思路:
// 首先在未排序序列中找到最小（或最大）元素，存放到排序序列的起始位置，
// 然后再从剩余未排序元素中继续寻找最小（或最大）元素，然后放到已排序序列的末尾。
//
// 复杂度分析:
// - 时间复杂度: O(n²)。选择排序的性能与输入数据无关。
// - 空间复杂度: O(1)。这是一个原地排序算法。
// - 稳定性: 不稳定。在交换过程中，相等元素的相对位置可能会改变。
func selectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			swap(arr, i, minIndex)
		}
	}
}

// 交换切片中两个位置的元素
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i] // i==j, err. invalid memory address or nil pointer dereference
}

// 对比函数：用Go内置排序作为基准
func comparator(arr []int) {
	sort.Ints(arr)
}

// 生成随机数组（模拟Java的generateRandomArray）
// maxSize: 数组最大长度, maxValue: 元素绝对值最大值
func generateRandomArray(maxSize, maxValue int) []int {
	// 初始化随机数种子（保证每次运行生成不同数组）
	rand.Seed(time.Now().UnixNano())
	// 生成[0, maxSize]范围内的随机长度
	length := rand.Intn(maxSize + 1)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		// 生成[-maxValue, maxValue]范围内的随机整数
		positive := rand.Intn(maxValue + 1)
		negative := rand.Intn(maxValue)
		arr[i] = positive - negative
	}
	return arr
}

// 复制切片（模拟Java的copyArray）
func copyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	res := make([]int, len(arr))
	copy(res, arr) // Go内置的copy函数，高效复制切片
	return res
}

// 判断两个切片是否相等（模拟Java的isEqual）
func isEqual(arr1, arr2 []int) bool {
	if (arr1 == nil && arr2 != nil) || (arr1 != nil && arr2 == nil) {
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

// 打印切片
func printArray(arr []int) {
	if arr == nil {
		return
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

// 主测试函数
func main() {
	testTime := 500000 // 测试次数
	maxSize := 100     // 数组最大长度
	maxValue := 100    // 元素绝对值最大值
	succeed := true

	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := copyArray(arr1)

		selectionSort(arr1)
		comparator(arr2)

		if !isEqual(arr1, arr2) {
			succeed = false
			fmt.Println("排序结果不一致：")
			printArray(arr1)
			printArray(arr2)
			break
		}
	}

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
	selectionSort(arr)
	fmt.Print("排序后：")
	printArray(arr)
}
