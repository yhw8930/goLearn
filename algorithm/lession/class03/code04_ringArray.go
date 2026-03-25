package main

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
