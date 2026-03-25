package main

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
