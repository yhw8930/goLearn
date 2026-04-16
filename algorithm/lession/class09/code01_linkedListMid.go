package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

// 上中点
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
