package main

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 定义结构体，把数组包起来
type GoldSplitter struct {
	arr []int
}

// NewGoldSplitter 构造函数
func NewGoldSplitter(arr []int) *GoldSplitter {
	return &GoldSplitter{arr: arr}
}

// ==================== 方法1：暴力递归 ====================
func (g *GoldSplitter) LessMoney1() int {
	if len(g.arr) == 0 {
		return 0
	}
	return g.process(g.arr, 0)
}

func (g *GoldSplitter) process(arr []int, pre int) int {
	if len(arr) == 1 {
		return pre
	}
	ans := math.MaxInt32
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			nextArr := g.copyAndMergeTwo(arr, i, j)
			ans = min(ans, g.process(nextArr, pre+arr[i]+arr[j]))
		}
	}
	return ans
}

func (g *GoldSplitter) copyAndMergeTwo(arr []int, i, j int) []int {
	ans := make([]int, len(arr)-1)
	ansi := 0
	for arri := 0; arri < len(arr); arri++ {
		if arri != i && arri != j {
			ans[ansi] = arr[arri]
			ansi++
		}
	}
	ans[ansi] = arr[i] + arr[j]
	return ans
}

// ==================== 方法2：贪心堆 ====================
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func (g *GoldSplitter) LessMoney2() int {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range g.arr {
		heap.Push(h, num)
	}
	sum := 0
	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		cur := a + b
		sum += cur
		heap.Push(h, cur)
	}
	return sum
}

func generateRandomArray(maxSize, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(maxSize + 1)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(maxValue + 1)
	}
	return arr
}

// ==================== 测试 ====================
func main() {
	testTime := 100000
	maxSize := 6
	maxValue := 1000
	fmt.Println("测试开始...")

	for i := 0; i < testTime; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		splitter := NewGoldSplitter(arr)
		ans1 := splitter.LessMoney1()
		ans2 := splitter.LessMoney2()
		if ans1 != ans2 {
			fmt.Println("Oops!")
			return
		}
	}
	fmt.Println("finish! ✅")
}
