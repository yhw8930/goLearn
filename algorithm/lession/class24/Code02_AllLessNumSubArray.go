package main

import (
	"fmt"
	"math/rand"
)

// AllLessNumSubArrayRight 解决“最大值减最小值达标的子数组数量”问题：
// 给定数组 arr 和整数 sum，统计有多少个子数组满足 max(subarray)-min(subarray) <= sum。
// 子数组必须连续；如果 arr 为空或 sum < 0，返回 0。
// 这个版本是暴力对数器：枚举所有 L、R，并扫描子数组求最大值和最小值。
//
// 时间复杂度：O(N^3)。
// 空间复杂度：O(1)。
func AllLessNumSubArrayRight(arr []int, sum int) int {
	if len(arr) == 0 || sum < 0 {
		return 0
	}
	n := len(arr)
	count := 0
	for l := 0; l < n; l++ {
		for r := l; r < n; r++ {
			maxValue := arr[l]
			minValue := arr[l]
			for i := l + 1; i <= r; i++ {
				if arr[i] > maxValue {
					maxValue = arr[i]
				}
				if arr[i] < minValue {
					minValue = arr[i]
				}
			}
			if maxValue-minValue <= sum {
				count++
			}
		}
	}
	return count
}

// AllLessNumSubArrayNum 返回满足 max-min <= sum 的子数组数量。
// 用两个双端队列维护当前窗口内的最大值和最小值。对每个左边界 L，右边界 R 尽量向右扩，
// 一旦窗口不达标就停止，此时以 L 开头且右端点在 [L,R) 范围内的子数组全部达标。
//
// 时间复杂度：O(N)，每个下标最多进入和离开窗口一次。
// 空间复杂度：O(N)，两个双端队列空间。
func AllLessNumSubArrayNum(arr []int, sum int) int {
	if len(arr) == 0 || sum < 0 {
		return 0
	}
	n := len(arr)
	count := 0
	maxWindow := make([]int, 0)
	minWindow := make([]int, 0)
	r := 0
	for l := 0; l < n; l++ {
		for r < n {
			for len(maxWindow) > 0 && arr[maxWindow[len(maxWindow)-1]] <= arr[r] {
				maxWindow = maxWindow[:len(maxWindow)-1]
			}
			maxWindow = append(maxWindow, r)
			for len(minWindow) > 0 && arr[minWindow[len(minWindow)-1]] >= arr[r] {
				minWindow = minWindow[:len(minWindow)-1]
			}
			minWindow = append(minWindow, r)
			if arr[maxWindow[0]]-arr[minWindow[0]] > sum {
				break
			}
			r++
		}
		count += r - l
		if maxWindow[0] == l {
			maxWindow = maxWindow[1:]
		}
		if minWindow[0] == l {
			minWindow = minWindow[1:]
		}
	}
	return count
}

func allLessNumSubArrayGenerateRandomArray(maxLen, maxValue int) []int {
	length := rand.Intn(maxLen + 1)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue+1)
	}
	return arr
}

func allLessNumSubArrayPrintArray(arr []int) {
	if arr == nil {
		return
	}
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
