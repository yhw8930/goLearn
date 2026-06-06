package main

// 题目：用固定长度数组实现队列，支持入队和出队。
// 普通数组队列如果一直移动头指针，会浪费前面已经弹出的空间。
// 环形数组让 push 和 poll 指针到达末尾后回到 0 位置继续使用。
// 核心思路：用 size 区分队列为空还是已满，用 nextIndex 处理指针循环。
// 时间复杂度：入队和出队都是 O(1)。
// 空间复杂度：O(N)，N 为固定数组容量。

// MyRingQueue 用环形数组实现队列（固定长度）
type MyRingQueue struct {
	arr   []int
	pushi int // 入队下标（end）
	polli int // 出队下标（begin）
	size  int // 当前元素个数
	limit int // 队列最大容量
}

// NewMyRingQueue 创建一个固定容量的队列
func NewMyRingQueue(limit int) *MyRingQueue {
	return &MyRingQueue{
		arr:   make([]int, limit),
		pushi: 0,
		polli: 0,
		size:  0,
		limit: limit,
	}
}

// Push 入队
func (q *MyRingQueue) Push(value int) {
	if q.size == q.limit {
		panic("队列满了，不能再加了")
	}
	q.size++
	q.arr[q.pushi] = value
	q.pushi = q.NextIndex(q.pushi)
}

// Pop 出队
func (q *MyRingQueue) Pop() int {
	if q.size == 0 {
		panic("队列空了，不能再拿了")
	}
	q.size--
	ans := q.arr[q.polli]
	q.polli = q.NextIndex(q.polli)
	return ans
}

// IsEmpty 判断是否为空
func (q *MyRingQueue) IsEmpty() bool {
	return q.size == 0
}

// nextIndex 计算下一个下标（环形）
func (q *MyRingQueue) NextIndex(i int) int {
	if i < q.limit-1 {
		return i + 1
	}
	return 0
}
