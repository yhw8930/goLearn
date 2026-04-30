package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

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
