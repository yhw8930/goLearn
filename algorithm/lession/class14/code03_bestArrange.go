package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 题目：给定多个会议的开始和结束时间，同一时间只能安排一个会议，求最多能安排多少场。
// 暴力方法枚举每场能否作为下一场会议，尝试所有合法选择。
// 核心思路：贪心按结束时间从早到晚排序，每次选择能参加且结束最早的会议。
// 越早结束会给后续留下越多时间，因此这个局部选择能得到最多会议数量。
// 时间复杂度：暴力为指数级，贪心排序为 O(NlogN)。
// 空间复杂度：暴力递归 O(N)，贪心排序视实现而定。

// Program 会议结构体
type Program struct {
	start int
	end   int
}

// MeetingArrange 会议安排器（结构体版）
type MeetingArrange struct {
	programs []Program
}

// NewMeetingArrange 构造函数
func NewMeetingArrange(programs []Program) *MeetingArrange {
	return &MeetingArrange{programs: programs}
}

// ==================== 方法1：暴力递归（用来验证答案）====================
// BestArrange1 是暴力搜索版本。
// 它枚举当前时间线下能安排的每一场会议，递归尝试所有后续安排。
// 时间复杂度：指数级，枚举所有可行会议选择顺序。
// 空间复杂度：O(N)，递归深度和复制数组。
func (m *MeetingArrange) BestArrange1() int {
	return m.process(m.programs, 0, 0)
}

func (m *MeetingArrange) process(programs []Program, done int, timeLine int) int {
	if len(programs) == 0 {
		return done
	}

	maxCount := done
	for i := 0; i < len(programs); i++ {
		if programs[i].start >= timeLine {
			nextPrograms := m.copyButExcept(programs, i)
			cur := m.process(nextPrograms, done+1, programs[i].end)
			if cur > maxCount {
				maxCount = cur
			}
		}
	}
	return maxCount
}

// copyButExcept 复制数组，排除第 i 个
func (m *MeetingArrange) copyButExcept(programs []Program, i int) []Program {
	ans := make([]Program, 0, len(programs)-1)
	for k := 0; k < len(programs); k++ {
		if k != i {
			ans = append(ans, programs[k])
		}
	}
	return ans
}

// ==================== 方法2：贪心算法（最优解）====================
// BestArrange2 是按结束时间排序的贪心版本。
// 每次选择最早结束且不冲突的会议，为后面保留最多空间。
// 时间复杂度：O(NlogN)，主要来自按结束时间排序。
// 空间复杂度：O(1) 到 O(N)，取决于排序实现。
func (m *MeetingArrange) BestArrange2() int {
	// 按结束时间升序排序
	sort.Slice(m.programs, func(i, j int) bool {
		return m.programs[i].end < m.programs[j].end
	})

	timeLine := 0
	result := 0
	for _, p := range m.programs {
		if timeLine <= p.start {
			result++
			timeLine = p.end
		}
	}
	return result
}

// ==================== 测试工具：随机生成会议 ====================
func generatePrograms(programSize, timeMax int) []Program {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(programSize + 1)
	ans := make([]Program, size)

	for i := 0; i < size; i++ {
		r1 := rand.Intn(timeMax + 1)
		r2 := rand.Intn(timeMax + 1)
		if r1 == r2 {
			ans[i] = Program{r1, r1 + 1}
		} else {
			ans[i] = Program{min(r1, r2), max(r1, r2)}
		}
	}
	return ans
}

// ==================== 主测试 ====================
func main() {
	programSize := 12
	timeMax := 20
	testTimes := 10000
	fmt.Println("测试开始...")

	for i := 0; i < testTimes; i++ {
		programs := generatePrograms(programSize, timeMax)
		ma := NewMeetingArrange(programs)
		ans1 := ma.BestArrange1()
		ans2 := ma.BestArrange2()
		if ans1 != ans2 {
			fmt.Println("Oops!")
			return
		}
	}

	fmt.Println("finish! ✅ 暴力与贪心完全一致")
}
