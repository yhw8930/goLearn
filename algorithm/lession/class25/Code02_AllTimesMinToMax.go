package main

import "math/rand"

// 子数组指标最大值：给定一个【全是正数】的数组 arr，定义一个子数组的指标为
// (子数组累加和) * (子数组中的最小值)。在所有子数组中求这个指标的最大值。
// 因为都是正数，所以以某个数 arr[j] 作为最小值时，子数组扩得越宽累加和越大，
// 于是只需找到每个 arr[j] 向左、向右能延伸到的“仍不小于它”的最大范围即可。

// AllTimesMinToMaxMax1 是对数器用的暴力解：枚举所有子数组，直接求累加和与最小值。
//
// 时间复杂度：O(N^3)。
// 空间复杂度：O(1)。
func AllTimesMinToMaxMax1(arr []int) int {
	maxValue := -1 << 62
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			minNum := 1 << 62
			sum := 0
			for k := i; k <= j; k++ {
				sum += arr[k]
				if arr[k] < minNum {
					minNum = arr[k]
				}
			}
			if minNum*sum > maxValue {
				maxValue = minNum * sum
			}
		}
	}
	return maxValue
}

// AllTimesMinToMaxMax2 用单调栈求解：以 arr[j] 为最小值时，
// 它左边最近的更小值和右边最近的更小值之间的区间就是合法范围，
// 配合前缀和 sums 在 O(1) 时间算出该区间的累加和。
//
// 时间复杂度：O(N)。
// 空间复杂度：O(N)。
func AllTimesMinToMaxMax2(arr []int) int {
	size := len(arr)
	sums := make([]int, size)
	sums[0] = arr[0]
	for i := 1; i < size; i++ {
		sums[i] = sums[i-1] + arr[i]
	}
	maxValue := -1 << 62
	stack := make([]int, 0, size)
	for i := 0; i < size; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var sum int
			if len(stack) == 0 {
				sum = sums[i-1]
			} else {
				sum = sums[i-1] - sums[stack[len(stack)-1]]
			}
			if cur := sum * arr[j]; cur > maxValue {
				maxValue = cur
			}
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		j := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		var sum int
		if len(stack) == 0 {
			sum = sums[size-1]
		} else {
			sum = sums[size-1] - sums[stack[len(stack)-1]]
		}
		if cur := sum * arr[j]; cur > maxValue {
			maxValue = cur
		}
	}
	return maxValue
}

// allTimesMinToMaxRandomArray 生成长度 10~29、元素 0~100 的随机正数数组。
func allTimesMinToMaxRandomArray() []int {
	arr := make([]int, rand.Intn(20)+10)
	for i := range arr {
		arr[i] = rand.Intn(101)
	}
	return arr
}
