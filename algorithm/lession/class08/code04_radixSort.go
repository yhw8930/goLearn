package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// RadixSort 基数排序（非负整数）
//
// 【问题】
// 对非负整数数组进行排序
//
// 【解题思路】
// 1. 计算最大值的位数
// 2. 从低位到高位依次进行稳定计数排序
// 3. 每一位使用count数组做桶排序
//
// 【时空复杂度】
// 时间：O(N * digit)
// 空间：O(N + 10)
func RadixSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	radixSort(arr, maxBits(arr))
}

// maxBits 计算最大数字的位数
//
// 【问题】
// 求数组中最大值的十进制位数
//
// 【解题思路】
// 找最大值，不断除10统计位数
//
// 【时空复杂度】
// 时间：O(N)
// 空间：O(1)
func maxBits(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	res := 0
	for max != 0 {
		res++
		max /= 10
	}
	return res
}

// radixSort 核心基数排序过程
//
// 【问题】
// 按位（从低到高）对数组进行稳定排序
//
// 【解题思路】
// 1. 每一位做一次计数排序
// 2. count数组统计当前位出现频率
// 3. 转为前缀和
// 4. 从右往左填充保证稳定性
// 5. 多轮处理每一位
//
// 【时空复杂度】
// 时间：O(N * digit)
// 空间：O(N)
func radixSort(arr []int, digit int) {
	const radix = 10
	n := len(arr)
	help := make([]int, n)

	for d := 1; d <= digit; d++ {

		count := make([]int, radix)

		// 统计当前位
		for i := 0; i < n; i++ {
			j := getDigit(arr[i], d)
			count[j]++
		}

		// 前缀和
		for i := 1; i < radix; i++ {
			count[i] += count[i-1]
		}

		// 稳定填充（从右往左）
		for i := n - 1; i >= 0; i-- {
			j := getDigit(arr[i], d)
			help[count[j]-1] = arr[i]
			count[j]--
		}

		// 回填
		for i := 0; i < n; i++ {
			arr[i] = help[i]
		}
	}
}

// getDigit 获取某一位数字
//
// 【问题】
// 获取x的第d位数字
//
// 【解题思路】
// (x / 10^(d-1)) % 10
//
// 【时空复杂度】
// 时间：O(1)
// 空间：O(1)
func getDigit(x, d int) int {
	return (x / int(math.Pow10(d-1))) % 10
}

//
// =======================
// test functions
// =======================
//

func comparator(arr []int) {
	// simple bubble sort (for test only)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func generateRandomArray(maxSize, maxValue int) []int {
	n := rand.Intn(maxSize + 1)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(maxValue + 1)
	}
	return arr
}

func copyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

func isEqual(a, b []int) bool {
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

func printArray(arr []int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	testTime := 20000
	maxSize := 100
	maxValue := 100000

	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := copyArray(arr1)

		RadixSort(arr1)
		comparator(arr2)

		if !isEqual(arr1, arr2) {
			fmt.Println("Fucked!")
			return
		}
	}

	fmt.Println("Nice!")

	arr := generateRandomArray(maxSize, maxValue)
	printArray(arr)
	RadixSort(arr)
	printArray(arr)
}
