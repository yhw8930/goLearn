package main

import "fmt"

// 题目：给定单链表，分别返回上中点、下中点、上中点前一个、下中点前一个。
// 链表长度奇偶不同，四个定义对应的返回位置不同。
// 核心思路：使用快慢指针，快指针一次走两步，慢指针一次走一步。
// 通过调整快慢指针的初始位置和循环停止条件，可以让慢指针停在目标节点。
// 时间复杂度：O(N)。
// 空间复杂度：O(1)。

type Node struct {
	Value int
	Next  *Node
}

// 上中点
// midOrUpMidNode 返回链表上中点。
// 长度为奇数时返回唯一中点，长度为偶数时返回两个中点中的上中点。
func midOrUpMidNode(head *Node) *Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	slow := head.Next
	fast := head.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 下中点
// midOrDownMidNode 返回链表下中点。
// 长度为奇数时返回唯一中点，长度为偶数时返回两个中点中的下中点。
func midOrDownMidNode(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head.Next
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 上中点的前一个
// midOrUpMidPreNode 返回上中点的前一个节点。
// 长度不足时没有前驱，返回 nil。
func midOrUpMidPreNode(head *Node) *Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	slow := head
	fast := head.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 下中点的前一个
// midOrDownMidPreNode 返回下中点的前一个节点。
// 通过调整快慢指针初始位置让 slow 停在目标前驱。
func midOrDownMidPreNode(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 对照方法1：上中点
// right1 是用数组收集节点的上中点对照方法。
// 它按长度直接计算目标下标，用于验证快慢指针版本。
// 时间复杂度：O(N)。
// 空间复杂度：O(N)，数组保存节点。
func right1(head *Node) *Node {
	if head == nil {
		return nil
	}
	arr := []*Node{}
	for cur := head; cur != nil; cur = cur.Next {
		arr = append(arr, cur)
	}
	return arr[(len(arr)-1)/2]
}

// 对照方法2：下中点
// right2 是用数组收集节点的下中点对照方法。
// 它按长度直接计算下中点下标。
// 时间复杂度：O(N)。
// 空间复杂度：O(N)，数组保存节点。
func right2(head *Node) *Node {
	if head == nil {
		return nil
	}
	arr := []*Node{}
	for cur := head; cur != nil; cur = cur.Next {
		arr = append(arr, cur)
	}
	return arr[len(arr)/2]
}

// 对照方法3：上中点前一个
// right3 是上中点前驱的数组对照方法。
// 它用于和 midOrUpMidPreNode 交叉验证。
// 时间复杂度：O(N)。
// 空间复杂度：O(N)，数组保存节点。
func right3(head *Node) *Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	arr := []*Node{}
	for cur := head; cur != nil; cur = cur.Next {
		arr = append(arr, cur)
	}
	return arr[(len(arr)-3)/2]
}

// 对照方法4：下中点前一个
// right4 是下中点前驱的数组对照方法。
// 它用于和 midOrDownMidPreNode 交叉验证。
// 时间复杂度：O(N)。
// 空间复杂度：O(N)，数组保存节点。
func right4(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}
	arr := []*Node{}
	for cur := head; cur != nil; cur = cur.Next {
		arr = append(arr, cur)
	}
	return arr[(len(arr)-2)/2]
}

func main() {
	// 构建链表 0~8
	head := &Node{0, nil}
	cur := head
	for i := 1; i <= 8; i++ {
		cur.Next = &Node{i, nil}
		cur = cur.Next
	}

	printNode(midOrUpMidNode(head))
	printNode(right1(head))

	printNode(midOrDownMidNode(head))
	printNode(right2(head))

	printNode(midOrUpMidPreNode(head))
	printNode(right3(head))

	printNode(midOrDownMidPreNode(head))
	printNode(right4(head))
}

func printNode(n *Node) {
	if n != nil {
		fmt.Println(n.Value)
	} else {
		fmt.Println("无")
	}
}
