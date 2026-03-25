package main

import (
	"fmt"
	"sort"
)

// 自定义结构体 Zuo
type Zuo struct {
	value int
}

// 自定义结构体 Node2
type Node2 struct {
	value int
}

func main() {
	// 1. Integer 等值测试
	a := 19000000
	b := 19000000
	fmt.Println(a == b) // true

	// 2. map 中 Integer key 测试
	test := make(map[int]string)
	test[a] = "我是3"
	_, exists := test[b]
	fmt.Println(exists) // true

	// 3. 自定义结构体作为 key（Go 中可直接比较，与 Java 不同）
	z1 := Zuo{1}
	z2 := Zuo{1}
	test2 := make(map[Zuo]string)
	test2[z1] = "我是z1"
	_, exists = test2[z2]
	fmt.Println(exists) // Go：true（Java 是 false，这是语言差异）

	fmt.Println("=====================")

	// 4. HashMap 基本操作（增删改查 O(1)）
	mp := make(map[int]string)
	mp[1000000] = "我是1000000"
	mp[2] = "我是2"
	mp[3] = "我是3"
	mp[4] = "我是4"
	mp[5] = "我是5"
	mp[6] = "我是6"
	mp[1000000] = "我是1000001"

	fmt.Println(mp[1] == "")  // false
	fmt.Println(mp[10] == "") // false

	fmt.Println(mp[4]) // 我是4
	fmt.Println(mp[10])

	mp[4] = "他是4"
	fmt.Println(mp[4])

	delete(mp, 4)
	fmt.Println(mp[4])

	fmt.Println("=====================")

	// 5. int 等值比较
	c := 100000
	d := 100000
	fmt.Println(c == d) // true

	e := 127
	f := 127
	fmt.Println(e == f) // true

	// 6. Node2 作为 key
	map2 := make(map[*Node2]string)
	Node21 := &Node2{1}
	Node22 := Node21
	map2[Node21] = "我是Node21"
	map2[Node22] = "我是Node21"
	fmt.Println(len(map2)) // 1

	fmt.Println("======================")

	// 7. 有序表（TreeMap）→ Go 用 sorted keys 模拟
	// Go 无内置 TreeMap，但可以用 map + 排序实现核心功能
	treeMap := make(map[int]string)
	treeMap[3] = "我是3"
	treeMap[4] = "我是4"
	treeMap[8] = "我是8"
	treeMap[5] = "我是5"
	treeMap[7] = "我是7"
	treeMap[1] = "我是1"
	treeMap[2] = "我是2"

	// 提取并排序 key
	var keys []int
	for k := range treeMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fmt.Println(treeMap[1] != "")  // true
	fmt.Println(treeMap[10] != "") // false

	fmt.Println(treeMap[4])
	fmt.Println(treeMap[10])

	treeMap[4] = "他是4"
	fmt.Println(treeMap[4])

	fmt.Println("新鲜：")
	// 最小 key
	fmt.Println(keys[0])
	// 最大 key
	fmt.Println(keys[len(keys)-1])

	// <= 4
	floor := findFloor(keys, 4)
	fmt.Println(floor)

	// >= 4
	ceiling := findCeiling(keys, 4)
	fmt.Println(ceiling)
}

// 找 <= target 的最大 key
func findFloor(keys []int, target int) int {
	for i := len(keys) - 1; i >= 0; i-- {
		if keys[i] <= target {
			return keys[i]
		}
	}
	return -1
}

// 找 >= target 的最小 key
func findCeiling(keys []int, target int) int {
	for _, k := range keys {
		if k >= target {
			return k
		}
	}
	return -1
}
