package main

import (
	"container/heap"
	"math"
)

// 题目：给定无向带权图，使用 Prim 算法求最小生成树。
// Prim 从某个点所在连通块出发，每次选择一条连接已访问点和未访问点的最小边。
// 核心思路：把已解锁节点的所有边放入小根堆，弹出最小边时检查它通向的点是否已访问。
// 如果是新点，就把边加入结果，并继续解锁该新点的所有出边。
// 前提：通常用于无向带权图。
// 时间复杂度：邻接表堆实现约 O(ElogE)，矩阵版本 O(V^2)。
// 空间复杂度：O(V+E)。

type edgeHeap []*Edge

func (h edgeHeap) Len() int           { return len(h) }
func (h edgeHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h edgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *edgeHeap) Push(x any) {
	*h = append(*h, x.(*Edge))
}

func (h *edgeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// PrimMST 求无向图的最小生成森林。
// 核心：从任意未访问点开始，把它的边解锁进小根堆；
// 每次取权值最小且能到达新点的边，并继续解锁新点的边。
//
// 时间复杂度：O(ElogE)。
// 空间复杂度：O(V+E)。
func PrimMST(graph *Graph) map[*Edge]struct{} {
	if graph == nil {
		return nil
	}
	h := &edgeHeap{}
	heap.Init(h)
	visited := make(map[*Node]struct{}, len(graph.Nodes))
	ans := make(map[*Edge]struct{})
	for _, node := range graph.Nodes {
		if _, ok := visited[node]; ok {
			continue
		}
		visited[node] = struct{}{}
		for _, edge := range node.Edges {
			heap.Push(h, edge)
		}
		for h.Len() > 0 {
			edge := heap.Pop(h).(*Edge)
			to := edge.To
			if _, ok := visited[to]; ok {
				continue
			}
			visited[to] = struct{}{}
			ans[edge] = struct{}{}
			for _, nextEdge := range to.Edges {
				heap.Push(h, nextEdge)
			}
		}
	}
	return ans
}

// PrimMatrix 使用邻接矩阵求连通图最小生成树的权值和。
// graph[i][j] 表示 i 到 j 的距离；math.MaxInt 表示无路。
//
// 时间复杂度：O(V^2)。
// 空间复杂度：O(V)。
func PrimMatrix(graph [][]int) int {
	size := len(graph)
	if size == 0 {
		return 0
	}
	distances := make([]int, size)
	visited := make([]bool, size)
	visited[0] = true
	copy(distances, graph[0])
	sum := 0
	for i := 1; i < size; i++ {
		minPath := math.MaxInt
		minIndex := -1
		for j := 0; j < size; j++ {
			if !visited[j] && distances[j] < minPath {
				minPath = distances[j]
				minIndex = j
			}
		}
		if minIndex == -1 {
			return sum
		}
		visited[minIndex] = true
		sum += minPath
		for j := 0; j < size; j++ {
			if !visited[j] && graph[minIndex][j] < distances[j] {
				distances[j] = graph[minIndex][j]
			}
		}
	}
	return sum
}
