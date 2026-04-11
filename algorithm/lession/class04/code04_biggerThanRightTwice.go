package main

import (
	"math/rand"
	"time"
)

// BiggerThanRightTwice 大于右边两倍
type BiggerThanRightTwice struct{}

/*
这个函数的目标是解决一个特定的计数问题：在一个数组中，对于每个数，找出它右边有多少个数乘以2后仍然比它小，然后求出总共有多少个这样的数对。
换句话说，我们要计算满足 i < j 且 arr[i] > arr[j] * 2 的数对 (i, j) 的总数量。
举个例子，对于数组 [6, 3, 2, 1, 0]：
对于 6 (i=0)，右边的 3*2=6 (不满足), 2*2=4 (满足), 1*2=2 (满足), 0*2=0 (满足)。所以有 3 个数对 (6,2), (6,1), (6,0)。
对于 3 (i=1)，右边的 2*2=4 (不满足), 1*2=2 (满足), 0*2=0 (满足)。所以有 2 个数对 (3,1), (3,0)。
对于 2 (i=2)，右边的 1*2=2 (不满足), 0*2=0 (满足)。所以有 1 个数对 (2,0)。
...
总数量是 3 + 2 + 1 = 6
*/
func (b *BiggerThanRightTwice) biggerTwice(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return b.process(arr, 0, len(arr)-1)
}

func (b *BiggerThanRightTwice) process(arr []int, l int, r int) int {
	if l == r {
		return 0
	}
	// l < r
	mid := l + ((r - l) >> 1)
	return b.process(arr, l, mid) + b.process(arr, mid+1, r) + b.merge(arr, l, mid, r)
}

func (b *BiggerThanRightTwice) merge(arr []int, L int, m int, r int) int {
	// [L....M]   [M+1....R]

	ans := 0
	// 目前囊括进来的数，是从[M+1, windowR)
	windowR := m + 1
	for i := L; i <= m; i++ {
		for windowR <= r && arr[i] > (arr[windowR]*2) {
			windowR++
		}
		ans += windowR - m - 1
	}

	help := make([]int, r-L+1)
	i := 0
	p1 := L
	p2 := m + 1
	for p1 <= m && p2 <= r {
		if arr[p1] <= arr[p2] {
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
	return ans
}

// for test
func (b *BiggerThanRightTwice) comparator(arr []int) int {
	ans := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > (arr[j] << 1) {
				ans++
			}
		}
	}
	return ans
}

// for test
func (b *BiggerThanRightTwice) generateRandomArray(maxSize int, maxValue int) []int {
	size := rand.Intn(maxSize + 1)
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue+1)
	}
	return arr
}

// for test
func (b *BiggerThanRightTwice) copyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

// for test
func (b *BiggerThanRightTwice) isEqual(arr1 []int, arr2 []int) bool {
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
func (b *BiggerThanRightTwice) printArray(arr []int) {
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
	bt := &BiggerThanRightTwice{}

	testTime := 500000
	maxSize := 100
	maxValue := 100
	println("测试开始")
	for i := 0; i < testTime; i++ {
		arr1 := bt.generateRandomArray(maxSize, maxValue)
		arr2 := bt.copyArray(arr1)
		if bt.biggerTwice(arr1) != bt.comparator(arr2) {
			println("Oops!")
			bt.printArray(arr1)
			bt.printArray(arr2)
			break
		}
	}
	println("测试结束")
}
