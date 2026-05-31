package main

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
