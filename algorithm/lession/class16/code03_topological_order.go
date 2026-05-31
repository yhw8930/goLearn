package main

import "sort"

// DirectedGraphNode 是 LintCode 拓扑排序题常见的图节点结构。
type DirectedGraphNode struct {
	Label     int
	Neighbors []*DirectedGraphNode
}

// TopSortBFS 使用入度表做拓扑排序。
//
// 时间复杂度：O(V+E)。
// 空间复杂度：O(V)。
func TopSortBFS(graph []*DirectedGraphNode) []*DirectedGraphNode {
	indegreeMap := make(map[*DirectedGraphNode]int, len(graph))
	for _, cur := range graph {
		indegreeMap[cur] = 0
	}
	for _, cur := range graph {
		for _, next := range cur.Neighbors {
			indegreeMap[next]++
		}
	}
	zeroQueue := make([]*DirectedGraphNode, 0)
	for cur, indegree := range indegreeMap {
		if indegree == 0 {
			zeroQueue = append(zeroQueue, cur)
		}
	}
	ans := make([]*DirectedGraphNode, 0, len(graph))
	for len(zeroQueue) > 0 {
		cur := zeroQueue[0]
		zeroQueue = zeroQueue[1:]
		ans = append(ans, cur)
		for _, next := range cur.Neighbors {
			indegreeMap[next]--
			if indegreeMap[next] == 0 {
				zeroQueue = append(zeroQueue, next)
			}
		}
	}
	return ans
}

type deepRecord struct {
	node *DirectedGraphNode
	deep int
}

// TopSortDFSByDeep 使用“最大后继深度”排序。
// 对任意边 cur -> next，cur 的 deep 一定大于 next，因此 deep 降序就是拓扑序。
//
// 时间复杂度：O(V+E+VlogV)，DFS 扫图后还要排序。
// 空间复杂度：O(V)，缓存和递归栈。
func TopSortDFSByDeep(graph []*DirectedGraphNode) []*DirectedGraphNode {
	order := make(map[*DirectedGraphNode]deepRecord, len(graph))
	for _, cur := range graph {
		deepDFS(cur, order)
	}
	records := make([]deepRecord, 0, len(order))
	for _, record := range order {
		records = append(records, record)
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].deep > records[j].deep
	})
	ans := make([]*DirectedGraphNode, 0, len(records))
	for _, record := range records {
		ans = append(ans, record.node)
	}
	return ans
}

func deepDFS(cur *DirectedGraphNode, order map[*DirectedGraphNode]deepRecord) deepRecord {
	if record, ok := order[cur]; ok {
		return record
	}
	follow := 0
	for _, next := range cur.Neighbors {
		if nextRecord := deepDFS(next, order); nextRecord.deep > follow {
			follow = nextRecord.deep
		}
	}
	record := deepRecord{node: cur, deep: follow + 1}
	order[cur] = record
	return record
}

type nodesRecord struct {
	node  *DirectedGraphNode
	nodes int64
}

// TopSortDFSByReachableNodes 使用“可到达点次”排序。
// 点次 = 自己 + 所有后继路径贡献。对 DAG 来说，前驱点次会大于后继，
// 所以点次降序可以得到一种拓扑序。
//
// 时间复杂度：O(V+E+VlogV)。
// 空间复杂度：O(V)。
func TopSortDFSByReachableNodes(graph []*DirectedGraphNode) []*DirectedGraphNode {
	order := make(map[*DirectedGraphNode]nodesRecord, len(graph))
	for _, cur := range graph {
		nodesDFS(cur, order)
	}
	records := make([]nodesRecord, 0, len(order))
	for _, record := range order {
		records = append(records, record)
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].nodes > records[j].nodes
	})
	ans := make([]*DirectedGraphNode, 0, len(records))
	for _, record := range records {
		ans = append(ans, record.node)
	}
	return ans
}

func nodesDFS(cur *DirectedGraphNode, order map[*DirectedGraphNode]nodesRecord) nodesRecord {
	if record, ok := order[cur]; ok {
		return record
	}
	var nodes int64
	for _, next := range cur.Neighbors {
		nodes += nodesDFS(next, order).nodes
	}
	record := nodesRecord{node: cur, nodes: nodes + 1}
	order[cur] = record
	return record
}
