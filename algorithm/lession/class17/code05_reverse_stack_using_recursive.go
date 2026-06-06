package main

// 题目：只允许使用递归函数，不借助额外数据结构，把一个栈逆序。
// 关键难点是普通栈只能弹出栈顶，但逆序需要把原栈底元素放到新栈顶。
// 核心思路：递归函数 removeBottom 弹出并返回当前栈底元素，同时保持其他元素顺序不变。
// 主递归每次拿出栈底，先逆序剩余栈，再把这个栈底元素压回栈顶。
// 时间复杂度：O(N^2)。
// 空间复杂度：O(N)，递归调用栈。

// ReverseStackUsingRecursive 只用递归反转栈。
// 这里用切片模拟栈，尾部是栈顶。
//
// 核心：removeBottom 每次移除并返回栈底元素，上面的元素按原相对顺序压回；
// 递归反转剩余栈后，再把原栈底元素压到栈顶。
//
// 时间复杂度：O(N^2)，每层递归都要向下找一次栈底。
// 空间复杂度：O(N)，递归调用栈。
func ReverseStackUsingRecursive(stack *[]int) {
	if stack == nil || len(*stack) == 0 {
		return
	}
	bottom := removeBottom(stack)
	ReverseStackUsingRecursive(stack)
	*stack = append(*stack, bottom)
}

func removeBottom(stack *[]int) int {
	n := len(*stack)
	result := (*stack)[n-1]
	*stack = (*stack)[:n-1]
	if len(*stack) == 0 {
		return result
	}
	last := removeBottom(stack)
	*stack = append(*stack, result)
	return last
}
