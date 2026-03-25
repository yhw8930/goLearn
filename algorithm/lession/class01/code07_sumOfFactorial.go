package main

import "fmt"

func main() {
	fmt.Println(sumFactorial(10))
	fmt.Println(sumFactorial2(10))
	fmt.Println(sumFactorial(10) == sumFactorial2(10))
}

// 计算 1! + 2! + 3! + ... + n!
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
