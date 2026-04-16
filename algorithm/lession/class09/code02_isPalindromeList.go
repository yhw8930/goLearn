package main

import (
	"container/list"
	"fmt"
)

// Code02Node 链表节点结构
type Code02Node struct {
	value int
	next  *Code02Node
}

// NewCode02Node 创建链表节点
func NewCode02Node(data int) *Code02Node {
	return &Code02Node{value: data}
}

// isPalindrome1 需要 O(n) 额外空间，使用栈存储全部节点
func isPalindrome1(head *Code02Node) bool {
	stack := list.New()
	cur := head
	// 把所有节点压入栈
	for cur != nil {
		stack.PushBack(cur)
		cur = cur.next
	}
	// 遍历链表，同时弹出栈节点对比
	for head != nil {
		Code02Node := stack.Remove(stack.Back()).(*Code02Node)
		if head.value != Code02Node.value {
			return false
		}
		head = head.next
	}
	return true
}

// isPalindrome2 需要 O(n/2) 额外空间，只存右半部分节点
func isPalindrome2(head *Code02Node) bool {
	if head == nil || head.next == nil {
		return true
	}
	right := head.next
	cur := head
	// 快慢指针找到右半部分第一个节点
	for cur.next != nil && cur.next.next != nil {
		right = right.next
		cur = cur.next.next
	}

	stack := list.New()
	// 右半部分节点入栈
	for right != nil {
		stack.PushBack(right)
		right = right.next
	}
	// 左半部分与栈中节点对比
	for stack.Len() > 0 {
		Code02Node := stack.Remove(stack.Back()).(*Code02Node)
		if head.value != Code02Node.value {
			return false
		}
		head = head.next
	}
	return true
}

// isPalindrome3 O(1) 额外空间，原地反转右半部分链表对比
func isPalindrome3(head *Code02Node) bool {
	if head == nil || head.next == nil {
		return true
	}

	// 快慢指针找到链表中点
	n1 := head
	n2 := head
	for n2.next != nil && n2.next.next != nil {
		n1 = n1.next      // n1 到达中点
		n2 = n2.next.next // n2 到达末尾
	}

	// 反转右半部分链表
	n2 = n1.next  // n2 指向右半部分第一个节点
	n1.next = nil // 中点断开
	var n3 *Code02Node
	for n2 != nil {
		n3 = n2.next // 保存下一个节点
		n2.next = n1 // 反转指针
		n1 = n2      // n1 移动
		n2 = n3      // n2 移动
	}

	n3 = n1   // 保存反转后的尾节点（用于恢复链表）
	n2 = head // n2 回到链表头部
	res := true

	// 对比左右两部分是否相等
	for n1 != nil && n2 != nil {
		if n1.value != n2.value {
			res = false
			break
		}
		n1 = n1.next
		n2 = n2.next
	}

	// 恢复链表（把反转的右半部分还原）
	n1 = n3.next
	n3.next = nil
	for n1 != nil {
		n2 = n1.next
		n1.next = n3
		n3 = n1
		n1 = n2
	}

	return res
}

// printLinkedList 打印链表
func printLinkedList(Code02Node *Code02Node) {
	fmt.Print("Linked List: ")
	for Code02Node != nil {
		fmt.Printf("%d ", Code02Node.value)
		Code02Node = Code02Node.next
	}
	fmt.Println()
}

func main() {
	// 测试用例1：空链表
	var head *Code02Node
	printLinkedList(head)
	fmt.Printf("%t | %t | %t |\n", isPalindrome1(head), isPalindrome2(head), isPalindrome3(head))
	printLinkedList(head)
	fmt.Println("=========================")

	// 测试用例2：单个节点
	head = NewCode02Node(1)
	printLinkedList(head)
	fmt.Printf("%t | %t | %t |\n", isPalindrome1(head), isPalindrome2(head), isPalindrome3(head))
	printLinkedList(head)
	fmt.Println("=========================")

	// 测试用例3：1->2
	head = NewCode02Node(1)
	head.next = NewCode02Node(2)
	printLinkedList(head)
	fmt.Printf("%t | %t | %t |\n", isPalindrome1(head), isPalindrome2(head), isPalindrome3(head))
	printLinkedList(head)
	fmt.Println("=========================")

	// 测试用例4：1->1
	head = NewCode02Node(1)
	head.next = NewCode02Node(1)
	printLinkedList(head)
	fmt.Printf("%t | %t | %t |\n", isPalindrome1(head), isPalindrome2(head), isPalindrome3(head))
	printLinkedList(head)
	fmt.Println("=========================")

	// 测试用例5：1->2->3
	head = NewCode02Node(1)
	head.next = NewCode02Node(2)
	head.next.next = NewCode02Node(3)
	printLinkedList(head)
	fmt.Printf("%t | %t | %t |\n", isPalindrome1(head), isPalindrome2(head), isPalindrome3(head))
	printLinkedList(head)
	fmt.Println("=========================")

	// 测试用例6：1->2->1
	head = NewCode02Node(1)
	head.next = NewCode02Node(2)
	head.next.next = NewCode02Node(1)
	printLinkedList(head)
	fmt.Printf("%t | %t | %t |\n", isPalindrome1(head), isPalindrome2(head), isPalindrome3(head))
	printLinkedList(head)
	fmt.Println("=========================")

	// 测试用例7：1->2->3->1
	head = NewCode02Node(1)
	head.next = NewCode02Node(2)
	head.next.next = NewCode02Node(3)
	head.next.next.next = NewCode02Node(1)
	printLinkedList(head)
	fmt.Printf("%t | %t | %t |\n", isPalindrome1(head), isPalindrome2(head), isPalindrome3(head))
	printLinkedList(head)
	fmt.Println("=========================")

	// 测试用例8：1->2->2->1
	head = NewCode02Node(1)
	head.next = NewCode02Node(2)
	head.next.next = NewCode02Node(2)
	head.next.next.next = NewCode02Node(1)
	printLinkedList(head)
	fmt.Printf("%t | %t | %t |\n", isPalindrome1(head), isPalindrome2(head), isPalindrome3(head))
	printLinkedList(head)
	fmt.Println("=========================")

	// 测试用例9：1->2->3->2->1
	head = NewCode02Node(1)
	head.next = NewCode02Node(2)
	head.next.next = NewCode02Node(3)
	head.next.next.next = NewCode02Node(2)
	head.next.next.next.next = NewCode02Node(1)
	printLinkedList(head)
	fmt.Printf("%t | %t | %t |\n", isPalindrome1(head), isPalindrome2(head), isPalindrome3(head))
	printLinkedList(head)
	fmt.Println("=========================")
}
