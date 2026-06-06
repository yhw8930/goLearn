package main

import "fmt"

// 题目：从图中某个起点开始做宽度优先遍历，打印或访问所有可达节点。
// 图可能有环，所以必须记录已经访问过的节点，避免重复入队和死循环。
// 核心思路：队列保存待扩展节点，先访问离起点距离更近的节点。
// 每弹出一个节点，就把它所有未访问过的邻居加入队列并标记为已访问。
// 时间复杂度：O(V+E)。
// 空间复杂度：O(V)。

// BFS 从 start 出发进行宽度优先遍历。
// 核心：队列保证按层推进，visited 保证有环图中每个点只进队一次。
//
// 时间复杂度：O(V+E)，从 start 能到达的点和边都会被扫描一次。
// 空间复杂度：O(V)，队列和 visited 最多保存所有可达点。
func BFS(start *Node) {
	if start == nil {
		return
	}
	queue := []*Node{start}
	visited := map[*Node]struct{}{start: {}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Println(cur.Value)
		for _, next := range cur.Nexts {
			if _, ok := visited[next]; ok {
				continue
			}
			visited[next] = struct{}{}
			queue = append(queue, next)
		}
	}
}
