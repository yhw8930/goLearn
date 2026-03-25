package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

// TwoQueueStack 两个队列实现栈
type TwoQueueStack[T any] struct {
	queue *list.List
	help  *list.List
}

// NewTwoQueueStack 构造函数
func NewTwoQueueStack[T any]() *TwoQueueStack[T] {
	return &TwoQueueStack[T]{
		queue: list.New(),
		help:  list.New(),
	}
}

// Push 入栈
func (s *TwoQueueStack[T]) Push(value T) {
	s.queue.PushBack(value)
}

// Poll 出栈
func (s *TwoQueueStack[T]) Poll() T {
	for s.queue.Len() > 1 {
		s.help.PushBack(s.queue.Remove(s.queue.Front()))
	}
	ans := s.queue.Remove(s.queue.Front()).(T)
	// 交换 queue 和 help
	tmp := s.queue
	s.queue = s.help
	s.help = tmp
	return ans
}

// Peek 查看栈顶（不删除）
func (s *TwoQueueStack[T]) Peek() T {
	for s.queue.Len() > 1 {
		s.help.PushBack(s.queue.Remove(s.queue.Front()))
	}
	ans := s.queue.Remove(s.queue.Front()).(T)
	s.help.PushBack(ans)
	// 交换 queue 和 help
	tmp := s.queue
	s.queue = s.help
	s.help = tmp
	return ans
}

// IsEmpty 是否为空
func (s *TwoQueueStack[T]) IsEmpty() bool {
	return s.queue.Len() == 0
}

// 对数器测试
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("test begin")
	myStack := NewTwoQueueStack[int]()
	testStack := list.New()

	testTime := 100000
	maxVal := 1000000

	for i := 0; i < testTime; i++ {
		if myStack.IsEmpty() {
			if testStack.Len() != 0 {
				fmt.Println("Oops")
				return
			}
			num := rand.Intn(maxVal)
			myStack.Push(num)
			testStack.PushBack(num)
		} else {
			r := rand.Float64()
			if r < 0.25 {
				num := rand.Intn(maxVal)
				myStack.Push(num)
				testStack.PushBack(num)
			} else if r < 0.5 {
				myPeek := myStack.Peek()
				tPeek := testStack.Back().Value.(int)
				if myPeek != tPeek {
					fmt.Println("Oops")
					return
				}
			} else if r < 0.75 {
				myPoll := myStack.Poll()
				tPop := testStack.Remove(testStack.Back()).(int)
				if myPoll != tPop {
					fmt.Println("Oops")
					return
				}
			} else {
				if myStack.IsEmpty() != (testStack.Len() == 0) {
					fmt.Println("Oops")
					return
				}
			}
		}
	}
	fmt.Println("test finish!")
}
