package main

import "fmt"

// 题目：给定学生认识关系矩阵 M，M[i][j]=1 表示 i 和 j 直接认识，求朋友圈数量。
// 朋友圈是传递关系：A 认识 B，B 认识 C，则 A/B/C 属于同一个朋友圈。
// 核心思路：初始每个人是一个集合，遍历矩阵上三角，把直接认识的两个人所在集合合并。
// 最终并查集剩余集合数量就是朋友圈数量。
// 时间复杂度：O(N^2)。
// 空间复杂度：O(N)。

// FriendCircles 结构体封装
type FriendCircles struct {
	parent []int
	size   []int
	help   []int
	sets   int
}

// NewFriendCircles 初始化并查集
func NewFriendCircles(N int) *FriendCircles {
	fc := &FriendCircles{
		parent: make([]int, N),
		size:   make([]int, N),
		help:   make([]int, N),
		sets:   N,
	}
	for i := 0; i < N; i++ {
		fc.parent[i] = i
		fc.size[i] = 1
	}
	return fc
}

// find 找代表节点 + 路径压缩
func (fc *FriendCircles) find(i int) int {
	hi := 0
	// 找到代表节点
	for i != fc.parent[i] {
		fc.help[hi] = i
		hi++
		i = fc.parent[i]
	}
	// 路径压缩
	for hi--; hi >= 0; hi-- {
		fc.parent[fc.help[hi]] = i
	}
	return i
}

// union 合并两个集合
func (fc *FriendCircles) union(i, j int) {
	f1 := fc.find(i)
	f2 := fc.find(j)

	if f1 != f2 {
		// 小集合挂到大集合
		if fc.size[f1] >= fc.size[f2] {
			fc.size[f1] += fc.size[f2]
			fc.parent[f2] = f1
		} else {
			fc.size[f2] += fc.size[f1]
			fc.parent[f1] = f2
		}
		fc.sets--
	}
}

// Sets 获取当前集合数量
func (fc *FriendCircles) Sets() int {
	return fc.sets
}

// findCircleNum 主函数（LeetCode 547）
func findCircleNum(M [][]int) int {
	N := len(M)
	fc := NewFriendCircles(N)

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if M[i][j] == 1 {
				fc.union(i, j)
			}
		}
	}
	return fc.Sets()
}

// 测试
func main() {
	M := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(findCircleNum(M)) // 输出 2 ✅
}
