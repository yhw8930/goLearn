package main

import "fmt"

/*
======================== 数据结构 ========================
*/

type code04Node struct {
	Value int
	Next  *code04Node
	Rand  *code04Node
}

/*
======================== 方法1：HashMap复制 ========================

1. 问题：
给定一个链表，每个节点除了 next 指针，还有一个 rand 指针（随机指向任意节点或 null）。
要求：复制整个链表，返回新链表头节点。

2. 思路：
- 用 HashMap 记录 “老节点 -> 新节点”
- 第一次遍历：创建所有新节点
- 第二次遍历：连接 next 和 rand 指针

核心思想：
用空间换映射关系

时间复杂度：O(N)
空间复杂度：O(N)
*/

func copyListWithRand1(head *code04Node) *code04Node {
	if head == nil {
		return nil
	}

	m := make(map[*code04Node]*code04Node)

	// 1. 复制节点
	cur := head
	for cur != nil {
		m[cur] = &code04Node{Value: cur.Value}
		cur = cur.Next
	}

	// 2. 连接 next + rand
	cur = head
	for cur != nil {
		m[cur].Next = m[cur.Next]
		m[cur].Rand = m[cur.Rand]
		cur = cur.Next
	}

	return m[head]
}

/*
======================== 方法2：O(1)空间复制（推荐） ========================

1. 问题：
同上

2. 思路：
分三步：

Step1：
在每个老节点后面插入复制节点
1 -> 2 -> 3
变成：
1 -> 1' -> 2 -> 2' -> 3 -> 3'

Step2：
设置 random 指针：
curCopy.rand = cur.rand.next

Step3：
拆分链表（恢复原链表 + 提取新链表）

核心思想：
不使用额外空间，利用链表结构“夹层映射”

时间复杂度：O(N)
空间复杂度：O(1)
*/

func copyListWithRand2(head *code04Node) *code04Node {
	if head == nil {
		return nil
	}

	cur := head
	var next *code04Node

	// Step1：复制节点插入原链表
	for cur != nil {
		next = cur.Next
		cur.Next = &code04Node{Value: cur.Value}
		cur.Next.Next = next
		cur = next
	}

	// Step2：设置 rand
	cur = head
	var copy *code04Node
	for cur != nil {
		next = cur.Next.Next
		copy = cur.Next

		if cur.Rand != nil {
			copy.Rand = cur.Rand.Next
		}

		cur = next
	}

	// Step3：拆分链表
	res := head.Next
	cur = head

	for cur != nil {
		next = cur.Next.Next
		copy = cur.Next

		cur.Next = next
		if next != nil {
			copy.Next = next.Next
		}

		cur = next
	}

	return res
}

/*
======================== 打印函数 ========================
*/

func printRandLinkedList(head *code04Node) {
	cur := head

	fmt.Print("order: ")
	for cur != nil {
		fmt.Print(cur.Value, " ")
		cur = cur.Next
	}
	fmt.Println()

	cur = head
	fmt.Print("rand:  ")
	for cur != nil {
		if cur.Rand == nil {
			fmt.Print("- ")
		} else {
			fmt.Print(cur.Rand.Value, " ")
		}
		cur = cur.Next
	}
	fmt.Println()
}

/*
======================== 测试 ========================
*/

func main() {
	var head *code04Node

	fmt.Println("空链表测试：")
	printRandLinkedList(copyListWithRand1(head))
	printRandLinkedList(copyListWithRand2(head))

	fmt.Println("=====================")

	head = &code04Node{Value: 1}
	head.Next = &code04Node{Value: 2}
	head.Next.Next = &code04Node{Value: 3}
	head.Next.Next.Next = &code04Node{Value: 4}
	head.Next.Next.Next.Next = &code04Node{Value: 5}
	head.Next.Next.Next.Next.Next = &code04Node{Value: 6}

	head.Rand = head.Next.Next.Next.Next.Next      // 1 -> 6
	head.Next.Rand = head.Next.Next.Next.Next      // 2 -> 5
	head.Next.Next.Rand = head.Next.Next.Next      // 3 -> 4
	head.Next.Next.Next.Rand = head.Next.Next      // 4 -> 3
	head.Next.Next.Next.Next.Rand = nil            // 5 -> null
	head.Next.Next.Next.Next.Next.Rand = head.Next // 6 -> 2

	fmt.Println("原始链表：")
	printRandLinkedList(head)

	fmt.Println("=====================")

	res1 := copyListWithRand1(head)
	fmt.Println("方法1（HashMap）：")
	printRandLinkedList(res1)

	fmt.Println("=====================")

	res2 := copyListWithRand2(head)
	fmt.Println("方法2（O(1)空间）：")
	printRandLinkedList(res2)

	fmt.Println("=====================")

	fmt.Println("原链表（验证未被破坏）：")
	printRandLinkedList(head)
}
