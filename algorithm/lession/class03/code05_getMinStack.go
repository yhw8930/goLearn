package main

import (
	"container/list"
	"fmt"
)

// MyStack1
type MyStack1 struct {
	stackData *list.List
	stackMin  *list.List
}

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
