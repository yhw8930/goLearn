package main

import (
	"container/list"
)

// Node 包装节点，对应 Java 的 Node<V>
type Node struct {
	Value interface{}
}

// UnionFind 并查集结构
type UnionFind struct {
	nodes   map[interface{}]*Node // 值 -> 节点
	parents map[*Node]*Node       // 父节点映射
	sizeMap map[*Node]int         // 代表节点 -> 集合大小
}

// NewUnionFind 构造函数
func NewUnionFind(values []interface{}) *UnionFind {
	uf := &UnionFind{
		nodes:   make(map[interface{}]*Node),
		parents: make(map[*Node]*Node),
		sizeMap: make(map[*Node]int),
	}

	// 初始化每个节点：自己是自己的父节点，size=1
	for _, v := range values {
		node := &Node{Value: v}
		uf.nodes[v] = node
		uf.parents[node] = node
		uf.sizeMap[node] = 1
	}
	return uf
}

// FindFather 找代表节点 + 路径压缩（核心）
func (uf *UnionFind) FindFather(cur *Node) *Node {
	path := list.New() // 用链表模拟栈

	// 一直往上找，直到找到自己是自己父节点的节点
	for cur != uf.parents[cur] {
		path.PushBack(cur)
		cur = uf.parents[cur]
	}

	// 路径压缩：沿途所有节点直接指向代表节点
	for path.Len() > 0 {
		node := path.Remove(path.Back()).(*Node)
		uf.parents[node] = cur
	}

	return cur
}

// IsSameSet 判断 a 和 b 是否在同一个集合
func (uf *UnionFind) IsSameSet(a, b interface{}) bool {
	nodeA := uf.nodes[a]
	nodeB := uf.nodes[b]
	if nodeA == nil || nodeB == nil {
		return false
	}
	return uf.FindFather(nodeA) == uf.FindFather(nodeB)
}

// Union 合并 a 和 b 所在的集合
func (uf *UnionFind) Union(a, b interface{}) {
	nodeA := uf.nodes[a]
	nodeB := uf.nodes[b]
	if nodeA == nil || nodeB == nil {
		return
	}

	headA := uf.FindFather(nodeA)
	headB := uf.FindFather(nodeB)

	if headA != headB {
		sizeA := uf.sizeMap[headA]
		sizeB := uf.sizeMap[headB]

		// 小集合挂到大集合下面
		var big, small *Node
		if sizeA >= sizeB {
			big = headA
			small = headB
		} else {
			big = headB
			small = headA
		}

		// 改父节点
		uf.parents[small] = big
		uf.sizeMap[big] = sizeA + sizeB
		delete(uf.sizeMap, small)
	}
}

// Sets 返回当前集合数量
func (uf *UnionFind) Sets() int {
	return len(uf.sizeMap)
}
