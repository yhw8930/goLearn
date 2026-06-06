package main

// 题目：给定链表头节点和目标值 num，删除链表中所有值等于 num 的节点。
// 头节点也可能需要删除，所以要先找到第一个不等于 num 的节点作为新头。
// 之后遍历剩余节点，用前驱节点决定是否跳过当前节点。
// 核心思路：删除节点本质是让前驱的 next 指向当前节点的下一个节点。
// 时间复杂度：O(N)。
// 空间复杂度：O(1)。

// 删除链表中所有等于 num 的节点
func removeValue(head *Node, num int) *Node {
	// 先跳过头部所有要删除的数
	for head != nil {
		if head.value != num {
			break
		}
		head = head.next
	}

	pre, cur := head, head
	for cur != nil {
		if cur.value == num {
			pre.next = cur.next
		} else {
			pre = cur
		}
		cur = cur.next
	}
	return head
}
