package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 题目：字符串 road 中 X 表示墙，. 表示居民点，灯可以照亮当前位置和相邻位置，求照亮所有居民点的最少灯数。
// 墙不需要照亮也不能放灯，居民点必须被自己或左右相邻位置的灯覆盖。
// 核心思路：暴力版本枚举每个位置放或不放灯，最后检查是否全部覆盖。
// 贪心版本从左到右扫描，遇到居民点就尽量把灯放在能覆盖更远的位置，然后跳过已覆盖区域。
// 时间复杂度：暴力为 O(2^N)，贪心为 O(N)。
// 空间复杂度：暴力递归和集合为 O(N)，贪心为 O(1)。

// ==================== 方法1：暴力递归（正确，用于验证）====================
// minLight1 是暴力递归版本。
// 每个位置尝试放灯或不放灯，最后检查所有居民点是否被照亮。
// 时间复杂度：O(2^N)。
// 空间复杂度：O(N)，递归深度和 lights 集合。
func minLight1(road string) int {
	if len(road) == 0 {
		return 0
	}
	lights := make(map[int]bool)
	return process([]byte(road), 0, lights)
}

func process(str []byte, index int, lights map[int]bool) int {
	if index == len(str) {
		for i := 0; i < len(str); i++ {
			if str[i] == '.' {
				has := lights[i-1] || lights[i] || lights[i+1]
				if !has {
					return math.MaxInt32
				}
			}
		}
		return len(lights)
	}

	no := process(str, index+1, lights)

	yes := math.MaxInt32
	if str[index] == '.' {
		lights[index] = true
		yes = process(str, index+1, lights)
		delete(lights, index)
	}

	return min(no, yes)
}

// ==================== 方法2：贪心算法（正确！面试首选）====================
// minLight2 是贪心版本。
// 从左到右遇到居民点就放一盏灯，并根据后一个位置情况跳过已覆盖区域。
// 时间复杂度：O(N)。
// 空间复杂度：O(1)。
func minLight2(road string) int {
	str := []byte(road)
	i := 0
	light := 0

	for i < len(str) {
		if str[i] == 'X' {
			i++
		} else {
			light++
			if i+1 == len(str) {
				break
			}
			if str[i+1] == 'X' {
				i += 2
			} else {
				i += 3
			}
		}
	}
	return light
}
func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	res := make([]byte, length)
	for i := 0; i < length; i++ {
		if rand.Float64() < 0.5 {
			res[i] = 'X'
		} else {
			res[i] = '.'
		}
	}
	return string(res)
}

// ==================== 主测试 ====================
func main() {
	testLen := 14
	testTimes := 1000
	fmt.Println("测试开始...")

	for i := 0; i < testTimes; i++ {
		s := randomString(testLen)
		ans1 := minLight1(s)
		ans2 := minLight2(s)

		if ans1 != ans2 {
			fmt.Println("出错！字符串：", s)
			fmt.Println("暴力：", ans1, " 贪心：", ans2)
			return
		}
	}

	fmt.Println("✅ 测试全部通过！暴力 == 贪心，完全正确")
}
