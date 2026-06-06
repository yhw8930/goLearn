package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

// 题目：先实现双端队列，再用它分别封装出栈和队列。
// 双端队列允许从头部和尾部插入、弹出元素。
// 栈要求后进先出，可以固定从头部压入、从头部弹出；队列要求先进先出，可以一端入另一端出。
// 核心思路：同一个底层双端结构，通过不同端点操作组合出不同数据结构语义。
// 时间复杂度：每次入队、出队、入栈、出栈都是 O(1)。
// 空间复杂度：O(N)。

type QueueNode[T any] struct {
	value T
	last  *QueueNode[T]
	next  *QueueNode[T]
}

type DoubleEndsQueue[T any] struct {
	head *QueueNode[T]
	tail *QueueNode[T]
}

func (q *DoubleEndsQueue[T]) addFromHead(value T) {
	cur := &QueueNode[T]{value: value}
	if q.head == nil {
		q.head = cur
		q.tail = cur
	} else {
		cur.next = q.head
		q.head.last = cur
		q.head = cur
	}
}

func (q *DoubleEndsQueue[T]) addFromBottom(value T) {
	cur := &QueueNode[T]{value: value}
	if q.head == nil {
		q.head = cur
		q.tail = cur
	} else {
		cur.last = q.tail
		q.tail.next = cur
		q.tail = cur
	}
}

func (q *DoubleEndsQueue[T]) popFromHead() *T {
	if q.head == nil {
		return nil
	}
	cur := q.head
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
		q.head.last = nil
		cur.next = nil
	}
	return &cur.value
}

func (q *DoubleEndsQueue[T]) popFromBottom() *T {
	if q.head == nil {
		return nil
	}
	cur := q.tail
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.tail = q.tail.last
		q.tail.next = nil
		cur.last = nil
	}
	return &cur.value
}

func (q *DoubleEndsQueue[T]) isEmpty() bool {
	return q.head == nil
}

// 栈
type MyStack[T any] struct {
	queue *DoubleEndsQueue[T]
}

// NewMyStack 基于双端队列创建栈。
// 栈只从同一端压入和弹出，因此满足后进先出。
func NewMyStack[T any]() *MyStack[T] {
	return &MyStack[T]{queue: &DoubleEndsQueue[T]{}}
}

func (s *MyStack[T]) Push(value T) {
	s.queue.addFromHead(value)
}

func (s *MyStack[T]) Pop() *T {
	return s.queue.popFromHead()
}

func (s *MyStack[T]) IsEmpty() bool {
	return s.queue.isEmpty()
}

// 队列
type MyQueue[T any] struct {
	queue *DoubleEndsQueue[T]
}

// NewMyQueue 基于双端队列创建队列。
// 队列从一端进入、另一端弹出，因此满足先进先出。
func NewMyQueue[T any]() *MyQueue[T] {
	return &MyQueue[T]{queue: &DoubleEndsQueue[T]{}}
}

func (q *MyQueue[T]) Push(value T) {
	q.queue.addFromHead(value)
}

func (q *MyQueue[T]) Poll() *T {
	return q.queue.popFromBottom()
}

func (q *MyQueue[T]) IsEmpty() bool {
	return q.queue.isEmpty()
}

func isEqual(o1, o2 *int) bool {
	if o1 == nil && o2 != nil {
		return false
	}
	if o1 != nil && o2 == nil {
		return false
	}
	if o1 == nil && o2 == nil {
		return true
	}
	return *o1 == *o2
}

func main() {
	rand.Seed(time.Now().UnixNano())
	oneTestDataNum := 100
	value := 10000
	testTimes := 100000

	fmt.Println("测试开始...")

	for i := 0; i < testTimes; i++ {
		myStack := NewMyStack[int]()
		myQueue := NewMyQueue[int]()
		stack := list.New()
		queue := list.New()

		for j := 0; j < oneTestDataNum; j++ {
			nums := rand.Intn(value)
			if myStack.IsEmpty() {
				myStack.Push(nums)
				stack.PushBack(nums)
			} else {
				if rand.Float64() < 0.5 {
					myStack.Push(nums)
					stack.PushBack(nums)
				} else {
					myPop := myStack.Pop()
					sPop := stack.Remove(stack.Back()).(int)
					if !isEqual(myPop, &sPop) {
						fmt.Println("oops! stack 出错")
						return
					}
				}
			}

			numq := rand.Intn(value)
			if myQueue.IsEmpty() {
				myQueue.Push(numq)
				queue.PushBack(numq)
			} else {
				if rand.Float64() < 0.5 {
					myQueue.Push(numq)
					queue.PushBack(numq)
				} else {
					myPoll := myQueue.Poll()
					qPoll := queue.Remove(queue.Front()).(int)
					if !isEqual(myPoll, &qPoll) {
						fmt.Println("oops! queue 出错")
						return
					}
				}
			}
		}
	}
	fmt.Println("finish! ✅ 测试全部通过")
}
