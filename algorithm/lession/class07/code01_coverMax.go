package main

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Line struct {
	start int
	end   int
}

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
