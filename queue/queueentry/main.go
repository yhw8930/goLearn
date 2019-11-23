package main

import (
	"fmt"
	"goLearn/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	//q.Push("fsfvdv")
	fmt.Println(q.Pop())

}
