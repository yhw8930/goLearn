package main

import "math"

// 题目：给定边权非负的图和起点 from，求 from 到所有可达节点的最短距离。
// Dijkstra 的前提是边权不能为负，否则已经确定的最短距离可能被后续路径推翻。
// 核心思路：维护起点到各节点的当前最短估计，每次选择未确定节点中距离最小的点。
// 确定该点后，用它的出边去松弛邻居距离；堆优化版本用小根堆加速选择最小距离点。
// 前提：边权非负。
// 时间复杂度：朴素版 O(V^2+E)，堆优化版 O((V+E)logV)。
// 空间复杂度：O(V)。

// Dijkstra1 求 from 到所有可达点的最短距离，要求边权非负。
// 朴素版本每轮从 distanceMap 中线性找一个未确定的最小距离点。
//
// 时间复杂度：O(V^2+E)。
// 空间复杂度：O(V)。
func Dijkstra1(from *Node) map[*Node]int {
	if from == nil {
		return nil
	}
	distanceMap := map[*Node]int{from: 0}
	selected := make(map[*Node]struct{})
	minNode := getMinDistanceAndUnselectedNode(distanceMap, selected)
	for minNode != nil {
		distance := distanceMap[minNode]
		for _, edge := range minNode.Edges {
			to := edge.To
			newDistance := distance + edge.Weight
			if oldDistance, ok := distanceMap[to]; !ok || newDistance < oldDistance {
				distanceMap[to] = newDistance
			}
		}
		selected[minNode] = struct{}{}
		minNode = getMinDistanceAndUnselectedNode(distanceMap, selected)
	}
	return distanceMap
}

func getMinDistanceAndUnselectedNode(distanceMap map[*Node]int, selected map[*Node]struct{}) *Node {
	var minNode *Node
	minDistance := math.MaxInt
	for node, distance := range distanceMap {
		if _, ok := selected[node]; !ok && distance < minDistance {
			minNode = node
			minDistance = distance
		}
	}
	return minNode
}

type nodeRecord struct {
	node     *Node
	distance int
}

type nodeHeap struct {
	nodes        []*Node
	heapIndexMap map[*Node]int
	distanceMap  map[*Node]int
	size         int
}

func newNodeHeap(capacity int) *nodeHeap {
	return &nodeHeap{
		nodes:        make([]*Node, capacity),
		heapIndexMap: make(map[*Node]int),
		distanceMap:  make(map[*Node]int),
	}
}

func (h *nodeHeap) isEmpty() bool {
	return h.size == 0
}

func (h *nodeHeap) addOrUpdateOrIgnore(node *Node, distance int) {
	if h.inHeap(node) {
		if distance < h.distanceMap[node] {
			h.distanceMap[node] = distance
			h.insertHeapify(h.heapIndexMap[node])
		}
		return
	}
	if !h.isEntered(node) {
		h.nodes[h.size] = node
		h.heapIndexMap[node] = h.size
		h.distanceMap[node] = distance
		h.insertHeapify(h.size)
		h.size++
	}
}

func (h *nodeHeap) pop() nodeRecord {
	record := nodeRecord{node: h.nodes[0], distance: h.distanceMap[h.nodes[0]]}
	h.swap(0, h.size-1)
	h.heapIndexMap[h.nodes[h.size-1]] = -1
	delete(h.distanceMap, h.nodes[h.size-1])
	h.nodes[h.size-1] = nil
	h.size--
	h.heapify(0)
	return record
}

func (h *nodeHeap) insertHeapify(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.distanceMap[h.nodes[index]] >= h.distanceMap[h.nodes[parent]] {
			break
		}
		h.swap(index, parent)
		index = parent
	}
}

func (h *nodeHeap) heapify(index int) {
	left := index*2 + 1
	for left < h.size {
		smallest := left
		if left+1 < h.size && h.distanceMap[h.nodes[left+1]] < h.distanceMap[h.nodes[left]] {
			smallest = left + 1
		}
		if h.distanceMap[h.nodes[index]] <= h.distanceMap[h.nodes[smallest]] {
			break
		}
		h.swap(index, smallest)
		index = smallest
		left = index*2 + 1
	}
}

func (h *nodeHeap) isEntered(node *Node) bool {
	_, ok := h.heapIndexMap[node]
	return ok
}

func (h *nodeHeap) inHeap(node *Node) bool {
	index, ok := h.heapIndexMap[node]
	return ok && index != -1
}

func (h *nodeHeap) swap(i, j int) {
	h.heapIndexMap[h.nodes[i]] = j
	h.heapIndexMap[h.nodes[j]] = i
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

// Dijkstra2 是加强堆版本，要求边权非负。
// addOrUpdateOrIgnore 实现三种状态：
// 1. 在堆中：发现更短距离就更新并向上调整。
// 2. 从没进过堆：加入堆。
// 3. 已经弹出过：最短距离已确定，忽略。
//
// 时间复杂度：O((V+E)logV)。
// 空间复杂度：O(V)。
func Dijkstra2(head *Node, size int) map[*Node]int {
	if head == nil || size <= 0 {
		return nil
	}
	h := newNodeHeap(size)
	h.addOrUpdateOrIgnore(head, 0)
	ans := make(map[*Node]int)
	for !h.isEmpty() {
		record := h.pop()
		cur := record.node
		distance := record.distance
		for _, edge := range cur.Edges {
			h.addOrUpdateOrIgnore(edge.To, edge.Weight+distance)
		}
		ans[cur] = distance
	}
	return ans
}
