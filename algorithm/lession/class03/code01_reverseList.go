package main

import (
	"math/rand"
	"time"
)

// 题目：分别反转单链表和双链表，返回新的头节点。
// 单链表反转时，需要保存 next，随后把当前节点的 next 指向前一个节点。
// 双链表除了修改 next，还要同步修改 last，让前后指针都反向。
// 核心思路：用 pre、next 辅助遍历链表，每处理一个节点就把它接到已反转部分前面。
// 时间复杂度：O(N)。
// 空间复杂度：O(1)。

type Node struct {
	value int
	next  *Node
}

type DoubleNode struct {
	value      int
	last, next *DoubleNode
}

// reverseLinkedList 是单链表原地反转的标准写法。
// pre 指向已反转部分头部，next 临时保存后续链表，逐个节点改 next 指向。
// 时间复杂度：O(N)。
// 空间复杂度：O(1)。
func reverseLinkedList(head *Node) *Node {
	var pre, next *Node
	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}

// reverseLinkedList2 是单链表反转的另一种变量组织方式。
// 核心仍然是遍历时保存下一个节点，并把当前节点接到已反转链表前面。
// 时间复杂度：O(N)。
// 空间复杂度：O(1)。
func reverseLinkedList2(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	revHead := reverseLinkedList2(head.next)
	tail := head.next
	tail.next = head
	head.next = nil
	return revHead
}

// reverseDoubleList 反转双链表。
// 每个节点既要修改 next，也要修改 last，保证两个方向的链接都被反过来。
// 时间复杂度：O(N)。
// 空间复杂度：O(1)。
func reverseDoubleList(head *DoubleNode) *DoubleNode {
	var pre, next *DoubleNode
	for head != nil {
		next = head.next
		head.next = pre
		head.last = next
		pre = head
		head = next
	}
	return pre
}

// 单链表暴力反转测试方法（容器法）
// testReverseLinkedList 是测试用对照实现。
// 它把链表节点收集到数组后按反向顺序重新连接。
// 时间复杂度：O(N)。
// 空间复杂度：O(N)，数组保存所有节点。
func testReverseLinkedList(head *Node) *Node {
	if head == nil {
		return nil
	}
	list := make([]*Node, 0)
	for head != nil {
		list = append(list, head)
		head = head.next
	}
	list[0].next = nil
	n := len(list)
	for i := 1; i < n; i++ {
		list[i].next = list[i-1]
	}
	return list[n-1]
}

// 双向链表暴力反转测试方法
// testReverseDoubleList 是双链表的测试对照实现。
// 它借助数组重连节点，用于和原地反转版本比对。
// 时间复杂度：O(N)。
// 空间复杂度：O(N)，数组保存所有节点。
func testReverseDoubleList(head *DoubleNode) *DoubleNode {
	if head == nil {
		return nil
	}
	list := make([]*DoubleNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.next
	}
	list[0].next = nil
	pre := list[0]
	n := len(list)
	for i := 1; i < n; i++ {
		cur := list[i]
		cur.last = nil
		cur.next = pre
		pre.last = cur
		pre = cur
	}
	return list[n-1]
}

// 生成随机单链表
func generateRandomLinkedList(length, value int) *Node {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(length + 1)
	if size == 0 {
		return nil
	}
	size--
	head := &Node{value: rand.Intn(value + 1)}
	pre := head
	for size != 0 {
		cur := &Node{value: rand.Intn(value + 1)}
		pre.next = cur
		pre = cur
		size--
	}
	return head
}

// 生成随机双向链表
func generateRandomDoubleList(length, value int) *DoubleNode {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(length + 1)
	if size == 0 {
		return nil
	}
	size--
	head := &DoubleNode{value: rand.Intn(value + 1)}
	pre := head
	for size != 0 {
		cur := &DoubleNode{value: rand.Intn(value + 1)}
		pre.next = cur
		cur.last = pre
		pre = cur
		size--
	}
	return head
}

// 获取单链表原始顺序
func getLinkedListOriginOrder(head *Node) []int {
	ans := make([]int, 0)
	for head != nil {
		ans = append(ans, head.value)
		head = head.next
	}
	return ans
}

// 检查单链表是否反转正确
func checkLinkedListReverse(origin []int, head *Node) bool {
	for i := len(origin) - 1; i >= 0; i-- {
		if origin[i] != head.value {
			return false
		}
		head = head.next
	}
	return true
}

// 获取双向链表原始顺序
func getDoubleListOriginOrder(head *DoubleNode) []int {
	ans := make([]int, 0)
	for head != nil {
		ans = append(ans, head.value)
		head = head.next
	}
	return ans
}

// 检查双向链表是否反转正确（正反都检查）
func checkDoubleListReverse(origin []int, head *DoubleNode) bool {
	end := head
	for i := len(origin) - 1; i >= 0; i-- {
		if origin[i] != head.value {
			return false
		}
		end = head
		head = head.next
	}
	for i := 0; i < len(origin); i++ {
		if origin[i] != end.value {
			return false
		}
		end = end.last
	}
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机种子

	len := 50
	value := 100
	testTime := 100000

	println("test begin!")

	for i := 0; i < testTime; i++ {
		// 单链表测试
		node1 := generateRandomLinkedList(len, value)
		list1 := getLinkedListOriginOrder(node1)
		node1 = reverseLinkedList(node1)
		if !checkLinkedListReverse(list1, node1) {
			println("Oops1!")
			return
		}

		// 单链表递归法测试
		node2 := generateRandomLinkedList(len, value)
		list2 := getLinkedListOriginOrder(node2)
		node2 = reverseLinkedList2(node2)
		if !checkLinkedListReverse(list2, node2) {
			println("Oops2!")
			return
		}

		// 单链表暴力法测试
		nodeTest1 := generateRandomLinkedList(len, value)
		listTest1 := getLinkedListOriginOrder(nodeTest1)
		nodeTest1 = testReverseLinkedList(nodeTest1)
		if !checkLinkedListReverse(listTest1, nodeTest1) {
			println("OopsTest1!")
			return
		}

		// 双向链表测试
		node3 := generateRandomDoubleList(len, value)
		list3 := getDoubleListOriginOrder(node3)
		node3 = reverseDoubleList(node3)
		if !checkDoubleListReverse(list3, node3) {
			println("Oops3!")
			return
		}

		// 双向链表暴力法测试
		node4 := generateRandomDoubleList(len, value)
		list4 := getDoubleListOriginOrder(node4)
		node4 = testReverseDoubleList(node4)
		if !checkDoubleListReverse(list4, node4) {
			println("Oops4!")
			return
		}
	}

	println("test finish!")
}
