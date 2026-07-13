package main

import (
	"fmt"
	"math/rand"
)

type Employee struct {
	Happy int
	Nexts []*Employee
}

// MaxHappy1 求公司聚会能获得的最大快乐值。公司组织是一棵多叉树，每名员工有快乐值；若员工参加，其直接下属不能参加。
// 递归参数表示上级是否参加：上级参加则当前员工不能来；否则分别尝试当前员工来或不来。
// 时间复杂度：最坏 O(2^N)，因为“上级不来”状态会重复计算子问题。空间复杂度：O(H)。
func MaxHappy1(boss *Employee) int {
	if boss == nil {
		return 0
	}
	return maxHappyProcess1(boss, false)
}

func maxHappyProcess1(cur *Employee, parentComes bool) int {
	if parentComes {
		ans := 0
		for _, next := range cur.Nexts {
			ans += maxHappyProcess1(next, false)
		}
		return ans
	}
	come, notCome := cur.Happy, 0
	for _, next := range cur.Nexts {
		come += maxHappyProcess1(next, true)
		notCome += maxHappyProcess1(next, false)
	}
	return max(come, notCome)
}

type maxHappyInfo struct{ no, yes int }

// MaxHappy2 使用树形 DP 求最大快乐值。
// 每名员工一次性返回“自己不参加”和“自己参加”两种最优值：参加时下属只能不参加，不参加时下属可选两者最大值。
// 时间复杂度：O(N)。空间复杂度：O(H)。
func MaxHappy2(boss *Employee) int {
	info := maxHappyProcess2(boss)
	return max(info.no, info.yes)
}

func maxHappyProcess2(x *Employee) maxHappyInfo {
	if x == nil {
		return maxHappyInfo{}
	}
	info := maxHappyInfo{yes: x.Happy}
	for _, next := range x.Nexts {
		nextInfo := maxHappyProcess2(next)
		info.no += max(nextInfo.no, nextInfo.yes)
		info.yes += nextInfo.no
	}
	return info
}

func maxHappyGenerateBoss(maxLevel, maxNexts, maxHappy int) *Employee {
	if rand.Float64() < 0.02 {
		return nil
	}
	boss := &Employee{Happy: rand.Intn(maxHappy + 1)}
	maxHappyGenerateNexts(boss, 1, maxLevel, maxNexts, maxHappy)
	return boss
}

func maxHappyGenerateNexts(e *Employee, level, maxLevel, maxNexts, maxHappy int) {
	if level > maxLevel {
		return
	}
	for i, size := 0, rand.Intn(maxNexts+1); i < size; i++ {
		next := &Employee{Happy: rand.Intn(maxHappy + 1)}
		e.Nexts = append(e.Nexts, next)
		maxHappyGenerateNexts(next, level+1, maxLevel, maxNexts, maxHappy)
	}
}

// main 随机生成公司层级，对比状态递归和树形 DP 两种最大快乐值算法。
func main() {
	for i := 0; i < 1000; i++ {
		boss := maxHappyGenerateBoss(4, 5, 100)
		if MaxHappy1(boss) != MaxHappy2(boss) {
			fmt.Println("Oops!")
			return
		}
	}
	fmt.Println("finish!")
}
