package main

import (
	"container/list"
	"fmt"
)

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
