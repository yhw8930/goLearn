package main

import (
	"math/rand"
	"time"
)

// ReversePair 逆序对
type ReversePair struct{}

/*
这个函数的功能是计算一个数组中**“逆序对”**的数量。
在一个数组中，如果一对元素 arr[i] 和 arr[j] 满足 i < j 并且 arr[i] > arr[j]，那么我们就称 (arr[i], arr[j]) 是一个逆序对。
简单来说，就是一个数比它右边的数大，这两个数就构成一个逆序对。
举个例子，对于数组 [3, 1, 4, 2]：
3 比 1 大，3 在 1 左边 -> (3, 1) 是一个逆序对
3 比 2 大，3 在 2 左边 -> (3, 2) 是一个逆序对
4 比 2 大，4 在 2 左边 -> (4, 2) 是一个逆序对
1 比 2 小，不是逆序对
所以，这个数组的逆序对总共有 3 个。
*/
func (r *ReversePair) reverPairNumber(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return r.process(arr, 0, len(arr)-1)
}

// arr[L..R]既要排好序，也要求逆序对数量返回
// 所有merge时，产生的逆序对数量，累加，返回
// 左 排序 merge并产生逆序对数量
// 右 排序 merge并产生逆序对数量
func (r *ReversePair) process(arr []int, l, right int) int {
	if l == right {
		return 0
	}
	mid := l + ((right - l) >> 1)
	return r.process(arr, l, mid) +
		r.process(arr, mid+1, right) +
		r.merge(arr, l, mid, right)
}

func (r *ReversePair) merge(arr []int, L, m, right int) int {
	help := make([]int, right-L+1)
	i := len(help) - 1
	p1 := m
	p2 := right
	res := 0

	for p1 >= L && p2 > m {
		if arr[p1] > arr[p2] {
			res += p2 - m
			help[i] = arr[p1]
			p1--
		} else {
			help[i] = arr[p2]
			p2--
		}
		i--
	}

	for p1 >= L {
		help[i] = arr[p1]
		i--
		p1--
	}
	for p2 > m {
		help[i] = arr[p2]
		i--
		p2--
	}

	for i := 0; i < len(help); i++ {
		arr[L+i] = help[i]
	}
	return res
}

// for test
func (r *ReversePair) comparator(arr []int) int {
	ans := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				ans++
			}
		}
	}
	return ans
}

// for test
func (r *ReversePair) generateRandomArray(maxSize, maxValue int) []int {
	size := rand.Intn(maxSize + 1)
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue+1)
	}
	return arr
}

// for test
func (r *ReversePair) copyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

// for test
func (r *ReversePair) isEqual(arr1, arr2 []int) bool {
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

// for test
func (r *ReversePair) printArray(arr []int) {
	if arr == nil {
		return
	}
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
	println()
}

// for test
func main() {
	rand.Seed(time.Now().UnixNano())
	rp := &ReversePair{}

	testTime := 500000
	maxSize := 100
	maxValue := 100
	println("测试开始")

	for i := 0; i < testTime; i++ {
		arr1 := rp.generateRandomArray(maxSize, maxValue)
		arr2 := rp.copyArray(arr1)
		if rp.reverPairNumber(arr1) != rp.comparator(arr2) {
			println("Oops!")
			rp.printArray(arr1)
			rp.printArray(arr2)
			break
		}
	}
	println("测试结束")
}
