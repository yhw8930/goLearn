package main

import (
	"container/heap"
	"fmt"
)

// IPOProgram 项目结构体（利润 p，成本 c）
type IPOProgram struct {
	p int // 利润 profit
	c int // 成本 capital
}

// IPO 结构体（封装所有数据，变成结构体函数）
type IPO struct {
	profits []int
	capital []int
}

// NewIPO 构造函数
func NewIPO(profits, capital []int) *IPO {
	return &IPO{
		profits: profits,
		capital: capital,
	}
}

// -------------------- 小根堆：按成本升序 --------------------
type MinCostHeap []IPOProgram

func (h MinCostHeap) Len() int           { return len(h) }
func (h MinCostHeap) Less(i, j int) bool { return h[i].c < h[j].c } // 成本小优先
func (h MinCostHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinCostHeap) Push(x interface{}) {
	*h = append(*h, x.(IPOProgram))
}

func (h *MinCostHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// -------------------- 大根堆：按利润降序 --------------------
type MaxProfitHeap []IPOProgram

func (h MaxProfitHeap) Len() int           { return len(h) }
func (h MaxProfitHeap) Less(i, j int) bool { return h[i].p > h[j].p } // 利润大优先
func (h MaxProfitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxProfitHeap) Push(x interface{}) {
	*h = append(*h, x.(IPOProgram))
}

func (h *MaxProfitHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// -------------------- 核心算法（结构体函数） --------------------
// FindMaximizedCapital 最多做 K 个项目，初始资金 W，返回最大资金
func (ipo *IPO) FindMaximizedCapital(K int, W int) int {
	// 小根堆：按成本排序
	minCost := &MinCostHeap{}
	heap.Init(minCost)

	// 大根堆：按利润排序
	maxProfit := &MaxProfitHeap{}
	heap.Init(maxProfit)

	// 所有项目先加入小根堆
	for i := 0; i < len(ipo.profits); i++ {
		heap.Push(minCost, IPOProgram{p: ipo.profits[i], c: ipo.capital[i]})
	}

	// 做 K 个项目
	for i := 0; i < K; i++ {
		// 把能做的项目（成本 <= W）全部移入大根堆
		for minCost.Len() > 0 && (*minCost)[0].c <= W {
			heap.Push(maxProfit, heap.Pop(minCost))
		}

		// 没有项目可做，直接返回
		if maxProfit.Len() == 0 {
			return W
		}

		// 选利润最大的做
		W += heap.Pop(maxProfit).(IPOProgram).p
	}

	return W
}

// -------------------- 测试 --------------------
func main() {
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}
	ipo := NewIPO(profits, capital)
	res := ipo.FindMaximizedCapital(2, 0)
	fmt.Println("最大资金：", res) // 输出 4 ✅
}
