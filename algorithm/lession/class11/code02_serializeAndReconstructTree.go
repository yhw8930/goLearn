package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

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
