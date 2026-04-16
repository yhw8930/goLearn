package main

import (
	"fmt"
	"math/rand"
	"time"
)

//
// =======================
// code02Trie1（数组实现）
// =======================
//

// code02Node code02Trie节点（数组版）
//
// 【问题】
// 构建仅支持 a-z 的前缀树节点
//
// 【解题思路】
// 用固定数组 nexts[26] 表示26条分支路径
// pass：经过该节点的字符串数量
// end：以该节点结尾的字符串数量
//
// 【时空复杂度】
// 时间：O(1)
// 空间：O(1)
type code02Node struct {
	pass  int
	end   int
	nexts [26]*code02Node
}

type code02Trie1 struct {
	root *code02Node
}

func Newcode02Trie1() *code02Trie1 {
	return &code02Trie1{root: &code02Node{}}
}

// Insert 插入字符串
//
// 【问题】
// 将字符串加入code02Trie，并维护路径统计信息
//
// 【解题思路】
// 1. root.pass++ 表示整体字符串数量+1
// 2. 按字符逐层走
// 3. 没有路径则创建节点
// 4. 每个节点 pass++
// 5. 最终节点 end++
//
// 【时空复杂度】
// 时间：O(L)
// 空间：O(L)
func (t *code02Trie1) Insert(word string) {
	if word == "" {
		return
	}
	node := t.root
	node.pass++

	for i := 0; i < len(word); i++ {
		path := word[i] - 'a'
		if node.nexts[path] == nil {
			node.nexts[path] = &code02Node{}
		}
		node = node.nexts[path]
		node.pass++
	}
	node.end++
}

// Search 查询字符串出现次数
//
// 【问题】
// 查询word在code02Trie中出现了多少次
//
// 【解题思路】
// 逐字符走code02Trie
// 若路径断裂返回0
// 否则返回 end
//
// 【时空复杂度】
// 时间：O(L)
// 空间：O(1)
func (t *code02Trie1) Search(word string) int {
	if word == "" {
		return 0
	}
	node := t.root
	for i := 0; i < len(word); i++ {
		path := word[i] - 'a'
		if node.nexts[path] == nil {
			return 0
		}
		node = node.nexts[path]
	}
	return node.end
}

// PrefixNumber 前缀出现次数
//
// 【问题】
// 统计有多少字符串以pre为前缀
//
// 【解题思路】
// 走到prefix最后节点
// 返回pass（经过该节点的字符串数量）
//
// 【时空复杂度】
// 时间：O(L)
// 空间：O(1)
func (t *code02Trie1) PrefixNumber(pre string) int {
	if pre == "" {
		return 0
	}
	node := t.root
	for i := 0; i < len(pre); i++ {
		path := pre[i] - 'a'
		if node.nexts[path] == nil {
			return 0
		}
		node = node.nexts[path]
	}
	return node.pass
}

// Delete 删除字符串
//
// 【问题】
// 从code02Trie中删除一个字符串（如果存在）
//
// 【解题思路】
// 1. 先判断是否存在
// 2. root.pass--
// 3. 沿路径 pass--
// 4. 如果某节点 pass==0，直接断开（剪枝）
// 5. 最后 end--
//
// 【时空复杂度】
// 时间：O(L)
// 空间：O(1)
func (t *code02Trie1) Delete(word string) {
	if t.Search(word) == 0 {
		return
	}

	node := t.root
	node.pass--

	for i := 0; i < len(word); i++ {
		path := word[i] - 'a'
		next := node.nexts[path]

		next.pass--

		if next.pass == 0 {
			node.nexts[path] = nil
			return
		}

		node = next
	}
	node.end--
}

//
// =======================
// code02Trie2（HashMap实现）
// =======================
//

type code02Node2 struct {
	pass  int
	end   int
	nexts map[byte]*code02Node2
}

type code02Trie2 struct {
	root *code02Node2
}

func Newcode02Trie2() *code02Trie2 {
	return &code02Trie2{root: &code02Node2{nexts: make(map[byte]*code02Node2)}}
}

