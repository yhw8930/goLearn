package main

import (
	"fmt"
	"strconv"
)

// ================================================================================
// 解法一：数组实现的并查集（速度快，适合 m*n 不大的情况）
// ================================================================================
type UnionFind11 struct {
	parent []int
	size   []int
	help   []int
	row    int
	col    int
	sets   int
}

func NewUnionFind11(m, n int) *UnionFind11 {
	length := m * n
	return &UnionFind11{
		parent: make([]int, length),
		size:   make([]int, length),
		help:   make([]int, length),
		row:    m,
		col:    n,
		sets:   0,
	}
}

// index 将二维坐标 (r,c) 转为一维索引
func (uf *UnionFind11) index(r, c int) int {
	return r*uf.col + c
}

// find 查找代表节点 + 路径压缩
func (uf *UnionFind11) find(i int) int {
	hi := 0
	for i != uf.parent[i] {
		uf.help[hi] = i
		hi++
		i = uf.parent[i]
	}
	// 路径压缩
	for hi--; hi >= 0; hi-- {
		uf.parent[uf.help[hi]] = i
	}
	return i
}

// union 合并两个位置
func (uf *UnionFind11) union(r1, c1, r2, c2 int) {
	// 越界直接返回
	if r1 < 0 || r1 >= uf.row || r2 < 0 || r2 >= uf.row || c1 < 0 || c1 >= uf.col || c2 < 0 || c2 >= uf.col {
		return
	}

	i1 := uf.index(r1, c1)
	i2 := uf.index(r2, c2)

	// 如果任意一个位置不是陆地，不合并
	if uf.size[i1] == 0 || uf.size[i2] == 0 {
		return
	}

	f1 := uf.find(i1)
	f2 := uf.find(i2)

	if f1 != f2 {
		// 小挂大
		if uf.size[f1] >= uf.size[f2] {
			uf.size[f1] += uf.size[f2]
			uf.parent[f2] = f1
		} else {
			uf.size[f2] += uf.size[f1]
			uf.parent[f1] = f2
		}
		uf.sets--
	}
}

// Connect 动态添加陆地 (r,c)，返回当前岛屿数量
func (uf *UnionFind11) Connect(r, c int) int {
	index := uf.index(r, c)
	// 只有当前是水（size=0），才进行初始化
	if uf.size[index] == 0 {
		uf.parent[index] = index
		uf.size[index] = 1
		uf.sets++ // 新增岛屿

		// 尝试与上下左右合并
		uf.union(r-1, c, r, c)
		uf.union(r+1, c, r, c)
		uf.union(r, c-1, r, c)
		uf.union(r, c+1, r, c)
	}
	return uf.sets
}

func numIslands21(m, n int, positions [][]int) []int {
	uf := NewUnionFind11(m, n)
	var ans []int
	for _, pos := range positions {
		ans = append(ans, uf.Connect(pos[0], pos[1]))
	}
	return ans
}

// ================================================================================
// 解法二：HashMap 优化版（适合地图超大，但操作数少的情况）
// ================================================================================
type UnionFind22 struct {
	parent map[string]string
	size   map[string]int
	help   []string
	sets   int
}

func NewUnionFind22() *UnionFind22 {
	return &UnionFind22{
		parent: make(map[string]string),
		size:   make(map[string]int),
		help:   []string{},
		sets:   0,
	}
}

func (uf *UnionFind22) find(cur string) string {
	// 找到代表节点
	for cur != uf.parent[cur] {
		uf.help = append(uf.help, cur)
		cur = uf.parent[cur]
	}
	// 路径压缩
	for _, s := range uf.help {
		uf.parent[s] = cur
	}
	uf.help = uf.help[:0] // 清空
	return cur
}

func (uf *UnionFind22) union(s1, s2 string) {
	// 两个 key 都存在才合并
	_, has1 := uf.parent[s1]
	_, has2 := uf.parent[s2]
	if !has1 || !has2 {
		return
	}

	f1 := uf.find(s1)
	f2 := uf.find(s2)

	if f1 != f2 {
		size1 := uf.size[f1]
		size2 := uf.size[f2]

		var big, small string
		if size1 >= size2 {
			big = f1
			small = f2
		} else {
			big = f2
			small = f1
		}

		uf.parent[small] = big
		uf.size[big] = size1 + size2
		uf.sets--
	}
}

// Connect 动态添加陆地 (r,c)
func (uf *UnionFind22) Connect(r, c int) int {
	key := strconv.Itoa(r) + "_" + strconv.Itoa(c)
	if _, exists := uf.parent[key]; !exists {
		uf.parent[key] = key
		uf.size[key] = 1
		uf.sets++

		up := strconv.Itoa(r-1) + "_" + strconv.Itoa(c)
		down := strconv.Itoa(r+1) + "_" + strconv.Itoa(c)
		left := strconv.Itoa(r) + "_" + strconv.Itoa(c-1)
		right := strconv.Itoa(r) + "_" + strconv.Itoa(c+1)

		uf.union(up, key)
		uf.union(down, key)
		uf.union(left, key)
		uf.union(right, key)
	}
	return uf.sets
}

func numIslands22(m, n int, positions [][]int) []int {
	uf := NewUnionFind22()
	var ans []int
	for _, pos := range positions {
		ans = append(ans, uf.Connect(pos[0], pos[1]))
	}
	return ans
}

// ================================================================================
// 测试主函数
// ================================================================================
func main() {
	m := 3
	n := 3
	positions := [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}}

	// 测试数组版
	res1 := numIslands21(m, n, positions)
	fmt.Println("数组并查集结果:", res1) // [1,1,2,3]

	// 测试 HashMap 版
	res2 := numIslands22(m, n, positions)
	fmt.Println("HashMap并查集结果:", res2) // [1,1,2,3]
}
