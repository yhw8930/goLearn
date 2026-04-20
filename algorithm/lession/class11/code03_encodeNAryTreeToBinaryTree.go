package main

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
