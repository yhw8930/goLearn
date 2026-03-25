package main

import "fmt"

func main() {
	// 异或交换两个数
	a := 5
	b := 7

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a)
	fmt.Println(b)

	// 测试只有一种数出现奇数次
	arr1 := []int{3, 3, 2, 3, 1, 1, 1, 3, 1, 1, 1}
	printOddTimesNum1(arr1)

	// 测试两种数出现奇数次
	arr2 := []int{4, 3, 4, 2, 2, 2, 4, 1, 1, 1, 3, 3, 1, 1, 1, 4, 2, 2}
	printOddTimesNum2(arr2)
}

// arr中，有一种数，出现奇数次
func printOddTimesNum1(arr []int) {
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor ^= arr[i]
	}
	fmt.Println(eor)
}

// 数组里有两种数出现了奇数次，其他数都出现偶数次，找出这两个数
// 时间 O (N)，空间 O (1) 找出这两个数
// a^b: 相同为0，不同为1。 a^a=0, 0^b=b，异或交换不能用在同一个内存地址上
func printOddTimesNum2(arr []int) {
	// 第一步，求出数组中的奇数结果. 因为 a 和 b 是两个不同的数，所以 eor != 0
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor ^= arr[i]
	}

	// 第二步，找出eor中，最右侧为1的位置
	// eor = a^b
	// 假设eor=10100000，那么rightOne=100000
	rightOne := eor & (-eor) // 提取一个数二进制中 最右侧的 1，其他位全部变成 0。

	// 第三步，根据rightOne，将数组分为两组.只异或 “该位是 1” 的数字
	a := 0
	for i := 0; i < len(arr); i++ {
		if (arr[i] & rightOne) != 0 {
			a ^= arr[i]
		}
	}

	b := eor ^ a
	fmt.Println(a)
	fmt.Println(b)
}

// num - 1 会把最右边的 1 变成 0，后面的 0 全变成 1
// & 运算 会刚好把这一整段抹成 0
// 结果就是：每次干掉最右边一个 1
func bit1Count(num int) int {
	count := 0
	for num != 0 {
		num &= num - 1
		count++
	}
	return count
}

// bit1Counts 统计二进制中 1 的个数
// 逻辑：每次提取最右侧的1，计数，然后删掉这个1
func bit1Count2(N int) int {
	count := 0

	//   011011010000
	//   000000010000     1

	//   011011000000

	for N != 0 {
		// 提取 N 最右边的 1（和 Java 的 N & ((~N)+1) 完全等价）
		rightOne := N & -N

		count++
		// 删掉最右边的 1
		N ^= rightOne
		// N -= rightOne 也可以
	}

	return count
}
