package main

// 题目：把 N 叉树编码成二叉树，并能从二叉树还原回 N 叉树。
// N 叉树每个节点可能有多个孩子，而二叉树每个节点只有 left 和 right。
// 核心思路：用 left 指向第一个孩子，用 right 串起兄弟节点，也就是“左孩子右兄弟”表示法。
// 解码时沿 left 找到孩子链表开头，再不断沿 right 收集所有兄弟节点。
// 时间复杂度：O(N)。
// 空间复杂度：O(H)，H 为递归深度。

// 提交时不要提交这个类
type Code03Node struct {
	val      int
	children []*Code03Node
}

// 提交时不要提交这个类
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// 只提交这个类即可
type Codec struct{}

// Encodes an n-ary tree to a binary tree.
func (c *Codec) encode(root *Code03Node) *TreeNode {
	if root == nil {
		return nil
	}
	head := &TreeNode{val: root.val}
	head.left = c.en(root.children)
	return head
}

func (c *Codec) en(children []*Code03Node) *TreeNode {
	var head *TreeNode
	var cur *TreeNode
	for _, child := range children {
		tNode := &TreeNode{val: child.val}
		if head == nil {
			head = tNode
		} else {
			cur.right = tNode
		}
		cur = tNode
		cur.left = c.en(child.children)
	}
	return head
}

// Decodes your binary tree to an n-ary tree.
func (c *Codec) decode(root *TreeNode) *Code03Node {
	if root == nil {
		return nil
	}
	return &Code03Node{
		val:      root.val,
		children: c.de(root.left),
	}
}

func (c *Codec) de(root *TreeNode) []*Code03Node {
	children := []*Code03Node{}
	for root != nil {
		cur := &Code03Node{
			val:      root.val,
			children: c.de(root.left),
		}
		children = append(children, cur)
		root = root.right
	}
	return children
}
