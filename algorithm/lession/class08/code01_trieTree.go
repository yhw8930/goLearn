package main

import "fmt"

/*
======================== 数据结构 ========================
*/

type Node1 struct {
	pass  int
	end   int
	nexts [26]*Node1
}

/*
======================== Trie1（数组实现） ========================

1. 问题：
实现一个前缀树（Trie），支持：
- insert(word)
- delete(word)
- search(word)
- prefixNumber(pre)

2. 思路：
- 用固定数组 nexts[26] 表示 a-z 路径
- pass：经过该节点的字符串数量
- end：以该节点结尾的字符串数量

核心操作：
- 插入：路径节点 pass++
- 删除：路径节点 pass--，为0则剪枝
- 查询：走到末尾看 end
- 前缀：看 pass

时间复杂度：O(K)（K为字符串长度）
空间复杂度：O(26*N)
*/

type Trie1 struct {
	root *Node1
}

func NewTrie1() *Trie1 {
	return &Trie1{root: &Node1{}}
}

func (t *Trie1) Insert(word string) {
	if word == "" {
		return
	}
	node := t.root
	node.pass++

	for _, c := range word {
		index := c - 'a'
		if node.nexts[index] == nil {
			node.nexts[index] = &Node1{}
		}
		node = node.nexts[index]
		node.pass++
	}
	node.end++
}

func (t *Trie1) Search(word string) int {
	if word == "" {
		return 0
	}
	node := t.root

	for _, c := range word {
		index := c - 'a'
		if node.nexts[index] == nil {
			return 0
		}
		node = node.nexts[index]
	}
	return node.end
}

func (t *Trie1) PrefixNumber(pre string) int {
	if pre == "" {
		return 0
	}
	node := t.root

	for _, c := range pre {
		index := c - 'a'
		if node.nexts[index] == nil {
			return 0
		}
		node = node.nexts[index]
	}
	return node.pass
}

func (t *Trie1) Delete(word string) {
	if t.Search(word) == 0 {
		return
	}

	node := t.root
	node.pass--

	for _, c := range word {
		index := c - 'a'
		next := node.nexts[index]

		next.pass--
		if next.pass == 0 {
			node.nexts[index] = nil
			return
		}

		node = next
	}
	node.end--
}

/*
======================== Trie2（HashMap实现） ========================

1. 问题：
同 Trie1，但用 HashMap 替代数组

2. 思路：
- nexts 用 map[rune]*Node2
- 逻辑同 Trie1

时间复杂度：O(K)
空间复杂度：O(N)
*/

type Node2 struct {
	pass  int
	end   int
	nexts map[rune]*Node2
}

type Trie2 struct {
	root *Node2
}

func NewTrie2() *Trie2 {
	return &Trie2{root: &Node2{nexts: make(map[rune]*Node2)}}
}

func (t *Trie2) Insert(word string) {
	if word == "" {
		return
	}
	node := t.root
	node.pass++

	for _, c := range word {
		if node.nexts[c] == nil {
			node.nexts[c] = &Node2{nexts: make(map[rune]*Node2)}
		}
		node = node.nexts[c]
		node.pass++
	}
	node.end++
}

func (t *Trie2) Search(word string) int {
	if word == "" {
		return 0
	}
	node := t.root

	for _, c := range word {
		if node.nexts[c] == nil {
			return 0
		}
		node = node.nexts[c]
	}
	return node.end
}

func (t *Trie2) PrefixNumber(pre string) int {
	if pre == "" {
		return 0
	}
	node := t.root

	for _, c := range pre {
		if node.nexts[c] == nil {
			return 0
		}
		node = node.nexts[c]
	}
	return node.pass
}

func (t *Trie2) Delete(word string) {
	if t.Search(word) == 0 {
		return
	}

	node := t.root
	node.pass--

	for _, c := range word {
		next := node.nexts[c]
		next.pass--

		if next.pass == 0 {
			delete(node.nexts, c)
			return
		}

		node = next
	}
	node.end--
}

/*
======================== Right（暴力对数器） ========================

1. 问题：
用 HashMap 模拟 Trie 功能，作为正确性对比

2. 思路：
- insert/delete：维护词频
- search：直接查 map
- prefix：遍历所有 key 判断 prefix

时间复杂度：
- search O(1)
- prefix O(N)

空间复杂度：O(N)
*/

type Right struct {
	box map[string]int
}

func NewRight() *Right {
	return &Right{box: make(map[string]int)}
}

func (r *Right) Insert(word string) {
	r.box[word]++
}

func (r *Right) Delete(word string) {
	if r.box[word] == 1 {
		delete(r.box, word)
	} else if r.box[word] > 1 {
		r.box[word]--
	}
}

func (r *Right) Search(word string) int {
	return r.box[word]
}

func (r *Right) PrefixNumber(pre string) int {
	ans := 0
	for k := range r.box {
		if len(k) >= len(pre) && k[:len(pre)] == pre {
			ans++
		}
	}
	return ans
}

/*
======================== 随机测试 ========================
*/

func randomString(maxLen int) string {
	l := int(float64(maxLen)*randFloat()) + 1
	bytes := make([]byte, l)

	for i := 0; i < l; i++ {
		bytes[i] = byte('a' + int(randFloat()*6))
	}
	return string(bytes)
}

func randFloat() float64 {
	return float64(seed()) / float64(1<<31)
}

var x int64 = 123456789

func seed() int {
	x = x*1103515245 + 12345
	return int((x / 65536) % 32768)
}

func main() {
	arrLen := 100
	strLen := 20
	testTimes := 1000

	for t := 0; t < testTimes; t++ {
		arr := make([]string, int(randFloat()*float64(arrLen))+1)
		for i := range arr {
			arr[i] = randomString(strLen)
		}

		trie1 := NewTrie1()
		trie2 := NewTrie2()
		right := NewRight()

		for _, s := range arr {
			op := randFloat()

			if op < 0.25 {
				trie1.Insert(s)
				trie2.Insert(s)
				right.Insert(s)
			} else if op < 0.5 {
				trie1.Delete(s)
				trie2.Delete(s)
				right.Delete(s)
			} else if op < 0.75 {
				a := trie1.Search(s)
				b := trie2.Search(s)
				c := right.Search(s)
				if a != b || b != c {
					fmt.Println("Oops search!")
				}
			} else {
				a := trie1.PrefixNumber(s)
				b := trie2.PrefixNumber(s)
				c := right.PrefixNumber(s)
				if a != b || b != c {
					fmt.Println("Oops prefix!")
				}
			}
		}
	}

	fmt.Println("finish!")
}
