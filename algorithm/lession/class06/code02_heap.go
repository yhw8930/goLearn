package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

// MyMaxHeap 自己实现的大根堆
type MyMaxHeap struct {
	heap     []int
	limit    int
	heapSize int
}

func NewMyMaxHeap(limit int) *MyMaxHeap {
	return &MyMaxHeap{
		heap:     make([]int, limit),
		limit:    limit,
		heapSize: 0,
	}
}

func (h *MyMaxHeap) IsEmpty() bool {
	return h.heapSize == 0
}

func (h *MyMaxHeap) IsFull() bool {
	return h.heapSize == h.limit
}

func (h *MyMaxHeap) Push(value int) {
	if h.heapSize == h.limit {
		panic("heap is full")
	}
	h.heap[h.heapSize] = value
	h.heapInsert(h.heapSize)
	h.heapSize++
}

// Pop 弹出最大值
func (h *MyMaxHeap) Pop() int {
	ans := h.heap[0]
	h.swap(0, h.heapSize-1)
	h.heapSize--
	h.heapify(0, h.heapSize)
	return ans
}

// heapInsert 向上调整
func (h *MyMaxHeap) heapInsert(index int) {
	for h.heap[index] > h.heap[(index-1)/2] {
		h.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// heapify 向下调整
func (h *MyMaxHeap) heapify(index int, heapSize int) {
	left := index*2 + 1
	for left < heapSize {
		largest := left
		if left+1 < heapSize && h.heap[left+1] > h.heap[left] {
			largest = left + 1
		}
		if h.heap[largest] <= h.heap[index] {
			largest = index
		}
		if largest == index {
			break
		}
		h.swap(index, largest)
		index = largest
		left = index*2 + 1
	}
}

func (h *MyMaxHeap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

// RightMaxHeap 暴力堆（用于测试对比）
type RightMaxHeap struct {
	arr   []int
	limit int
	size  int
}

func NewRightMaxHeap(limit int) *RightMaxHeap {
	return &RightMaxHeap{
		arr:   make([]int, limit),
		limit: limit,
		size:  0,
	}
}

func (r *RightMaxHeap) IsEmpty() bool {
	return r.size == 0
}

func (r *RightMaxHeap) IsFull() bool {
	return r.size == r.limit
}

func (r *RightMaxHeap) Push(value int) {
	if r.size == r.limit {
		panic("heap is full")
	}
	r.arr[r.size] = value
	r.size++
}

func (r *RightMaxHeap) Pop() int {
	maxIndex := 0
	for i := 1; i < r.size; i++ {
		if r.arr[i] > r.arr[maxIndex] {
			maxIndex = i
		}
	}
	ans := r.arr[maxIndex]
	r.arr[maxIndex] = r.arr[r.size-1]
	r.size--
	return ans
}

func main() {
	// 第一部分：测试 Go 优先级队列（大根堆）
	pq := &IntHeap{}
	heap.Init(pq)

	pq.Push(5)
	pq.Push(5)
	pq.Push(5)
	pq.Push(3)
	fmt.Println(pq.Peek()) // 看堆顶 5

	pq.Push(7)
	pq.Push(0)
	pq.Push(7)
	pq.Push(0)
	pq.Push(7)
	pq.Push(0)
	fmt.Println(pq.Peek()) // 看堆顶 7

	for pq.Len() > 0 {
		fmt.Println(pq.Pop())
	}

	// 第二部分：对数器测试
	rand.Seed(time.Now().UnixNano())
	value := 1000
	limit := 100
	testTimes := 1000000

	fmt.Println("测试开始...")
	for i := 0; i < testTimes; i++ {
		curLimit := rand.Intn(limit) + 1
		my := NewMyMaxHeap(curLimit)
		test := NewRightMaxHeap(curLimit)
		curOpTimes := rand.Intn(limit)

		for j := 0; j < curOpTimes; j++ {
			if my.IsEmpty() != test.IsEmpty() {
				fmt.Println("Oops! 空状态错误")
				return
			}
			if my.IsFull() != test.IsFull() {
				fmt.Println("Oops! 满状态错误")
				return
			}

			if my.IsEmpty() {
				curValue := rand.Intn(value)
				my.Push(curValue)
				test.Push(curValue)
			} else if my.IsFull() {
				if my.Pop() != test.Pop() {
					fmt.Println("Oops! Pop 值错误")
					return
				}
			} else {
				if rand.Float64() < 0.5 {
					curValue := rand.Intn(value)
					my.Push(curValue)
					test.Push(curValue)
				} else {
					if my.Pop() != test.Pop() {
						fmt.Println("Oops! Pop 值错误")
						return
					}
				}
			}
		}
	}
	fmt.Println("finish!")
}

// 给 IntHeap 加一个 Peek 方法（看堆顶）
func (h *IntHeap) Peek() int {
	if h.Len() == 0 {
		return 0
	}
	return (*h)[0]
}
