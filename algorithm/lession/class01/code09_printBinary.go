package main

import "fmt"

func main() {
	printBinary(11)
}

// printBinary 打印一个数的二进制
// num & (1 << i) 能判断 num 的第 i 位是 0 还是 1，核心是利用按位与（&:都为1，结果为1；否则是0）运算的特性 + 掩码（1 << i）的构造逻辑
func printBinary(n int) {
	for i := 31; i >= 0; i-- {
		if n&(1<<i) == 0 {
			fmt.Print(0)
		} else {
			fmt.Print(1)
		}

	}
}
