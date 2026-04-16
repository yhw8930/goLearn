package main

import (
	"fmt"
)

type HeapGreater[T any] struct {
	heap     []T
	indexMap map[any]int
	heapSize int
	comp     func(T, T) int
}

func NewHeapGreater[T any](comp func(T, T) int) *HeapGreater[T] {
	return &HeapGreater[T]{
		heap:     make([]T, 0),
		indexMap: make(map[any]int),
		heapSize: 0,
		comp:     comp,
	}
}

func (h *HeapGreater[T]) IsEmpty() bool {
	return h.heapSize == 0
}

func (h *HeapGreater[T]) Size() int {
	return h.heapSize
}

func (h *HeapGreater[T]) Contains(obj T) bool {
	_, ok := h.indexMap[obj]
	return ok
}

func (h *HeapGreater[T]) Peek() T {
	return h.heap[0]
}

func (h *HeapGreater[T]) Push(obj T) {
	h.heap = append(h.heap, obj)
	h.indexMap[obj] = h.heapSize
	h.heapInsert(h.heapSize)
	h.heapSize++
}

func (h *HeapGreater[T]) Pop() T {
	ans := h.heap[0]
	h.swap(0, h.heapSize-1)
	delete(h.indexMap, ans)
	h.heap = h.heap[:h.heapSize-1]
	h.heapSize--
	h.heapify(0)
	return ans
}

func (h *HeapGreater[T]) Remove(obj T) {
	replace := h.heap[h.heapSize-1]
	index := h.indexMap[obj]
	delete(h.indexMap, obj)
	h.heap = h.heap[:h.heapSize-1]
	h.heapSize--

	if any(obj) != any(replace) {
		h.heap[index] = replace
		h.indexMap[replace] = index
		h.Resign(replace)
	}
}

func (h *HeapGreater[T]) Resign(obj T) {
	index := h.indexMap[obj]
	h.heapInsert(index)
	h.heapify(index)
}

func (h *HeapGreater[T]) GetAllElements() []T {
	ans := make([]T, len(h.heap))
	copy(ans, h.heap)
	return ans
}

func (h *HeapGreater[T]) heapInsert(index int) {
	for h.comp(h.heap[index], h.heap[(index-1)/2]) < 0 {
		h.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func (h *HeapGreater[T]) heapify(index int) {
	left := index*2 + 1
	for left < h.heapSize {
		best := left
		if left+1 < h.heapSize && h.comp(h.heap[left+1], h.heap[left]) < 0 {
			best = left + 1
		}
		if h.comp(h.heap[best], h.heap[index]) >= 0 {
			break
		}
		h.swap(best, index)
		index = best
		left = index*2 + 1
	}
}

func (h *HeapGreater[T]) swap(i, j int) {
	o1 := h.heap[i]
	o2 := h.heap[j]
	h.heap[i] = o2
	h.heap[j] = o1
	h.indexMap[o2] = i
	h.indexMap[o1] = j
}

// Customer 顾客结构体
type Customer struct {
	id        int
	buy       int
	enterTime int
}

// CandidateComparator 候选区比较器：购买数降序，时间升序
func CandidateComparator(a, b *Customer) int {
	if a.buy != b.buy {
		return b.buy - a.buy // 降序
	}
	return a.enterTime - b.enterTime // 升序
}

// DaddyComparator 得奖区比较器：购买数升序，时间升序
func DaddyComparator(a, b *Customer) int {
	if a.buy != b.buy {
		return a.buy - b.buy // 升序
	}
	return a.enterTime - b.enterTime // 升序
}

// WhosYourDaddy 核心管理类
type WhosYourDaddy struct {
	customers  map[int]*Customer
	candHeap   *HeapGreater[*Customer]
	daddyHeap  *HeapGreater[*Customer]
	daddyLimit int
}

func NewWhosYourDaddy(limit int) *WhosYourDaddy {
	return &WhosYourDaddy{
		customers:  make(map[int]*Customer),
		candHeap:   NewHeapGreater(CandidateComparator),
		daddyHeap:  NewHeapGreater(DaddyComparator),
		daddyLimit: limit,
	}
}

// operate 处理购买或退货事件
func (w *WhosYourDaddy) operate(time, id int, buyOrRefund bool) {
	// 无效退货：用户不存在且要退货
	if !buyOrRefund && w.customers[id] == nil {
		return
	}

	// 用户不存在则创建
	if w.customers[id] == nil {
		w.customers[id] = &Customer{id: id, buy: 0, enterTime: 0}
	}

	c := w.customers[id]
	if buyOrRefund {
		c.buy++
	} else {
		c.buy--
	}

	// 购买数为0时删除用户
	if c.buy == 0 {
		delete(w.customers, id)
	}

	// 用户不在任何堆中
	if !w.candHeap.Contains(c) && !w.daddyHeap.Contains(c) {
		if w.daddyHeap.Size() < w.daddyLimit {
			c.enterTime = time
			w.daddyHeap.Push(c)
		} else {
			c.enterTime = time
			w.candHeap.Push(c)
		}
	} else if w.candHeap.Contains(c) { // 在候选区
		if c.buy == 0 {
			w.candHeap.Remove(c)
		} else {
			w.candHeap.Resign(c)
		}
	} else { // 在得奖区
		if c.buy == 0 {
			w.daddyHeap.Remove(c)
		} else {
			w.daddyHeap.Resign(c)
		}
	}

	w.daddyMove(time)
}

// getDaddies 获取当前得奖区用户ID列表
func (w *WhosYourDaddy) getDaddies() []int {
	customers := w.daddyHeap.GetAllElements()
	ans := make([]int, len(customers))
	for i, c := range customers {
		ans[i] = c.id
	}
	return ans
}

// daddyMove 调整得奖区和候选区
func (w *WhosYourDaddy) daddyMove(time int) {
	if w.candHeap.IsEmpty() {
		return
	}

	if w.daddyHeap.Size() < w.daddyLimit {
		p := w.candHeap.Pop()
		p.enterTime = time
		w.daddyHeap.Push(p)
	} else {
		if w.candHeap.Peek().buy > w.daddyHeap.Peek().buy {
			oldDaddy := w.daddyHeap.Pop()
			newDaddy := w.candHeap.Pop()
			oldDaddy.enterTime = time
			newDaddy.enterTime = time
			w.daddyHeap.Push(newDaddy)
			w.candHeap.Push(oldDaddy)
		}
	}
}

// topK 主函数：处理事件序列，返回每一步的得奖区用户
func topK(arr []int, op []bool, k int) [][]int {
	ans := make([][]int, len(arr))
	whoDaddies := NewWhosYourDaddy(k)

	for i := 0; i < len(arr); i++ {
		whoDaddies.operate(i, arr[i], op[i])
		ans[i] = whoDaddies.getDaddies()
	}
	return ans
}

func main() {
	// 测试代码
	arr := []int{1, 2, 3, 4, 5}
	op := []bool{true, true, false, true, true}
	k := 2

	result := topK(arr, op, k)
	for i, daddies := range result {
		fmt.Printf("步骤%d: 得奖用户 %v\n", i, daddies)
	}
}
