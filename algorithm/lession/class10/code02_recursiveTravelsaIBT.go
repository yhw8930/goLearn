package main

import "fmt"

// 题目：递归实现二叉树先序、中序、后序遍历。
// 三种遍历都遵循递归处理左子树和右子树，只是访问当前节点的时机不同。
// 核心思路：先序是在递归左右子树前访问当前节点，中序是在左右之间访问，后序是在左右之后访问。
// 递归终止条件是当前节点为空，递归调用栈天然保存从根到当前节点的路径。
// 时间复杂度：O(N)。
// 空间复杂度：O(H)，H 为树高。

type code02Node struct {
	value int
	left  *code02Node
	right *code02Node
}

// Pre 前序遍历（根-左-右）
//
// 【问题】
// 递归方式输出二叉树所有节点（前序）
//
// 【解题思路】
// 1. 先访问根节点
// 2. 递归左子树
// 3. 递归右子树
//
// 【时空复杂度】
// 时间：O(N)
// 空间：O(H) H为树高（递归栈）
func Pre(head *code02Node) {
	if head == nil {
		return
	}
	fmt.Println(head.value)
	Pre(head.left)
	Pre(head.right)
}

// In 中序遍历（左-根-右）
//
// 【问题】
// 递归方式输出二叉树所有节点（中序）
//
// 【解题思路】
// 1. 递归左子树
// 2. 访问根节点
// 3. 递归右子树
//
// 【时空复杂度】
// 时间：O(N)
// 空间：O(H)
func In(head *code02Node) {
	if head == nil {
		return
	}
	In(head.left)
	fmt.Println(head.value)
	In(head.right)
}

// Pos 后序遍历（左-右-根）
//
// 【问题】
// 递归方式输出二叉树所有节点（后序）
//
// 【解题思路】
// 1. 递归左子树
// 2. 递归右子树
// 3. 访问根节点
//
// 【时空复杂度】
// 时间：O(N)
// 空间：O(H)
func Pos(head *code02Node) {
	if head == nil {
		return
	}
	Pos(head.left)
	Pos(head.right)
	fmt.Println(head.value)
}

func main() {
	head := &code02Node{value: 1}
	head.left = &code02Node{value: 2}
	head.right = &code02Node{value: 3}
	head.left.left = &code02Node{value: 4}
	head.left.right = &code02Node{value: 5}
	head.right.left = &code02Node{value: 6}
	head.right.right = &code02Node{value: 7}

	Pre(head)
	fmt.Println("========")
	In(head)
	fmt.Println("========")
	Pos(head)
	fmt.Println("========")
}
