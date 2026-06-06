package main

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// 题目：给定多条线段，每条线段有起点和终点，求最多有多少条线段在同一区域重合。
// 线段重合可以理解为存在某个点，被最多线段同时覆盖。
// 核心思路：先按起点排序，从左到右扫描线段，把当前仍未结束的线段终点放入小根堆。
// 扫描到新线段起点时，弹出所有终点小于等于该起点的线段，堆大小就是当前重合数。
// 时间复杂度：O(NlogN)。
// 空间复杂度：O(N)。

type Line struct {
	start int
	end   int
}

// maxCover1 是暴力验证方法。
// 它枚举可能的中间点，统计每个点被多少线段覆盖。
// 时间复杂度：O(N*R)，R 为枚举点范围，适合作为暴力验证。
// 空间复杂度：O(1)。
func maxCover1(lines [][]int) int {
	minVal := math.MaxInt32
	maxVal := math.MinInt32
	for i := 0; i < len(lines); i++ {
		minVal = min(minVal, lines[i][0])
		maxVal = max(maxVal, lines[i][1])
	}
	cover := 0
	for p := float64(minVal) + 0.5; p < float64(maxVal); p += 1 {
		cur := 0
		for i := 0; i < len(lines); i++ {
			if float64(lines[i][0]) < p && float64(lines[i][1]) > p {
				cur++
			}
		}
		cover = max(cover, cur)
	}
	return cover
}

// maxCover2 是堆优化方法。
// 它按起点扫描线段，用小根堆维护当前还没结束的线段终点。
// 时间复杂度：O(NlogN)。
// 空间复杂度：O(N)，小根堆保存未结束线段。
func maxCover2(m [][]int) int {
	lines := make([]Line, len(m))
	for i := 0; i < len(m); i++ {
		lines[i] = Line{m[i][0], m[i][1]}
	}
	sort.Slice(lines, func(i, j int) bool {
		return lines[i].start < lines[j].start
	})

	h := &IntHeap{}
	heap.Init(h)
	maxVal := 0
	for i := 0; i < len(lines); i++ {
		for h.Len() > 0 && h.Peek() <= lines[i].start {
			heap.Pop(h)
		}
		heap.Push(h, lines[i].end)
		maxVal = max(maxVal, h.Len())
	}
	return maxVal
}

func generateLines(N, L, R int) [][]int {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(N) + 1
	ans := make([][]int, size)
	for i := 0; i < size; i++ {
		ans[i] = make([]int, 2) // 初始化内层切片
		a := L + rand.Intn(R-L+1)
		b := L + rand.Intn(R-L+1)
		if a == b {
			b = a + 1
		}
		ans[i][0] = min(a, b)
		ans[i][1] = max(a, b)
	}
	return ans
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() int {
	return (*h)[0]
}

func main() {
	l1 := Line{4, 9}
	l2 := Line{1, 4}
	l3 := Line{7, 15}
	l4 := Line{2, 4}
	l5 := Line{4, 6}
	l6 := Line{3, 7}

	h := &LineHeap{}
	heap.Init(h)
	heap.Push(h, l1)
	heap.Push(h, l2)
	heap.Push(h, l3)
	heap.Push(h, l4)
	heap.Push(h, l5)
	heap.Push(h, l6)

	for h.Len() > 0 {
		cur := heap.Pop(h).(Line)
		fmt.Println(cur.start, ",", cur.end)
	}

	fmt.Println("test begin")
	N := 100
	L := 0
	R := 200
	testTimes := 200000
	for i := 0; i < testTimes; i++ {
		lines := generateLines(N, L, R)
		ans1 := maxCover1(lines)
		ans2 := maxCover2(lines)
		if ans1 != ans2 {
			fmt.Println("Oops!")
		}
	}
	fmt.Println("test end")
}

type LineHeap []Line

func (h LineHeap) Len() int           { return len(h) }
func (h LineHeap) Less(i, j int) bool { return h[i].start < h[j].start }
func (h LineHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *LineHeap) Push(x interface{}) {
	*h = append(*h, x.(Line))
}

func (h *LineHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
