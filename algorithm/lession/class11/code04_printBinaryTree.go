package main

import (
	"fmt"
	"strings"
)

// 题目：把二叉树按更直观的横向结构打印出来，便于观察树形。
// 普通遍历只能看到节点访问顺序，不容易看出左右子树层级。
// 核心思路：先递归打印右子树，再打印当前节点，最后打印左子树，相当于把树逆时针旋转 90 度。
// 通过递归高度控制缩进，并用特殊标记区分当前节点来自左边还是右边。
// 时间复杂度：O(N)。
// 空间复杂度：O(H)。

type Code04_Node struct {
	value int
	left  *Code04_Node
	right *Code04_Node
}

func Code04_NewNode(data int) *Code04_Node {
	return &Code04_Node{value: data}
}

func Code04_printTree(head *Code04_Node) {
	fmt.Println("Binary Tree:")
	Code04_printInOrder(head, 0, "H", 17)
	fmt.Println()
}

func Code04_printInOrder(head *Code04_Node, height int, to string, length int) {
	if head == nil {
		return
	}
	Code04_printInOrder(head.right, height+1, "v", length)

	val := to + fmt.Sprintf("%d", head.value) + to
	lenM := len(val)
	lenL := (length - lenM) / 2
	lenR := length - lenM - lenL
	val = Code04_getSpace(lenL) + val + Code04_getSpace(lenR)

	fmt.Println(Code04_getSpace(height*length) + val)
	Code04_printInOrder(head.left, height+1, "^", length)
}

func Code04_getSpace(num int) string {
	return strings.Repeat(" ", num)
}

func main() {
	// 测试用例 1
	head := Code04_NewNode(1)
	head.left = Code04_NewNode(-222222222)
	head.right = Code04_NewNode(3)
	head.left.left = Code04_NewNode(-9223372036854775808)
	head.right.left = Code04_NewNode(55555555)
	head.right.right = Code04_NewNode(66)
	head.left.left.right = Code04_NewNode(777)
	Code04_printTree(head)

	// 测试用例 2
	head = Code04_NewNode(1)
	head.left = Code04_NewNode(2)
	head.right = Code04_NewNode(3)
	head.left.left = Code04_NewNode(4)
	head.right.left = Code04_NewNode(5)
	head.right.right = Code04_NewNode(6)
	head.left.left.right = Code04_NewNode(7)
	Code04_printTree(head)

	// 测试用例 3
	head = Code04_NewNode(1)
	head.left = Code04_NewNode(1)
	head.right = Code04_NewNode(1)
	head.left.left = Code04_NewNode(1)
	head.right.left = Code04_NewNode(1)
	head.right.right = Code04_NewNode(1)
	head.left.left.right = Code04_NewNode(1)
	Code04_printTree(head)
}
