package main

import "math/rand"

// 单调栈：给定一个数组 arr，要求对每个位置 i，求出它左边离它最近且比它小的数的下标，
// 以及右边离它最近且比它小的数的下标（找不到记为 -1）。结果是一个 N×2 的二维数组，
// res[i][0] 是 i 左边最近的更小值下标，res[i][1] 是 i 右边最近的更小值下标。
//
// 核心思想：维护一个“栈底到栈顶值严格递增”的单调栈，栈里存下标。
// 当新元素小于栈顶对应的值时，弹出栈顶 j：此时新元素就是 j 右边最近的更小值，
// 弹出后新的栈顶就是 j 左边最近的更小值。

// GetNearLessNoRepeat 适用于数组中无重复值的情况。
// 栈中存放严格递增的下标，遍历时遇到更小值就结算被弹出的下标。
//
// 时间复杂度：O(N)，每个下标进栈出栈各一次。
// 空间复杂度：O(N)。
func GetNearLessNoRepeat(arr []int) [][2]int {
	res := make([][2]int, len(arr))
	stack := make([]int, 0, len(arr)) // 只存位置
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			leftLessIndex := -1
			if len(stack) > 0 {
				leftLessIndex = stack[len(stack)-1]
			}
			res[j][0] = leftLessIndex
			res[j][1] = i
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		j := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		leftLessIndex := -1
		if len(stack) > 0 {
			leftLessIndex = stack[len(stack)-1]
		}
		res[j][0] = leftLessIndex
		res[j][1] = -1
	}
	return res
}

// GetNearLess 适用于数组中可能有重复值的情况。
// 栈里每个元素是一组下标（值相等的下标压在同一层），结算时整组一起赋值，
// 左边最近更小值取下一层那组的最后一个下标。
//
// 时间复杂度：O(N)。
// 空间复杂度：O(N)。
func GetNearLess(arr []int) [][2]int {
	res := make([][2]int, len(arr))
	stack := make([][]int, 0, len(arr)) // 每层是值相等的下标列表
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1][0]] > arr[i] {
			popIs := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			leftLessIndex := -1
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				leftLessIndex = top[len(top)-1]
			}
			for _, popi := range popIs {
				res[popi][0] = leftLessIndex
				res[popi][1] = i
			}
		}
		if len(stack) > 0 && arr[stack[len(stack)-1][0]] == arr[i] {
			top := &stack[len(stack)-1]
			*top = append(*top, i)
		} else {
			stack = append(stack, []int{i})
		}
	}
	for len(stack) > 0 {
		popIs := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		leftLessIndex := -1
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			leftLessIndex = top[len(top)-1]
		}
		for _, popi := range popIs {
			res[popi][0] = leftLessIndex
			res[popi][1] = -1
		}
	}
	return res
}

// nearLessRightWay 是对数器用的暴力解：对每个 i 向左、向右各扫一遍找最近的更小值。
//
// 时间复杂度：O(N^2)。
// 空间复杂度：O(N)。
func nearLessRightWay(arr []int) [][2]int {
	res := make([][2]int, len(arr))
	for i := 0; i < len(arr); i++ {
		leftLessIndex := -1
		rightLessIndex := -1
		for cur := i - 1; cur >= 0; cur-- {
			if arr[cur] < arr[i] {
				leftLessIndex = cur
				break
			}
		}
		for cur := i + 1; cur < len(arr); cur++ {
			if arr[cur] < arr[i] {
				rightLessIndex = cur
				break
			}
		}
		res[i][0] = leftLessIndex
		res[i][1] = rightLessIndex
	}
	return res
}

// nearLessRandomArrayNoRepeat 生成 0..n-1 的随机排列（无重复）。
func nearLessRandomArrayNoRepeat(size int) []int {
	arr := make([]int, rand.Intn(size)+1)
	for i := range arr {
		arr[i] = i
	}
	for i := range arr {
		swapIndex := rand.Intn(len(arr))
		arr[swapIndex], arr[i] = arr[i], arr[swapIndex]
	}
	return arr
}

// nearLessRandomArray 生成可能有重复值的随机数组。
func nearLessRandomArray(size, max int) []int {
	arr := make([]int, rand.Intn(size)+1)
	for i := range arr {
		arr[i] = rand.Intn(max) - rand.Intn(max)
	}
	return arr
}

func nearLessIsEqual(res1, res2 [][2]int) bool {
	if len(res1) != len(res2) {
		return false
	}
	for i := range res1 {
		if res1[i][0] != res2[i][0] || res1[i][1] != res2[i][1] {
			return false
		}
	}
	return true
}
