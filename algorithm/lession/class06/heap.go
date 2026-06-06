package main

// 题目：提供一个 int 小根堆适配器，供标准库 container/heap 使用。
// Go 的 heap 包要求类型实现 Len、Less、Swap、Push、Pop 五个方法。
// 核心思路：Less 使用小于号时堆顶就是最小值，Push/Pop 负责维护底层切片。
// 这个文件只定义通用堆能力，具体算法文件可以直接复用它。
// 时间复杂度：标准堆 Push/Pop 为 O(logN)。
// 空间复杂度：O(N)。

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 小根堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
