package main

import "fmt"

// 题目：给定 N，计算 1! + 2! + ... + N! 的结果。
// 直接每次重新计算 i! 会重复做大量乘法。
// 可以维护一个变量 cur 表示当前阶乘，处理 i 时由上一次阶乘乘以 i 得到。
// 核心思路：一边累乘得到当前阶乘，一边把当前阶乘累加到答案中。
// 时间复杂度：复用阶乘结果的写法为 O(N)。
// 空间复杂度：O(1)。

func main() {
	fmt.Println(sumFactorial(10))
	fmt.Println(sumFactorial2(10))
	fmt.Println(sumFactorial(10) == sumFactorial2(10))
}

// 计算 1! + 2! + 3! + ... + n!
// sumFactorial 逐项计算阶乘和。
// 如果每项都从头计算阶乘，会产生重复乘法。
func sumFactorial(n int64) int64 {
	res := int64(0)
	for i := int64(1); i <= n; i++ {
		res += factorial(i)
	}
	return res
}

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// sumFactorial2 是更好的累乘写法。
// cur 保存上一个阶乘结果，下一项只需要再乘当前 i。
func sumFactorial2(n int64) int64 {
	if n == 0 {
		return 1
	}
	res := int64(0)
	cur := int64(1)
	for i := int64(1); i <= n; i++ {
		cur *= i
		res += cur
	}
	return res
}