// Insert 插入字符串
//
// 【问题】
// 使用HashMap实现code02Trie插入
//
// 【解题思路】
// 用map代替固定数组
// 动态创建分支
//
// 【时空复杂度】
// 时间：O(L)
// 空间：O(L)
func (t *code02Trie2) Insert(word string) {
	if word == "" {
		return
	}
	node := t.root
	node.pass++

	for i := 0; i < len(word); i++ {
		c := word[i]
		if node.nexts[c] == nil {
			node.nexts[c] = &code02Node2{nexts: make(map[byte]*code02Node2)}
		}
		node = node.nexts[c]
		node.pass++
	}
	node.end++
}

// Search 查询字符串
//
// 【问题】
// 查询字符串出现次数
//
// 【解题思路】
// map逐层查找
//
// 【时空复杂度】
// 时间：O(L)
// 空间：O(1)
func (t *code02Trie2) Search(word string) int {
	if word == "" {
		return 0
	}
	node := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if node.nexts[c] == nil {
			return 0
		}
		node = node.nexts[c]
	}
	return node.end
}

// PrefixNumber 前缀统计
//
// 【问题】
// 统计前缀出现次数
//
// 【解题思路】
// 走到前缀节点返回pass
//
// 【时空复杂度】
// 时间：O(L)
// 空间：O(1)
func (t *code02Trie2) PrefixNumber(pre string) int {
	if pre == "" {
		return 0
	}
	node := t.root
	for i := 0; i < len(pre); i++ {
		c := pre[i]
		if node.nexts[c] == nil {
			return 0
		}
		node = node.nexts[c]
	}
	return node.pass
}

// Delete 删除字符串
//
// 【问题】
// 从code02Trie中删除字符串
//
// 【解题思路】
// 类似code02Trie1，但map结构删除用delete
//
// 【时空复杂度】
// 时间：O(L)
// 空间：O(1)
func (t *code02Trie2) Delete(word string) {
	if t.Search(word) == 0 {
		return
	}

	node := t.root
	node.pass--

	for i := 0; i < len(word); i++ {
		c := word[i]
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

//
// =======================
// code02Right（对数器）
// =======================
//

type code02Right struct {
	box map[string]int
}

func Newcode02Right() *code02Right {
	return &code02Right{box: make(map[string]int)}
}

func (r *code02Right) Insert(word string) {
	r.box[word]++
}

func (r *code02Right) Delete(word string) {
	if r.box[word] > 0 {
		if r.box[word] == 1 {
			delete(r.box, word)
		} else {
			r.box[word]--
		}
	}
}

func (r *code02Right) Search(word string) int {
	return r.box[word]
}

func (r *code02Right) PrefixNumber(pre string) int {
	count := 0
	for k, v := range r.box {
		if len(k) >= len(pre) && k[:len(pre)] == pre {
			count += v
		}
	}
	return count
}

//
// =======================
// test
// =======================
//

func code02RandomString(maxLen int) string {
	length := rand.Intn(maxLen) + 1
	s := make([]byte, length)
	for i := 0; i < length; i++ {
		s[i] = byte('a' + rand.Intn(6))
	}
	return string(s)
}

func randomArray(arrLen, strLen int) []string {
	n := rand.Intn(arrLen) + 1
	res := make([]string, n)
	for i := range res {
		res[i] = code02RandomString(strLen)
	}
	return res
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 2000; i++ {
		arr := randomArray(100, 20)

		t1 := Newcode02Trie1()
		t2 := Newcode02Trie2()
		r := Newcode02Right()

		for _, s := range arr {
			x := rand.Float64()

			if x < 0.25 {
				t1.Insert(s)
				t2.Insert(s)
				r.Insert(s)
			} else if x < 0.5 {
				t1.Delete(s)
				t2.Delete(s)
				r.Delete(s)
			} else if x < 0.75 {
				if t1.Search(s) != t2.Search(s) || t2.Search(s) != r.Search(s) {
					fmt.Println("Oops search")
					return
				}
			} else {
				if t1.PrefixNumber(s) != t2.PrefixNumber(s) ||
					t2.PrefixNumber(s) != r.PrefixNumber(s) {
					fmt.Println("Oops prefix")
					return
				}
			}
		}
	}

	fmt.Println("finish")
}
