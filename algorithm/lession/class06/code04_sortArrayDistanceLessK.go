package main

import (
	"container/heap"
	"math/rand"
	"sort"
	"time"
)

// ******************************************************************
// 堆排序已经有 IntHeap 了，这里不再定义！不会冲突！
// ******************************************************************

func sortedArrDistanceLessK(arr []int, k int) {
	if k == 0 {
		return
	}
	pq := &IntHeap{}
	heap.Init(pq)
	n := len(arr)
	index := 0

	for ; index <= min(n-1, k-1); index++ {
		heap.Push(pq, arr[index])
	}

	i := 0
	for ; index < n; i, index = i+1, index+1 {
		heap.Push(pq, arr[index])
		arr[i] = heap.Pop(pq).(int)
	}

	for pq.Len() > 0 {
		arr[i] = heap.Pop(pq).(int)
		i++
	}
}

// =================================================================
// 全部加前缀 lessK_ ，和堆排序完全不冲突
// =================================================================
func lessK_comparator(arr []int) {
	sort.Ints(arr)
}

func lessK_randomArrayNoMoveMoreK(maxSize, maxValue, K int) []int {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(maxSize + 1)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue+1)
	}
	sort.Ints(arr)

	isSwap := make([]bool, size)
	for i := 0; i < size; i++ {
		j := min(i+rand.Intn(K+1), size-1)
		if !isSwap[i] && !isSwap[j] {
			isSwap[i] = true
			isSwap[j] = true
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return arr
}

func lessK_copyArray(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

func lessK_isEqual(arr1, arr2 []int) bool {
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

func lessK_printArray(arr []int) {
	for _, num := range arr {
		print(num, " ")
	}
	println()
}

// =================================================================
// main 也改名 lessK_main
// =================================================================
func lessK_main() {
	println("test begin")
	rand.Seed(time.Now().UnixNano())
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		k := rand.Intn(maxSize) + 1
		arr := lessK_randomArrayNoMoveMoreK(maxSize, maxValue, k)
		arr1 := lessK_copyArray(arr)
		arr2 := lessK_copyArray(arr)

		sortedArrDistanceLessK(arr1, k)
		lessK_comparator(arr2)

		if !lessK_isEqual(arr1, arr2) {
			succeed = false
			println("K :", k)
			lessK_printArray(arr)
			lessK_printArray(arr1)
			lessK_printArray(arr2)
			break
		}
	}

	if succeed {
		println("Nice!")
	} else {
		println("Fucking fucked!")
	}
}
