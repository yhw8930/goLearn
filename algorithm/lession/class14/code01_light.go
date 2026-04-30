package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// ==================== 方法1：暴力递归（正确，用于验证）====================
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
