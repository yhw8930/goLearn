package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 题目：把二叉树序列化成线性结构，并能根据该结构重建出完全相同的树。
// 只记录非空节点值无法唯一恢复树结构，所以必须显式记录空节点。
// 核心思路：先序、后序或层序序列化都可以，只要反序列化时使用同一套顺序和空节点标记。
// 重建时按序列逐项消费，遇到空标记返回 nil，遇到值则创建节点并继续构造子树。
// 时间复杂度：O(N)。
// 空间复杂度：O(N)，用于保存序列和递归/队列结构。

// 二叉树节点
type Code02Node struct {
	value int
	left  *Code02Node
	right *Code02Node
}

func NewNode(data int) *Code02Node {
	return &Code02Node{value: data}
}

/*
 * 二叉树可以通过先序、后序或者按层遍历的方式序列化和反序列化，
 * 以下代码全部实现了。
 * 但是，二叉树无法通过中序遍历的方式实现序列化和反序列化
 * 因为不同的两棵树，可能得到同样的中序序列，即便补了空位置也可能一样。
 * 比如如下两棵树
 *         __2
 *        /
 *       1
 *       和
 *       1__
 *          \
 *           2
 * 补足空位置的中序遍历结果都是{ null, 1, null, 2, null}
 * */

// 先序序列化
// preSerial 使用先序方式序列化。
// 它按头、左、右顺序记录节点，并用空标记保留结构。
// 时间复杂度：O(N)。
// 空间复杂度：O(N)，序列保存所有节点和空标记。
func preSerial(head *Code02Node) *list.List {
	ans := list.New()
	pres(head, ans)
	return ans
}

func pres(head *Code02Node, ans *list.List) {
	if head == nil {
		ans.PushBack(nil)
	} else {
		ans.PushBack(strconv.Itoa(head.value))
		pres(head.left, ans)
		pres(head.right, ans)
	}
}

// 中序序列化
// inSerial 使用中序方式序列化。
// 中序序列本身不适合单独反序列化，主要用于展示遍历序列化形式。
// 时间复杂度：O(N)。
// 空间复杂度：O(N)。
func inSerial(head *Code02Node) *list.List {
	ans := list.New()
	ins(head, ans)
	return ans
}

func ins(head *Code02Node, ans *list.List) {
	if head == nil {
		ans.PushBack(nil)
	} else {
		ins(head.left, ans)
		ans.PushBack(strconv.Itoa(head.value))
		ins(head.right, ans)
	}
}

// 后序序列化
// posSerial 使用后序方式序列化。
// 它按左、右、头顺序记录节点，并保留空节点。
// 时间复杂度：O(N)。
// 空间复杂度：O(N)。
func posSerial(head *Code02Node) *list.List {
	ans := list.New()
	poss(head, ans)
	return ans
}

func poss(head *Code02Node, ans *list.List) {
	if head == nil {
		ans.PushBack(nil)
	} else {
		poss(head.left, ans)
		poss(head.right, ans)
		ans.PushBack(strconv.Itoa(head.value))
	}
}

// 先序反序列化
// buildByPreQueue 根据先序队列反序列化。
// 每消费一个值就递归构造当前节点的左子树和右子树。
// 时间复杂度：O(N)。
// 空间复杂度：O(H)，递归深度；输入队列 O(N)。
func buildByPreQueue(prelist *list.List) *Code02Node {
	if prelist == nil || prelist.Len() == 0 {
		return nil
	}
	return preb(prelist)
}

func preb(prelist *list.List) *Code02Node {
	e := prelist.Front()
	prelist.Remove(e)
	value, ok := e.Value.(string)
	if !ok {
		return nil
	}
	val, _ := strconv.Atoi(value)
	head := NewNode(val)
	head.left = preb(prelist)
	head.right = preb(prelist)
	return head
}

// 后序反序列化
// buildByPosQueue 根据后序队列反序列化。
// 后序反序列化通常先反转为栈，再按头、右、左的顺序恢复。
// 时间复杂度：O(N)。
// 空间复杂度：O(N)，栈和递归消耗。
func buildByPosQueue(poslist *list.List) *Code02Node {
	if poslist == nil || poslist.Len() == 0 {
		return nil
	}
	stack := list.New()
	for poslist.Len() > 0 {
		e := poslist.Front()
		poslist.Remove(e)
		stack.PushBack(e.Value)
	}
	return posb(stack)
}

func posb(posstack *list.List) *Code02Node {
	e := posstack.Back()
	posstack.Remove(e)
	value, ok := e.Value.(string)
	if !ok {
		return nil
	}
	val, _ := strconv.Atoi(value)
	head := NewNode(val)
	head.right = posb(posstack)
	head.left = posb(posstack)
	return head
}

