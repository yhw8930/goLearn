package main

import (
	"math/rand"
	"time"
)

// QuickSortRecursiveUnrecursive 结构体全包，解决冲突
type QuickSortRecursiveUnrecursive struct{}

// 荷兰国旗问题
func (q *QuickSortRecursiveUnrecursive) netherlandsFlag(arr []int, L, R int) []int {
	if L > R {
		return []int{-1, -1}
	}
	if L == R {
		return []int{L, R}
	}
	less := L - 1
	more := R
	index := L
	for index < more {
		if arr[index] == arr[R] {
			index++
		} else if arr[index] < arr[R] {
			q.swap(arr, index, less+1)
			index++
			less++
		} else {
			q.swap(arr, index, more-1)
			more--
		}
	}
	q.swap(arr, more, R)
	return []int{less + 1, more}
}

func (q *QuickSortRecursiveUnrecursive) swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

// 快排递归版本
func (q *QuickSortRecursiveUnrecursive) quickSort1(arr []int) {
	if len(arr) < 2 {
		return
	}
	q.process(arr, 0, len(arr)-1)
}

func (q *QuickSortRecursiveUnrecursive) process(arr []int, L, R int) {
	if L >= R {
		return
	}
	q.swap(arr, L+rand.Intn(R-L+1), R)
	equalArea := q.netherlandsFlag(arr, L, R)
	q.process(arr, L, equalArea[0]-1)
	q.process(arr, equalArea[1]+1, R)
}

// Op 快排非递归版本需要的辅助结构体
type Op struct {
	l int
	r int
}

// 快排3.0 非递归版本  ✅ 栈用切片实现，不依赖任何包！
func (q *QuickSortRecursiveUnrecursive) quickSort2(arr []int) {
	if len(arr) < 2 {
		return
	}
	N := len(arr)
	q.swap(arr, rand.Intn(N), N-1)
	equalArea := q.netherlandsFlag(arr, 0, N-1)
	el := equalArea[0]
	er := equalArea[1]

	// 用切片模拟栈
	var stack []Op
	stack = append(stack, Op{l: 0, r: el - 1})
	stack = append(stack, Op{l: er + 1, r: N - 1})

	for len(stack) > 0 {
		// 弹出栈顶
		op := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if op.l < op.r {
			q.swap(arr, op.l+rand.Intn(op.r-op.l+1), op.r)
			equalArea = q.netherlandsFlag(arr, op.l, op.r)
			el = equalArea[0]
			er = equalArea[1]

			stack = append(stack, Op{l: op.l, r: el - 1})
			stack = append(stack, Op{l: er + 1, r: op.r})
		}
	}
}

// 生成随机数组
func (q *QuickSortRecursiveUnrecursive) generateRandomArray(maxSize, maxValue int) []int {
	size := rand.Intn(maxSize + 1)
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue+1)
	}
	return arr
}

// 拷贝数组
func (q *QuickSortRecursiveUnrecursive) copyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

// 对比数组
func (q *QuickSortRecursiveUnrecursive) isEqual(arr1, arr2 []int) bool {
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

// 打印数组
func (q *QuickSortRecursiveUnrecursive) printArray(arr []int) {
	if arr == nil {
		return
	}
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
	println()
}

// 主测试函数
func main() {
	rand.Seed(time.Now().UnixNano())
	q := &QuickSortRecursiveUnrecursive{}

	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	println("test begin")
	for i := 0; i < testTime; i++ {
		arr1 := q.generateRandomArray(maxSize, maxValue)
		arr2 := q.copyArray(arr1)

		q.quickSort1(arr1)
		q.quickSort2(arr2)

		if !q.isEqual(arr1, arr2) {
			succeed = false
			break
		}
	}
	println("test end")
	print("测试", testTime, "组是否全部通过：")
	if succeed {
		println("是")
	} else {
		println("否")
	}
}
