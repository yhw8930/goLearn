package main

import "fmt"

// Hanoi1 使用六个具体方向函数打印 N 层汉诺塔移动过程。
//
// 时间复杂度：O(2^N)，一共移动 2^N-1 次。
// 空间复杂度：O(N)，递归深度为 N。
func Hanoi1(n int) {
	if n > 0 {
		leftToRight(n)
	}
}

func leftToRight(n int) {
	if n == 1 {
		fmt.Println("Move 1 from left to right")
		return
	}
	leftToMid(n - 1)
	fmt.Printf("Move %d from left to right\n", n)
	midToRight(n - 1)
}

func leftToMid(n int) {
	if n == 1 {
		fmt.Println("Move 1 from left to mid")
		return
	}
	leftToRight(n - 1)
	fmt.Printf("Move %d from left to mid\n", n)
	rightToMid(n - 1)
}

func rightToMid(n int) {
	if n == 1 {
		fmt.Println("Move 1 from right to mid")
		return
	}
	rightToLeft(n - 1)
	fmt.Printf("Move %d from right to mid\n", n)
	leftToMid(n - 1)
}

func midToRight(n int) {
	if n == 1 {
		fmt.Println("Move 1 from mid to right")
		return
	}
	midToLeft(n - 1)
	fmt.Printf("Move %d from mid to right\n", n)
	leftToRight(n - 1)
}

func midToLeft(n int) {
	if n == 1 {
		fmt.Println("Move 1 from mid to left")
		return
	}
	midToRight(n - 1)
	fmt.Printf("Move %d from mid to left\n", n)
	rightToLeft(n - 1)
}

func rightToLeft(n int) {
	if n == 1 {
		fmt.Println("Move 1 from right to left")
		return
	}
	rightToMid(n - 1)
	fmt.Printf("Move %d from right to left\n", n)
	midToLeft(n - 1)
}

// Hanoi2 使用一个通用递归函数打印汉诺塔移动过程。
// 含义：把 1..n 从 from 移到 to，other 是辅助柱。
//
// 时间复杂度：O(2^N)。
// 空间复杂度：O(N)。
func Hanoi2(n int) {
	if n > 0 {
		hanoiFunc(n, "left", "right", "mid")
	}
}

func hanoiFunc(n int, from, to, other string) {
	if n == 1 {
		fmt.Printf("Move 1 from %s to %s\n", from, to)
		return
	}
	hanoiFunc(n-1, from, other, to)
	fmt.Printf("Move %d from %s to %s\n", n, from, to)
	hanoiFunc(n-1, other, to, from)
}

type hanoiRecord struct {
	finish1 bool
	base    int
	from    string
	to      string
	other   string
}

// Hanoi3 使用栈模拟 Hanoi2 的递归过程。
//
// 时间复杂度：O(2^N)。
// 空间复杂度：O(N)。
func Hanoi3(n int) {
	if n < 1 {
		return
	}
	stack := []hanoiRecord{{base: n, from: "left", to: "right", other: "mid"}}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.base == 1 {
			fmt.Printf("Move 1 from %s to %s\n", cur.from, cur.to)
			if len(stack) > 0 {
				stack[len(stack)-1].finish1 = true
			}
			continue
		}
		if !cur.finish1 {
			stack = append(stack, cur)
			stack = append(stack, hanoiRecord{base: cur.base - 1, from: cur.from, to: cur.other, other: cur.to})
		} else {
			fmt.Printf("Move %d from %s to %s\n", cur.base, cur.from, cur.to)
			stack = append(stack, hanoiRecord{base: cur.base - 1, from: cur.other, to: cur.to, other: cur.from})
		}
	}
}
