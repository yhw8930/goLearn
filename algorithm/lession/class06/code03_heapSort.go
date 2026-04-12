package main

import (
	"math/rand"
	"sort"
	"time"
)

// 堆排序额外空间复杂度O(1)
func heapSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// O(N) 从下往上堆化（左神原版）
	for i := len(arr) - 1; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}

	heapSize := len(arr)
	swap(arr, 0, heapSize-1)
	heapSize--

	// O(N*logN)
	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		swap(arr, 0, heapSize-1)
		heapSize--
	}
}

// heapInsert 向上调整（当前数往上飘）
func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// heapify 向下调整（当前数往下沉）左神原版
func heapify(arr []int, index int, heapSize int) {
	left := index*2 + 1 // 左孩子下标
	for left < heapSize {
		// 两个孩子中谁最大
		largest := left
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		}
		// 父和孩子谁大
		if arr[largest] <= arr[index] {
			largest = index
		}
		if largest == index {
			break
		}
		swap(arr, largest, index)
		index = largest
		left = index*2 + 1
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 对数器
func comparator(arr []int) {
	sort.Ints(arr)
}

// 生成随机数组
func generateRandomArray(maxSize, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(maxSize + 1)
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue+1)
	}
	return arr
}

// 拷贝数组
func copyArray(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

// 判断是否相等
func isEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// 打印数组
func printArray(arr []int) {
	for _, num := range arr {
		print(num, " ")
	}
	println()
}

// main 测试（完全对应Java）
func main() {
	// 测试优先级队列（小根堆）
	heap := []int{}
	push := func(x int) {
		heap = append(heap, x)
		up(heap)
	}
	pop := func() int {
		top := heap[0]
		swap(heap, 0, len(heap)-1)
		heap = heap[:len(heap)-1]
		heapify(heap, 0, len(heap))
		return top
	}

	push(6)
	push(8)
	push(0)
	push(2)
	push(9)
	push(1)

	for len(heap) > 0 {
		println(pop())
	}

	// 对数器测试
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := copyArray(arr1)
		heapSort(arr1)
		comparator(arr2)
		if !isEqual(arr1, arr2) {
			succeed = false
			break
		}
	}

	if succeed {
		println("Nice!")
	} else {
		println("Fucking fucked!")
	}

	// 演示排序
	arr := generateRandomArray(maxSize, maxValue)
	printArray(arr)
	heapSort(arr)
	printArray(arr)
}

// 小根堆向上调整（用于演示Java PriorityQueue）
func up(arr []int) {
	index := len(arr) - 1
	for arr[index] < arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}
