package main

import "sort"

// 题目：给定无向带权图，使用 Kruskal 算法求最小生成树。
// 最小生成树要求连接所有点且总边权最小，并且不能形成环。
// 核心思路：把所有边按权重从小到大排序或放入小根堆。
// 依次尝试最小边，如果边两端点不在同一集合，就加入结果并合并集合；否则跳过避免成环。
// 前提：通常用于无向带权图。
// 时间复杂度：O(ElogE)。
// 空间复杂度：O(V+E)。

type unionFind struct {
	father map[*Node]*Node
	size   map[*Node]int
}

func newUnionFind(nodes map[int]*Node) *unionFind {
	uf := &unionFind{
		father: make(map[*Node]*Node, len(nodes)),
		size:   make(map[*Node]int, len(nodes)),
	}
	for _, node := range nodes {
		uf.father[node] = node
		uf.size[node] = 1
	}
	return uf
}

func (uf *unionFind) find(node *Node) *Node {
	path := make([]*Node, 0)
	for node != uf.father[node] {
		path = append(path, node)
		node = uf.father[node]
	}
	for _, cur := range path {
		uf.father[cur] = node
	}
	return node
}

func (uf *unionFind) isSameSet(a, b *Node) bool {
	return uf.find(a) == uf.find(b)
}

func (uf *unionFind) union(a, b *Node) {
	if a == nil || b == nil {
		return
	}
	af, bf := uf.find(a), uf.find(b)
	if af == bf {
		return
	}
	if uf.size[af] <= uf.size[bf] {
		uf.father[af] = bf
		uf.size[bf] += uf.size[af]
		delete(uf.size, af)
	} else {
		uf.father[bf] = af
		uf.size[af] += uf.size[bf]
		delete(uf.size, bf)
	}
}

// KruskalMST 求无向图的最小生成森林。
// 核心：所有边按权值从小到大尝试加入；如果边两端已经在同一集合，
// 加入会形成环，跳过；否则加入答案并合并集合。
//
// 注意：用当前 Graph 表达无向图时，通常需要把无向边拆成两条方向相反的边。
//
// 时间复杂度：O(ElogE)，主要来自边排序。
// 空间复杂度：O(V+E)。
func KruskalMST(graph *Graph) map[*Edge]struct{} {
	if graph == nil {
		return nil
	}
	uf := newUnionFind(graph.Nodes)
	edges := make([]*Edge, 0, len(graph.Edges))
	for edge := range graph.Edges {
		edges = append(edges, edge)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})
	ans := make(map[*Edge]struct{})
	for _, edge := range edges {
		if !uf.isSameSet(edge.From, edge.To) {
			ans[edge] = struct{}{}
			uf.union(edge.From, edge.To)
		}
	}
	return ans
}
