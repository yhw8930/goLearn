package main

import (
	"container/list"
	"fmt"
)

// 题目：只用两个栈实现队列的 add、poll、peek。
// 栈是后进先出，队列是先进先出，两者顺序相反。
// push 栈负责收集新元素，pop 栈为空时把 push 栈元素全部倒入 pop 栈，顺序就会反过来。
// 核心思路：倒数据必须一次倒干净，且 pop 栈非空时不能倒，才能保持队列顺序。
// 时间复杂度：单次操作均摊 O(1)。
// 空间复杂度：O(N)。

type TwoStacksQueue struct {
	stackPush *list.List
	stackPop  *list.List
}

func NewTwoStacksQueue() *TwoStacksQueue {
	return &TwoStacksQueue{
		stackPush: list.New(),
		stackPop:  list.New(),
	}
}

func (q *TwoStacksQueue) pushToPop() {
	if q.stackPop.Len() == 0 {
		for q.stackPush.Len() != 0 {
			val := q.stackPush.Remove(q.stackPush.Back())
			q.stackPop.PushBack(val)
		}
	}
}

func (q *TwoStacksQueue) Add(pushInt int) {
	q.stackPush.PushBack(pushInt)
	q.pushToPop()
}

func (q *TwoStacksQueue) Poll() int {
	if q.stackPop.Len() == 0 && q.stackPush.Len() == 0 {
		panic("Queue is empty!")
	}
	q.pushToPop()
	return q.stackPop.Remove(q.stackPop.Back()).(int)
}

func (q *TwoStacksQueue) Peek() int {
	if q.stackPop.Len() == 0 && q.stackPush.Len() == 0 {
		panic("Queue is empty!")
	}
	q.pushToPop()
	return q.stackPop.Back().Value.(int)
}

// 测试主函数
func main() {
	test := NewTwoStacksQueue()
	test.Add(1)
	test.Add(2)
	test.Add(3)

	fmt.Println(test.Peek())
	fmt.Println(test.Poll())

	fmt.Println(test.Peek())
	fmt.Println(test.Poll())

	fmt.Println(test.Peek())
	fmt.Println(test.Poll())
}
