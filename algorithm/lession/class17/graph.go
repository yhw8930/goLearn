package main

// 题目：为图算法提供统一的点、边、图结构。
// Node 保存点编号、入度、出度、邻接点列表和从该点出发的边。
// Edge 保存边权以及 from/to 两端点，Graph 统一保存所有点和边。
// 核心思路：BFS、DFS、拓扑排序、最小生成树、最短路都可以复用这一套邻接结构。
// 时间复杂度：结构创建和插入按元素数量线性增长。
// 空间复杂度：O(V+E)。

// Node 表示图中的一个点。
type Node struct {
	Value int
	In    int
	Out   int
	Nexts []*Node
	Edges []*Edge
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

// Edge 表示一条有权边。
type Edge struct {
	Weight int
	From   *Node
	To     *Node
}

// Graph 使用点集和边集描述一张图。
type Graph struct {
	Nodes map[int]*Node
	Edges map[*Edge]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]*Node),
		Edges: make(map[*Edge]struct{}),
	}
}
