package main

import (
	"math/rand"
	"time"
)

// 题目：实现数组划分，并基于划分实现快速排序的多个版本。
// 划分目标是围绕基准值，把数组分成小于区、等于区和大于区。
// 核心思路：荷兰国旗过程用 less、more、index 三个边界维护三个区域。
// 快速排序每次随机或固定选基准，划分后只递归处理小于区和大于区，等于区不用再排。
// 时间复杂度：随机快排期望 O(NlogN)，最坏 O(N^2)。
// 空间复杂度：期望 O(logN)，来自递归深度。

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
// partition 是单边划分版本。
// 它只把小于等于基准的数放到左侧，返回基准最终位置。
// 时间复杂度：O(R-L+1)。
// 空间复杂度：O(1)。
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
// netherlandsFlag 是荷兰国旗划分。
// 它一次划出小于区、等于区和大于区，特别适合处理大量重复值。
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

// quickSort1 基于普通 partition 递归排序。
// 每次只确定一个基准位置，再递归处理左右两侧。
// 时间复杂度：固定基准快排平均 O(NlogN)，最坏 O(N^2)。
// 空间复杂度：平均 O(logN)，最坏 O(N)。
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

// quickSort2 基于荷兰国旗划分递归排序。
// 等于基准的一整段都不用再处理，重复值较多时更高效。
// 时间复杂度：平均 O(NlogN)，最坏 O(N^2)，重复值多时通常更优。
// 空间复杂度：平均 O(logN)，最坏 O(N)。
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

// quickSort3 是随机快速排序。
// 排序前随机选择基准并交换到末尾，降低遇到最坏划分的概率。
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
