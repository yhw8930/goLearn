package main

import (
	"math/rand"
	"time"
)

// PartitionAndQuickSort 结构体，全包函数，解决重名冲突
type PartitionAndQuickSort struct{}

func (p *PartitionAndQuickSort) swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

// arr[L..R]上，以arr[R]位置的数做划分值
// <= X > X
// <= X X
func (p *PartitionAndQuickSort) partition(arr []int, L, R int) int {
	if L > R {
		return -1
	}
	if L == R {
		return L
	}
	lessEqual := L - 1
	index := L
	for index < R {
		if arr[index] <= arr[R] {
			p.swap(arr, index, lessEqual+1)
			lessEqual++
		}
		index++
	}
	p.swap(arr, lessEqual+1, R)
	return lessEqual + 1
}

// arr[L...R] 玩荷兰国旗问题的划分，以arr[R]做划分值
// <arr[R] ==arr[R] > arr[R]
func (p *PartitionAndQuickSort) netherlandsFlag(arr []int, L, R int) []int {
	if L > R { // L...R L>R
		return []int{-1, -1}
	}
	if L == R {
		return []int{L, R}
	}
	less := L - 1 // < 区 右边界
	more := R     // > 区 左边界
	index := L
	for index < more { // 当前位置，不能和 >区的左边界撞上
		if arr[index] == arr[R] {
			index++
		} else if arr[index] < arr[R] {
			p.swap(arr, index, less+1)
			less++
			index++
		} else { // >
			p.swap(arr, index, more-1)
			more--
		}
	}
	p.swap(arr, more, R) // <[R]   =[R]   >[R]
	return []int{less + 1, more}
}

func (p *PartitionAndQuickSort) quickSort1(arr []int) {
	if len(arr) < 2 {
		return
	}
	p.process1(arr, 0, len(arr)-1)
}

func (p *PartitionAndQuickSort) process1(arr []int, L, R int) {
	if L >= R {
		return
	}
	// L..R partition arr[R] [ <=arr[R] arr[R] >arr[R] ]
	M := p.partition(arr, L, R)
	p.process1(arr, L, M-1)
	p.process1(arr, M+1, R)
}

func (p *PartitionAndQuickSort) quickSort2(arr []int) {
	if len(arr) < 2 {
		return
	}
	p.process2(arr, 0, len(arr)-1)
}

// arr[L...R] 排有序，快排2.0方式
func (p *PartitionAndQuickSort) process2(arr []int, L, R int) {
	if L >= R {
		return
	}
	// [ equalArea[0]  ,  equalArea[0]]
	equalArea := p.netherlandsFlag(arr, L, R)
	p.process2(arr, L, equalArea[0]-1)
	p.process2(arr, equalArea[1]+1, R)
}

func (p *PartitionAndQuickSort) quickSort3(arr []int) {
	if len(arr) < 2 {
		return
	}
	p.process3(arr, 0, len(arr)-1)
}

func (p *PartitionAndQuickSort) process3(arr []int, L, R int) {
	if L >= R {
		return
	}
	p.swap(arr, L+rand.Intn(R-L+1), R)
	equalArea := p.netherlandsFlag(arr, L, R)
	p.process3(arr, L, equalArea[0]-1)
	p.process3(arr, equalArea[1]+1, R)
}

// for test
func (p *PartitionAndQuickSort) generateRandomArray(maxSize, maxValue int) []int {
	size := rand.Intn(maxSize + 1)
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue+1)
	}
	return arr
}

// for test
func (p *PartitionAndQuickSort) copyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

// for test
func (p *PartitionAndQuickSort) isEqual(arr1, arr2 []int) bool {
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
func (p *PartitionAndQuickSort) printArray(arr []int) {
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
	p := &PartitionAndQuickSort{}

	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		arr1 := p.generateRandomArray(maxSize, maxValue)
		arr2 := p.copyArray(arr1)
		arr3 := p.copyArray(arr1)

		p.quickSort1(arr1)
		p.quickSort2(arr2)
		p.quickSort3(arr3)

		if !p.isEqual(arr1, arr2) || !p.isEqual(arr2, arr3) {
			succeed = false
			break
		}
	}

	if succeed {
		println("Nice!")
	} else {
		println("Oops!")
	}
}
