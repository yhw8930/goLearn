package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

// ================================================================================
// 解法一：感染法 (DFS) - 最简单
// ================================================================================
func numIslands3(board [][]byte) int {
	islands := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '1' {
				islands++
				infect(board, i, j)
			}
		}
	}
	return islands
}

func infect(board [][]byte, i, j int) {
	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) || board[i][j] != '1' {
		return
	}
	board[i][j] = '0' // 感染成0
	infect(board, i-1, j)
	infect(board, i+1, j)
	infect(board, i, j-1)
	infect(board, i, j+1)
}

// ================================================================================
// 解法二：并查集 (Map 泛型实现)
// ================================================================================
type Dot struct{}

type UnionFind1 struct {
	nodes   map[interface{}]*Node
	parents map[*Node]*Node
	sizeMap map[*Node]int
}

type Node struct {
	value interface{}
}

func NewUnionFind1(values []interface{}) *UnionFind1 {
	uf := &UnionFind1{
		nodes:   make(map[interface{}]*Node),
		parents: make(map[*Node]*Node),
		sizeMap: make(map[*Node]int),
	}
	for _, v := range values {
		node := &Node{value: v}
		uf.nodes[v] = node
		uf.parents[node] = node
		uf.sizeMap[node] = 1
	}
	return uf
}

func (uf *UnionFind1) findFather(cur *Node) *Node {
	path := list.New()
	for cur != uf.parents[cur] {
		path.PushBack(cur)
		cur = uf.parents[cur]
	}
	for path.Len() > 0 {
		node := path.Remove(path.Back()).(*Node)
		uf.parents[node] = cur
	}
	return cur
}

func (uf *UnionFind1) Union(a, b interface{}) {
	aHead := uf.findFather(uf.nodes[a])
	bHead := uf.findFather(uf.nodes[b])
	if aHead != bHead {
		aSize := uf.sizeMap[aHead]
		bSize := uf.sizeMap[bHead]
		var big, small *Node
		if aSize >= bSize {
			big = aHead
			small = bHead
		} else {
			big = bHead
			small = aHead
		}
		uf.parents[small] = big
		uf.sizeMap[big] = aSize + bSize
		delete(uf.sizeMap, small)
	}
}

func (uf *UnionFind1) Sets() int {
	return len(uf.sizeMap)
}

func numIslands1(board [][]byte) int {
	row := len(board)
	col := len(board[0])
	dots := make([][]*Dot, row)
	for i := range dots {
		dots[i] = make([]*Dot, col)
	}
	var dotList []interface{}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == '1' {
				dots[i][j] = &Dot{}
				dotList = append(dotList, dots[i][j])
			}
		}
	}

	uf := NewUnionFind1(dotList)

	// 合并第一行
	for j := 1; j < col; j++ {
		if board[0][j-1] == '1' && board[0][j] == '1' {
			uf.Union(dots[0][j-1], dots[0][j])
		}
	}
	// 合并第一列
	for i := 1; i < row; i++ {
		if board[i-1][0] == '1' && board[i][0] == '1' {
			uf.Union(dots[i-1][0], dots[i][0])
		}
	}
	// 合并其他
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if board[i][j] == '1' {
				if board[i][j-1] == '1' {
					uf.Union(dots[i][j-1], dots[i][j])
				}
				if board[i-1][j] == '1' {
					uf.Union(dots[i-1][j], dots[i][j])
				}
			}
		}
	}
	return uf.Sets()
}

// ================================================================================
// 解法三：并查集 (数组实现 - 最快)
// ================================================================================
type UnionFind2 struct {
	parent []int
	size   []int
	help   []int
	col    int
	sets   int
}

func NewUnionFind2(board [][]byte) *UnionFind2 {
	row := len(board)
	col := len(board[0])
	lenTotal := row * col
	uf := &UnionFind2{
		parent: make([]int, lenTotal),
		size:   make([]int, lenTotal),
		help:   make([]int, lenTotal),
		col:    col,
		sets:   0,
	}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if board[r][c] == '1' {
				i := uf.index(r, c)
				uf.parent[i] = i
				uf.size[i] = 1
				uf.sets++
			}
		}
	}
	return uf
}

func (uf *UnionFind2) index(r, c int) int {
	return r*uf.col + c
}

func (uf *UnionFind2) find(i int) int {
	hi := 0
	for i != uf.parent[i] {
		uf.help[hi] = i
		hi++
		i = uf.parent[i]
	}
	for hi--; hi >= 0; hi-- {
		uf.parent[uf.help[hi]] = i
	}
	return i
}

func (uf *UnionFind2) Union(r1, c1, r2, c2 int) {
	i1 := uf.index(r1, c1)
	i2 := uf.index(r2, c2)
	f1 := uf.find(i1)
	f2 := uf.find(i2)
	if f1 != f2 {
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

func (uf *UnionFind2) Sets() int {
	return uf.sets
}

func numIslands2(board [][]byte) int {
	row := len(board)
	col := len(board[0])
	uf := NewUnionFind2(board)

	// 第一行
	for j := 1; j < col; j++ {
		if board[0][j-1] == '1' && board[0][j] == '1' {
			uf.Union(0, j-1, 0, j)
		}
	}
	// 第一列
	for i := 1; i < row; i++ {
		if board[i-1][0] == '1' && board[i][0] == '1' {
			uf.Union(i-1, 0, i, 0)
		}
	}
	// 其他
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if board[i][j] == '1' {
				if board[i][j-1] == '1' {
					uf.Union(i, j-1, i, j)
				}
				if board[i-1][j] == '1' {
					uf.Union(i-1, j, i, j)
				}
			}
		}
	}
	return uf.Sets()
}

// ================================================================================
// 测试工具
// ================================================================================
func generateRandomMatrix(row, col int) [][]byte {
	rand.Seed(time.Now().UnixNano())
	board := make([][]byte, row)
	for i := range board {
		board[i] = make([]byte, col)
		for j := range board[i] {
			if rand.Float64() < 0.5 {
				board[i][j] = '1'
			} else {
				board[i][j] = '0'
			}
		}
	}
	return board
}

func copyBoard(board [][]byte) [][]byte {
	row := len(board)
	col := len(board[0])
	ans := make([][]byte, row)
	for i := range ans {
		ans[i] = make([]byte, col)
		copy(ans[i], board[i])
	}
	return ans
}

// ================================================================================
// 主函数测试
// ================================================================================
func main() {
	row, col := 1000, 1000
	board1 := generateRandomMatrix(row, col)
	board2 := copyBoard(board1)
	board3 := copyBoard(board1)

	fmt.Println("感染法、并查集(Map)、并查集(数组) 运行结果 & 时间")
	fmt.Printf("矩阵规模: %d * %d\n", row, col)

	start := time.Now()
	fmt.Println("感染法结果:", numIslands3(board1))
	fmt.Printf("感染法耗时: %d ms\n", time.Since(start).Milliseconds())

	start = time.Now()
	fmt.Println("并查集(Map)结果:", numIslands1(board2))
	fmt.Printf("并查集(Map)耗时: %d ms\n", time.Since(start).Milliseconds())

	start = time.Now()
	fmt.Println("并查集(数组)结果:", numIslands2(board3))
	fmt.Printf("并查集(数组)耗时: %d ms\n", time.Since(start).Milliseconds())
}
