package main

import (
	"math"
	"math/rand"
	"time"
)

type MergeSort struct{}

// 递归方法实现
func (m *MergeSort) mergeSort1(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	m.process(arr, 0, len(arr)-1)
}

// 请把arr[L..R]排有序
// l...r N
// T(N) = 2 * T(N / 2) + O(N)
// O(N * logN)
func (m *MergeSort) process(arr []int, L, R int) {
	if L == R { // base case
		return
	}
	mid := L + ((R - L) >> 1)
	m.process(arr, L, mid)
	m.process(arr, mid+1, R)
	m.merge(arr, L, mid, R)
}

func (m *MergeSort) merge(arr []int, L, M, R int) {
	help := make([]int, R-L+1)
	i := 0
	p1 := L
	p2 := M + 1
	for p1 <= M && p2 <= R {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}
	// 要么p1越界了，要么p2越界了
	for p1 <= M {
		help[i] = arr[p1]
		i++
		p1++
	}
	for p2 <= R {
		help[i] = arr[p2]
		i++
		p2++
	}
	for i := 0; i < len(help); i++ {
		arr[L+i] = help[i]
	}
}

// 非递归方法实现
func (m *MergeSort) mergeSort2(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	N := len(arr)
	// 步长
	mergeSize := 1
	for mergeSize < N { // log N
		// 当前左组的，第一个位置
		L := 0
		for L < N {
			if mergeSize >= N-L {
				break
			}
			M := L + mergeSize - 1
			R := M + int(math.Min(float64(mergeSize), float64(N-M-1)))
			m.merge(arr, L, M, R)
			L = R + 1
		}
		// 防止溢出
		if mergeSize > N/2 {
			break
		}
		mergeSize <<= 1
	}
}

func (m *MergeSort) generateRandomArray(maxSize, maxValue int) []int {
	size := rand.Intn(maxSize + 1)
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue+1)
	}
	return arr
}

func (m *MergeSort) copyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

func (m *MergeSort) isEqual(arr1, arr2 []int) bool {
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

func (m *MergeSort) printArray(arr []int) {
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
	ms := &MergeSort{} // 实例化

	testTime := 500000
	maxSize := 100
	maxValue := 100
	println("测试开始")

	for i := 0; i < testTime; i++ {
		arr1 := ms.generateRandomArray(maxSize, maxValue)
		arr2 := ms.copyArray(arr1)
		ms.mergeSort1(arr1)
		ms.mergeSort2(arr2)
		if !ms.isEqual(arr1, arr2) {
			println("出错了！")
			ms.printArray(arr1)
			ms.printArray(arr2)
			break
		}
	}
	println("测试结束")
}
