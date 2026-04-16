package main

import "fmt"

type code03Node struct {
	value int
	left  *code03Node
	right *code03Node
}

// code03Pre 非递归前序遍历（根-左-右）
//
// 【问题】
// 用栈实现二叉树前序遍历
//
// 【解题思路】
// 1. 根节点入栈
// 2. 弹出即访问
// 3. 先压右再压左（保证左先出）
//
// 【时空复杂度】
// 时间：O(N)
// 空间：O(N)
func code03Pre(head *code03Node) {
	fmt.Print("code03Pre-order: ")
	if head != nil {
		stack := []*code03Node{head}

		for len(stack) > 0 {
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Print(n.value, " ")

			if n.right != nil {
				stack = append(stack, n.right)
			}
			if n.left != nil {
				stack = append(stack, n.left)
			}
		}
	}
	fmt.Println()
}

// code03In 非递归中序遍历（左-根-右）
//
// 【问题】
// 用栈实现中序遍历
//
// 【解题思路】
// 1. 一路压左节点
// 2. 无左子树时弹出访问
// 3. 转向右子树重复
//
// 【时空复杂度】
// 时间：O(N)
// 空间：O(N)
func code03In(cur *code03Node) {
	fmt.Print("code03In-order: ")

	stack := []*code03Node{}

	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Print(cur.value, " ")
			cur = cur.right
		}
	}

	fmt.Println()
}

// Pos1 后序遍历（双栈实现）
//
// 【问题】
// 用两个栈实现后序遍历
//
// 【解题思路】
// 1. s1实现：根-右-左
// 2. 压入s2反转顺序
// 3. s2弹出即左-右-根
//
// 【时空复杂度】
// 时间：O(N)
// 空间：O(N)
func Pos1(head *code03Node) {
	fmt.Print("pos-order: ")

	if head == nil {
		fmt.Println()
		return
	}

	s1 := []*code03Node{head}
	s2 := []*code03Node{}

	for len(s1) > 0 {
		n := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]

		s2 = append(s2, n)

		if n.left != nil {
			s1 = append(s1, n.left)
		}
		if n.right != nil {
			s1 = append(s1, n.right)
		}
	}

	for i := len(s2) - 1; i >= 0; i-- {
		fmt.Print(s2[i].value, " ")
	}

	fmt.Println()
}

// Pos2 后序遍历（单栈实现）
//
// 【问题】
// 用一个栈实现后序遍历
//
// 【解题思路】
// 1. 记录上一次访问节点
// 2. 判断是否能从子树回退
// 3. 满足条件则弹出
//
// 【时空复杂度】
// 时间：O(N)
// 空间：O(N)
func Pos2(head *code03Node) {
	fmt.Print("pos-order: ")

	if head == nil {
		fmt.Println()
		return
	}

	stack := []*code03Node{head}
	var last *code03Node = nil

	for len(stack) > 0 {
		c := stack[len(stack)-1]

		if c.left != nil && last != c.left && last != c.right {
			stack = append(stack, c.left)
		} else if c.right != nil && last != c.right {
			stack = append(stack, c.right)
		} else {
			fmt.Print(c.value, " ")
			stack = stack[:len(stack)-1]
			last = c
		}
	}

	fmt.Println()
}

func main() {
	head := &code03Node{value: 1}
	head.left = &code03Node{value: 2}
	head.right = &code03Node{value: 3}
	head.left.left = &code03Node{value: 4}
	head.left.right = &code03Node{value: 5}
	head.right.left = &code03Node{value: 6}
	head.right.right = &code03Node{value: 7}

	code03Pre(head)
	fmt.Println("========")
	code03In(head)
	fmt.Println("========")
	Pos1(head)
	fmt.Println("========")
	Pos2(head)
	fmt.Println("========")
}
