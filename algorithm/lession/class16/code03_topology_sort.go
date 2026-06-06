package main

// 题目：给定有向无环图，返回一个拓扑排序结果。
// 拓扑序要求每条有向边 from->to 中，from 必须排在 to 前面。
// 核心思路：统计每个节点入度，所有入度为 0 的节点都可以作为当前输出节点。
// 输出一个节点后，把它指向的后继入度减一，新变成 0 的节点继续入队。
// 前提：有向无环图。
// 时间复杂度：O(V+E)。
// 空间复杂度：O(V)。

// SortedTopology 对有向无环图做拓扑排序。
// 核心：先收集所有入度为 0 的点；每弹出一个点，就“删除”它的出边，
// 让后继点的剩余入度减 1，新的 0 入度点继续入队。
//
// 时间复杂度：O(V+E)。
// 空间复杂度：O(V)。
func SortedTopology(graph *Graph) []*Node {
	if graph == nil {
		return nil
	}
	inMap := make(map[*Node]int, len(graph.Nodes))
	zeroQueue := make([]*Node, 0)
	for _, node := range graph.Nodes {
		inMap[node] = node.In
		if node.In == 0 {
			zeroQueue = append(zeroQueue, node)
		}
	}
	ans := make([]*Node, 0, len(graph.Nodes))
	for len(zeroQueue) > 0 {
		cur := zeroQueue[0]
		zeroQueue = zeroQueue[1:]
		ans = append(ans, cur)
		for _, next := range cur.Nexts {
			inMap[next]--
			if inMap[next] == 0 {
				zeroQueue = append(zeroQueue, next)
			}
		}
	}
	return ans
}
