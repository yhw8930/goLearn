package main

import (
	"container/list"
	"fmt"
)

// 题目：实现一个栈，除了 push/pop，还能 O(1) 返回当前栈内最小值。
// 普通栈想找最小值需要遍历所有元素。
// 额外维护一个最小值栈，让它记录每个数据栈状态下的当前最小值。
// 核心思路：每次压入数据时同步更新最小值栈，弹出时两个栈同步弹出。
// 时间复杂度：Push、Pop、GetMin 都是 O(1)。
// 空间复杂度：O(N)。

// MyStack1
type MyStack1 struct {
	stackData *list.List
	stackMin  *list.List
}

// NewMyStack1 使用“只在更小或相等时压入最小栈”的策略。
// 弹出数据时，如果弹出的值等于当前最小值，最小栈也同步弹出。
func NewMyStack1() *MyStack1 {
	return &MyStack1{stackData: list.New(), stackMin: list.New()}
}
func (s *MyStack1) Push(newNum int) {
	if s.stackMin.Len() == 0 {
		s.stackMin.PushBack(newNum)
	} else if newNum <= s.GetMin() {
		s.stackMin.PushBack(newNum)
	}
	s.stackData.PushBack(newNum)
}
func (s *MyStack1) Pop() int {
	if s.stackData.Len() == 0 {
		panic("empty")
	}
	val := s.stackData.Remove(s.stackData.Back()).(int)
	if val == s.GetMin() {
		s.stackMin.Remove(s.stackMin.Back())
	}
	return val
}
func (s *MyStack1) GetMin() int {
	return s.stackMin.Back().Value.(int)
}

// MyStack2
type MyStack2 struct {
	stackData *list.List
	stackMin  *list.List
}

// NewMyStack2 让最小栈和数据栈保持相同长度。
// 每次压入数据时，最小栈同步压入当前状态下的最小值。
func NewMyStack2() *MyStack2 {
	return &MyStack2{stackData: list.New(), stackMin: list.New()}
}
func (s *MyStack2) Push(newNum int) {
	if s.stackMin.Len() == 0 {
		s.stackMin.PushBack(newNum)
	} else if newNum < s.GetMin() {
		s.stackMin.PushBack(newNum)
	} else {
		s.stackMin.PushBack(s.GetMin())
	}
	s.stackData.PushBack(newNum)
}
func (s *MyStack2) Pop() int {
	if s.stackData.Len() == 0 {
		panic("empty")
	}
	s.stackMin.Remove(s.stackMin.Back())
	return s.stackData.Remove(s.stackData.Back()).(int)
}
func (s *MyStack2) GetMin() int {
	return s.stackMin.Back().Value.(int)
}

// 测试主函数
func main() {
	stack1 := NewMyStack1()
	stack1.Push(3)
	fmt.Println(stack1.GetMin()) // 3
	stack1.Push(4)
	fmt.Println(stack1.GetMin()) // 3
	stack1.Push(1)
	fmt.Println(stack1.GetMin()) // 1
	fmt.Println(stack1.Pop())    // 1
	fmt.Println(stack1.GetMin()) // 3

	fmt.Println("=============")

	stack2 := NewMyStack2()
	stack2.Push(3)
	fmt.Println(stack2.GetMin()) // 3
	stack2.Push(4)
	fmt.Println(stack2.GetMin()) // 3
	stack2.Push(1)
	fmt.Println(stack2.GetMin()) // 1
	fmt.Println(stack2.Pop())    // 1
	fmt.Println(stack2.GetMin()) // 3
}
