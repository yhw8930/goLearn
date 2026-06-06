package main

import "fmt"

// 题目：给定链表头节点和 pivot，把链表调整为小于区、等于区、大于区的顺序。
// 要求每个节点仍然是原节点，只改变节点之间的连接关系。
// 核心思路：数组方法把节点放入切片后做荷兰国旗划分。
// 链表方法更直接：分别收集小于、等于、大于三条链，最后按顺序连接非空链表。
// 时间复杂度：O(N)。
// 空间复杂度：数组版 O(N)，链表版 O(1)。

/*
======================== 数据结构 ========================
*/

type Code03Node struct {
	Value int
	Next  *Code03Node
}

/*
======================== 方法1：数组 partition ========================

1. 问题：
给定一个单链表和一个 pivot，将链表按 pivot 分成三部分：
小于、等于、大于，并返回新的头节点

2. 思路：
- 遍历链表，将所有节点放入数组
- 在数组上做“荷兰国旗问题”的 partition
- 再将数组重新串成链表

本质：链表问题转数组问题

时间复杂度：O(N)
空间复杂度：O(N)
*/

// listPartition1 是数组辅助版本。
// 它把链表节点放入数组，再在数组上做荷兰国旗划分，最后重新串链。
// 时间复杂度：O(N)。
// 空间复杂度：O(N)，数组保存所有节点。
func listPartition1(head *Code03Node, pivot int) *Code03Node {
	if head == nil {
		return head
	}

	cur := head
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}

	Code03NodeArr := make([]*Code03Node, count)
	cur = head
	for i := 0; i < count; i++ {
		Code03NodeArr[i] = cur
		cur = cur.Next
	}

	arrPartition(Code03NodeArr, pivot)

	for i := 1; i < count; i++ {
		Code03NodeArr[i-1].Next = Code03NodeArr[i]
	}
	Code03NodeArr[count-1].Next = nil

	return Code03NodeArr[0]
}

func arrPartition(Code03NodeArr []*Code03Node, pivot int) {
	small := -1
	big := len(Code03NodeArr)
	index := 0

	for index != big {
		if Code03NodeArr[index].Value < pivot {
			small++
			swap(Code03NodeArr, small, index)
			index++
		} else if Code03NodeArr[index].Value == pivot {
			index++
		} else {
			big--
			swap(Code03NodeArr, big, index)
		}
	}
}

func swap(arr []*Code03Node, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

/*
======================== 方法2：三链表（推荐） ========================

1. 问题：
同上

2. 思路：
- 遍历链表，将节点拆分到三个链表：
  small（小于 pivot）、equal（等于 pivot）、big（大于 pivot）
- 每个链表维护 head 和 tail
- 最后按 small → equal → big 连接

关键点：
- 每次都要断开原链表（head.Next = nil）
- 拼接时要考虑某一段为空的情况

时间复杂度：O(N)
空间复杂度：O(1)
*/

// listPartition2 是原地链表版本。
// 它分别维护小于、等于、大于三条链的头尾，最后把三段连接起来。
// 时间复杂度：O(N)。
// 空间复杂度：O(1)，只使用有限几个头尾指针。
func listPartition2(head *Code03Node, pivot int) *Code03Node {
	var sH, sT *Code03Node
	var eH, eT *Code03Node
	var mH, mT *Code03Node

	var next *Code03Node

	for head != nil {
		next = head.Next
		head.Next = nil

		if head.Value < pivot {
			if sH == nil {
				sH = head
				sT = head
			} else {
				sT.Next = head
				sT = head
			}
		} else if head.Value == pivot {
			if eH == nil {
				eH = head
				eT = head
			} else {
				eT.Next = head
				eT = head
			}
		} else {
			if mH == nil {
				mH = head
				mT = head
			} else {
				mT.Next = head
				mT = head
			}
		}

		head = next
	}

	// small → equal
	if sT != nil {
		sT.Next = eH
		if eT == nil {
			eT = sT
		}
	}

	// equal → big
	if eT != nil {
		eT.Next = mH
	}

	if sH != nil {
		return sH
	}
	if eH != nil {
		return eH
	}
	return mH
}

/*
======================== 工具函数 ========================
*/

func code03PrintLinkedList(head *Code03Node) {
	fmt.Print("Linked List: ")
	for head != nil {
		fmt.Print(head.Value, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	head := &Code03Node{7, nil}
	head.Next = &Code03Node{9, nil}
	head.Next.Next = &Code03Node{1, nil}
	head.Next.Next.Next = &Code03Node{8, nil}
	head.Next.Next.Next.Next = &Code03Node{5, nil}
	head.Next.Next.Next.Next.Next = &Code03Node{2, nil}
	head.Next.Next.Next.Next.Next.Next = &Code03Node{5, nil}

	code03PrintLinkedList(head)

	head = listPartition2(head, 5)

	code03PrintLinkedList(head)
}
