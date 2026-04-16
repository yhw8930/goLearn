package main

import "fmt"

// CountSort 计数排序
//
// 【问题】
// 对整数数组进行排序，支持负数
//
// 【解题思路】
// 1. 找最小值和最大值
// 2. 用 value-minVal 作为桶下标统计频率
// 3. 按桶顺序回填数组
//
// 【时空复杂度】
// 时间：O(N + (maxVal-minVal))
// 空间：O(maxVal-minVal)
func CountSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	minVal := arr[0]
	maxVal := arr[0]
	for i := 1; i < len(arr); i++ {
		minVal = min(arr[i], minVal)
		maxVal = max(arr[i], maxVal)
	}
	bucket := make([]int, maxVal-minVal+1)
	for _, v := range arr {
		bucket[v-minVal]++
	}
	index := 0
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			arr[index] = i + minVal
			index++
			bucket[i]--
		}
	}
}

func copyArray1(arr []int) []int {
	if arr == nil {
		return nil
	}
	ans := make([]int, len(arr))
	copy(ans, arr)
	return ans
}

func isEqual1(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func comparator1(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func printArray1(arr []int) {
	fmt.Println(arr)
}

func main() {
	testCases := [][]int{
		{3, -1, 2, -5, 0, 2, -1},
		{-3, -7, -1, -7, -2},
		{-2, -1, 0, 3, 5},
		{0},
		{},
		{5, 4, 3, 2, 1, 0, -1, -2},
		{2, 2, 2, -1, -1, 0, 0},
	}

	for i, tc := range testCases {
		arr1 := copyArray1(tc)
		arr2 := copyArray1(tc)
		CountSort(arr1)
		comparator1(arr2)

		fmt.Printf("case %d\n", i+1)
		fmt.Print("origin: ")
		printArray1(tc)
		fmt.Print("sorted: ")
		printArray1(arr1)

		if !isEqual1(arr1, arr2) {
			fmt.Print("expect: ")
			printArray1(arr2)
			fmt.Println("failed")
			return
		}
		fmt.Println("passed")
	}
}
