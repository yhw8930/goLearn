package main

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 题目：一根金条要切成指定长度数组，每次切割代价等于当前被切金条长度，求最小总代价。
// 正向切割难以贪心选择，但可以反过来看成把小段金条合并成大金条。
// 核心思路：每次合并两块最小的金条，产生的代价最小，并把合并后的新金条继续放回集合。
// 这就是哈夫曼编码的贪心模型，用小根堆能高效取出当前最小的两块。
// 时间复杂度：暴力为指数级，堆贪心为 O(NlogN)。
// 空间复杂度：O(N)。

// 定义结构体，把数组包起来
type GoldSplitter struct {
	arr []int
}

// NewGoldSplitter 构造函数
func NewGoldSplitter(arr []int) *GoldSplitter {
	return &GoldSplitter{arr: arr}
}

// ==================== 方法1：暴力递归 ====================
// LessMoney1 是暴力递归版本。
// 它枚举每次合并哪两块金条，尝试所有合并顺序。
// 时间复杂度：指数级，枚举所有合并顺序。
// 空间复杂度：O(N)，递归深度和临时数组。
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

// LessMoney2 是小根堆贪心版本。
// 每次取出当前最小的两块合并，累计代价后再放回堆。
// 时间复杂度：O(NlogN)。
// 空间复杂度：O(N)，小根堆保存当前金条。
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
