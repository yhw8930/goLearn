package main

// 题目：实现一个加强堆，支持判断元素是否存在、删除任意元素、元素优先级变化后重排。
// 普通堆只能高效处理堆顶，无法快速定位某个任意元素。
// 核心思路：额外维护 indexMap，记录每个元素当前在堆数组中的下标。
// 当元素变化或删除时，可以 O(logN) 从对应下标向上/向下调整，而不用线性查找。
// 时间复杂度：Push/Pop/Remove/Resign 为 O(logN)，Contains 为 O(1)。
// 空间复杂度：O(N)。

//type HeapGreater[T any] struct {
//	heap     []T
//	indexMap map[any]int
//	heapSize int
//	comp     func(T, T) int
//}
//
//func NewHeapGreater[T any](comp func(T, T) int) *HeapGreater[T] {
//	return &HeapGreater[T]{
//		heap:     make([]T, 0),
//		indexMap: make(map[any]int),
//		heapSize: 0,
//		comp:     comp,
//	}
//}
//
//func (h *HeapGreater[T]) IsEmpty() bool {
//	return h.heapSize == 0
//}
//
//func (h *HeapGreater[T]) Size() int {
//	return h.heapSize
//}
//
//func (h *HeapGreater[T]) Contains(obj T) bool {
//	_, ok := h.indexMap[obj]
//	return ok
//}
//
//func (h *HeapGreater[T]) Peek() T {
//	return h.heap[0]
//}
//
//func (h *HeapGreater[T]) Push(obj T) {
//	h.heap = append(h.heap, obj)
//	h.indexMap[obj] = h.heapSize
//	h.heapInsert(h.heapSize)
//	h.heapSize++
//}
//
//func (h *HeapGreater[T]) Pop() T {
//	ans := h.heap[0]
//	h.swap(0, h.heapSize-1)
//	delete(h.indexMap, ans)
//	h.heap = h.heap[:h.heapSize-1]
//	h.heapSize--
//	h.heapify(0)
//	return ans
//}
//
//func (h *HeapGreater[T]) Remove(obj T) {
//	replace := h.heap[h.heapSize-1]
//	index := h.indexMap[obj]
//	delete(h.indexMap, obj)
//	h.heap = h.heap[:h.heapSize-1]
//	h.heapSize--
//
//	if any(obj) != any(replace) {
//		h.heap[index] = replace
//		h.indexMap[replace] = index
//		h.Resign(replace)
//	}
//}
//
//func (h *HeapGreater[T]) Resign(obj T) {
//	index := h.indexMap[obj]
//	h.heapInsert(index)
//	h.heapify(index)
//}
//
//func (h *HeapGreater[T]) GetAllElements() []T {
//	ans := make([]T, len(h.heap))
//	copy(ans, h.heap)
//	return ans
//}
//
//func (h *HeapGreater[T]) heapInsert(index int) {
//	for h.comp(h.heap[index], h.heap[(index-1)/2]) < 0 {
//		h.swap(index, (index-1)/2)
//		index = (index - 1) / 2
//	}
//}
//
//func (h *HeapGreater[T]) heapify(index int) {
//	left := index*2 + 1
//	for left < h.heapSize {
//		best := left
//		if left+1 < h.heapSize && h.comp(h.heap[left+1], h.heap[left]) < 0 {
//			best = left + 1
//		}
//		if h.comp(h.heap[best], h.heap[index]) >= 0 {
//			break
//		}
//		h.swap(best, index)
//		index = best
//		left = index*2 + 1
//	}
//}
//
//func (h *HeapGreater[T]) swap(i, j int) {
//	o1 := h.heap[i]
//	o2 := h.heap[j]
//	h.heap[i] = o2
//	h.heap[j] = o1
//	h.indexMap[o2] = i
//	h.indexMap[o1] = j
//}
