package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

// GetIntersectNode 两个链表的第一个相交节点
//
// 【问题】
// 给定两个可能有环/无环的链表，返回第一个相交节点
//
// 【解题思路】
// 1. 分别判断两个链表是否有环
// 2. 三种情况：
//   - 都无环：按无环相交方式处理
//   - 都有环：按有环相交方式处理
//   - 一个有环一个无环：一定不相交
//
// 【时空复杂度】
// 时间：O(N)
// 空间：O(1)
func GetIntersectNode(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}

	loop1 := getLoopNode(head1)
	loop2 := getLoopNode(head2)

	if loop1 == nil && loop2 == nil {
		return noLoop(head1, head2)
	}
	if loop1 != nil && loop2 != nil {
		return bothLoop(head1, loop1, head2, loop2)
	}
	return nil
}

// getLoopNode 找链表第一个入环节点
//
// 【问题】
// 判断链表是否有环，并返回入环节点
//
// 【解题思路】
// 1. 快慢指针相遇判环
// 2. 相遇后 fast回到head
// 3. slow和fast同步走，再次相遇即入环点
//
// 【时空复杂度】
// 时间：O(N)
// 空间：O(1)
func getLoopNode(head *Node) *Node {
	if head == nil || head.next == nil || head.next.next == nil {
		return nil
	}

	slow := head.next
	fast := head.next.next

	for slow != fast {
		if fast.next == nil || fast.next.next == nil {
			return nil
		}
		slow = slow.next
		fast = fast.next.next
	}

	fast = head
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}

	return slow
}

// noLoop 无环链表相交问题
//
// 【问题】
// 两个无环链表，找第一个相交节点
//
// 【解题思路】
// 1. 先算长度差
// 2. 长链先走差值
// 3. 同步移动找相交点
//
// 【时空复杂度】
// 时间：O(N)
// 空间：O(1)
func noLoop(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}

	cur1, cur2 := head1, head2
	n := 0

	for cur1.next != nil {
		n++
		cur1 = cur1.next
	}
	for cur2.next != nil {
		n--
		cur2 = cur2.next
	}

	if cur1 != cur2 {
		return nil
	}

	if n > 0 {
		cur1 = head1
		cur2 = head2
	} else {
		cur1 = head2
		cur2 = head1
		n = -n
	}

	for n != 0 {
		cur1 = cur1.next
		n--
	}

	for cur1 != cur2 {
		cur1 = cur1.next
		cur2 = cur2.next
	}

	return cur1
}

// bothLoop 有环链表相交问题
//
// 【问题】
// 两个有环链表，找第一个相交节点
//
// 【解题思路】
// 1. loop1 == loop2：转化为“环前相交问题”
// 2. loop1 != loop2：
//   - 判断环内是否可达
//   - 能走到说明相交，否则不相交
//
// 【时空复杂度】
// 时间：O(N)
// 空间：O(1)
func bothLoop(head1, loop1, head2, loop2 *Node) *Node {

	if loop1 == loop2 {
		cur1, cur2 := head1, head2
		n := 0

		for cur1 != loop1 {
			n++
			cur1 = cur1.next
		}
		for cur2 != loop2 {
			n--
			cur2 = cur2.next
		}

		if n > 0 {
			cur1 = head1
			cur2 = head2
		} else {
			cur1 = head2
			cur2 = head1
			n = -n
		}

		for n != 0 {
			cur1 = cur1.next
			n--
		}

		for cur1 != cur2 {
			cur1 = cur1.next
			cur2 = cur2.next
		}

		return cur1
	}

	cur := loop1.next
	for cur != loop1 {
		if cur == loop2 {
			return loop1
		}
		cur = cur.next
	}

	return nil
}

//
// =======================
// test
// =======================
//

func main() {

	// 1->2->3->4->5->6->7
	h1 := &Node{value: 1}
	h1.next = &Node{value: 2}
	h1.next.next = &Node{value: 3}
	h1.next.next.next = &Node{value: 4}
	h1.next.next.next.next = &Node{value: 5}
	h1.next.next.next.next.next = &Node{value: 6}
	h1.next.next.next.next.next.next = &Node{value: 7}

	// 0->9->8->6...
	h2 := &Node{value: 0}
	h2.next = &Node{value: 9}
	h2.next.next = &Node{value: 8}
	h2.next.next.next = h1.next.next.next.next.next

	fmt.Println(GetIntersectNode(h1, h2).value)

	// create cycle case
	h1 = &Node{value: 1}
	h1.next = &Node{value: 2}
	h1.next.next = &Node{value: 3}
	h1.next.next.next = &Node{value: 4}
	h1.next.next.next.next = &Node{value: 5}
	h1.next.next.next.next.next = &Node{value: 6}
	h1.next.next.next.next.next.next = &Node{value: 7}
	h1.next.next.next.next.next.next = h1.next.next.next // 7->4

	h2 = &Node{value: 0}
	h2.next = &Node{value: 9}
	h2.next.next = &Node{value: 8}
	h2.next.next.next = h1.next

	fmt.Println(GetIntersectNode(h1, h2).value)
}
