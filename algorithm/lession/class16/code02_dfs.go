package main

import "fmt"

// 题目：从图中某个起点开始做深度优先遍历，访问所有可达节点。
// 图可能有环，需要 set 记录已访问节点。
// 核心思路：用栈模拟递归，优先沿着一条未访问边走到底。
// 当发现未访问邻居时，把当前节点和邻居都压回栈，先处理邻居，从而保持深度优先效果。
// 时间复杂度：O(V+E)。
// 空间复杂度：O(V)。

// DFS 从 node 出发进行深度优先遍历，使用显式栈模拟递归。
// 遇到一个没访问过的邻居时，先把当前点压回栈，再压入邻居，
// 这样下一轮会沿着这条路一直向深处走。
//
// 时间复杂度：O(V+E)。
// 空间复杂度：O(V)，栈和 visited 最多保存所有可达点。
func DFS(node *Node) {
	if node == nil {
		return
	}
	stack := []*Node{node}
	visited := map[*Node]struct{}{node: {}}
	fmt.Println(node.Value)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, next := range cur.Nexts {
			if _, ok := visited[next]; ok {
				continue
			}
			stack = append(stack, cur, next)
			visited[next] = struct{}{}
			fmt.Println(next.Value)
			break
		}
	}
}
