package main

import (
	"math/rand"
	"time"
)

type SmallSum struct{}

//对于一个数组中的每个数，它左边所有比它小的数的总和，就是这个数的小和。整个数组的小和就是所有数的小和之和。
//
//举个例子，对于数组 [1, 3, 4, 2, 5]：
//
//1 的左边没有数，小和是 0
//3 的左边有 1 比它小，小和是 1
//4 的左边有 1, 3 比它小，小和是 1 + 3 = 4
//2 的左边有 1 比它小，小和是 1
//5 的左边有 1, 3, 4, 2 比它小，小和是 1 + 3 + 4 + 2 = 10
//所以整个数组的小和就是 0 + 1 + 4 + 1 + 10 = 16。

// 算法核心思想是：在对数组进行归-并排序的过程中，计算产生的小和, 将时间复杂度优化到了 O(N * logN)。
// 分解 (Divide)： process 函数递归地将数组分成两半，直到每个子数组只包含一个元素。这和归并排序的分解步骤完全一样。
// 合并与计算 (Merge & Conquer)： 这是算法的关键。当合并两个已经排好序的子数组（左组和右组）时，我们可以高效地计算出“跨左右两组”的小和。
// 在 merge 函数中，当比较左组的数 arr[p1] 和右组的数 arr[p2] 时：
// 如果 arr[p1] < arr[p2]，这意味着右组中从 arr[p2] 到结尾的所有数都比 arr[p1] 大。因此，arr[p1] 会为这些数（共 r - p2 + 1 个）都贡献一次小和。所以，在这一步产生的小和就是 arr[p1] * (r - p2 + 1)。
// 如果 arr[p1] >= arr[p2]，则 arr[p1] 不会为 arr[p2] 产生小和（因为它不比 arr[p2] 小），所以小和贡献为 0。
// 通过在归并排序的 merge 过程中增加这个计算步骤，smallSum 可以在排序的同时，高效地计算出整个数组的小和。
func (s *SmallSum) smallSum(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return s.process(arr, 0, len(arr)-1)
}

// arr[L..R]既要排好序，也要求小和返回
// 所有merge时，产生的小和，累加
// 左 排序   merge
// 右 排序  merge
// merge
func (s *SmallSum) process(arr []int, l int, r int) int {
	if l == r {
		return 0
	}
	// l < r
	mid := l + ((r - l) >> 1)
	return s.process(arr, l, mid) +
		s.process(arr, mid+1, r) +
		s.merge(arr, l, mid, r)
}

func (s *SmallSum) merge(arr []int, L int, m int, r int) int {
	help := make([]int, r-L+1)
	i := 0
	p1 := L
	p2 := m + 1
	res := 0

	for p1 <= m && p2 <= r {
		if arr[p1] < arr[p2] {
			res += (r - p2 + 1) * arr[p1]
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}

	for p1 <= m {
		help[i] = arr[p1]
		i++
		p1++
	}
	for p2 <= r {
		help[i] = arr[p2]
		i++
		p2++
	}

	for i := 0; i < len(help); i++ {
		arr[L+i] = help[i]
	}
	return res
}

// for test
func (s *SmallSum) comparator(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	res := 0
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				res += arr[j]
			}
		}
	}
	return res
}

// ================= 对数器全部放进结构体 =================
func (s *SmallSum) generateRandomArray(maxSize int, maxValue int) []int {
	size := rand.Intn(maxSize + 1)
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue+1)
	}
	return arr
}

func (s *SmallSum) copyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

func (s *SmallSum) isEqual(arr1 []int, arr2 []int) bool {
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

func (s *SmallSum) printArray(arr []int) {
	if arr == nil {
		return
	}
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
	println()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ss := &SmallSum{} // 实例化

	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		arr1 := ss.generateRandomArray(maxSize, maxValue)
		arr2 := ss.copyArray(arr1)
		if ss.smallSum(arr1) != ss.comparator(arr2) {
			succeed = false
			ss.printArray(arr1)
			ss.printArray(arr2)
			break
		}
	}

	if succeed {
		println("Nice!")
	} else {
		println("Fucking fucked!")
	}
}
