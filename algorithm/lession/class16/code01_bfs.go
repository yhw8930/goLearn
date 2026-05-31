package main

import "fmt"

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
