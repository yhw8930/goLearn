package main

import "fmt"

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