// 按层序列化
// levelSerial 使用层序方式序列化。
// 队列按层展开节点，同时记录空孩子保证结构不丢失。
// 时间复杂度：O(N)。
// 空间复杂度：O(N)。
func levelSerial(head *Code02Node) *list.List {
	ans := list.New()
	if head == nil {
		ans.PushBack(nil)
	} else {
		ans.PushBack(strconv.Itoa(head.value))
		queue := list.New()
		queue.PushBack(head)
		for queue.Len() > 0 {
			e := queue.Front()
			queue.Remove(e)
			head = e.Value.(*Code02Node)
			if head.left != nil {
				ans.PushBack(strconv.Itoa(head.left.value))
				queue.PushBack(head.left)
			} else {
				ans.PushBack(nil)
			}
			if head.right != nil {
				ans.PushBack(strconv.Itoa(head.right.value))
				queue.PushBack(head.right)
			} else {
				ans.PushBack(nil)
			}
		}
	}
	return ans
}

// 按层反序列化
// buildByLevelQueue 根据层序队列反序列化。
// 它按队列顺序为每个弹出的父节点依次生成左、右孩子。
// 时间复杂度：O(N)。
// 空间复杂度：O(N)，队列保存待连接节点。
func buildByLevelQueue(levelList *list.List) *Code02Node {
	if levelList == nil || levelList.Len() == 0 {
		return nil
	}
	head := generateNode(levelList)
	queue := list.New()
	if head != nil {
		queue.PushBack(head)
	}
	var code02Node *Code02Node
	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		code02Node = e.Value.(*Code02Node)
		code02Node.left = generateNode(levelList)
		code02Node.right = generateNode(levelList)
		if code02Node.left != nil {
			queue.PushBack(code02Node.left)
		}
		if code02Node.right != nil {
			queue.PushBack(code02Node.right)
		}
	}
	return head
}

func generateNode(lst *list.List) *Code02Node {
	e := lst.Front()
	lst.Remove(e)
	val, ok := e.Value.(string)
	if !ok {
		return nil
	}
	v, _ := strconv.Atoi(val)
	return NewNode(v)
}

// for test
func generateRandomBST(maxLevel int, maxValue int) *Code02Node {
	rand.Seed(time.Now().UnixNano())
	return generate(1, maxLevel, maxValue)
}

// for test
func generate(level int, maxLevel int, maxValue int) *Code02Node {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}
	head := NewNode(rand.Intn(maxValue))
	head.left = generate(level+1, maxLevel, maxValue)
	head.right = generate(level+1, maxLevel, maxValue)
	return head
}

// for test
func isSameValueStructure(head1 *Code02Node, head2 *Code02Node) bool {
	if head1 == nil && head2 != nil {
		return false
	}
	if head1 != nil && head2 == nil {
		return false
	}
	if head1 == nil && head2 == nil {
		return true
	}
	if head1.value != head2.value {
		return false
	}
	return isSameValueStructure(head1.left, head2.left) && isSameValueStructure(head1.right, head2.right)
}

// for test
func printTree(head *Code02Node) {
	fmt.Println("Binary Tree:")
	printInOrder(head, 0, "H", 17)
	fmt.Println()
}

func printInOrder(head *Code02Node, height int, to string, length int) {
	if head == nil {
		return
	}
	printInOrder(head.right, height+1, "v", length)
	val := to + strconv.Itoa(head.value) + to
	lenM := len(val)
	lenL := (length - lenM) / 2
	lenR := length - lenM - lenL
	val = getSpace(lenL) + val + getSpace(lenR)
	fmt.Println(getSpace(height*length) + val)
	printInOrder(head.left, height+1, "^", length)
}

func getSpace(num int) string {
	space := " "
	buf := ""
	for i := 0; i < num; i++ {
		buf += space
	}
	return buf
}

func main() {
	maxLevel := 5
	maxValue := 100
	testTimes := 10000
	fmt.Println("test begin")
	for i := 0; i < testTimes; i++ {
		head := generateRandomBST(maxLevel, maxValue)
		pre := preSerial(head)
		pos := posSerial(head)
		level := levelSerial(head)
		preBuild := buildByPreQueue(pre)
		posBuild := buildByPosQueue(pos)
		levelBuild := buildByLevelQueue(level)
		if !isSameValueStructure(preBuild, posBuild) || !isSameValueStructure(posBuild, levelBuild) {
			fmt.Println("Oops!")
		}
	}
	fmt.Println("test finish!")
}
