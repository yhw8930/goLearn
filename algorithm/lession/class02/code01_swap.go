package main

import "fmt"

// swap 异或交换数组中两个位置的值
func swapCode10(arr []int, i, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}

func main() {
	// 交换两个独立变量
	a := 16
	b := 603

	fmt.Println(a)
	fmt.Println(b)

	// 异或交换
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a)
	fmt.Println(b)

	// 测试数组中交换同一个位置的元素
	arr := []int{3, 1, 100}

	i := 0
	j := 0

	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]

	fmt.Println(arr[i], " , ", arr[j])

	// 测试 swap 函数交换同一个位置
	fmt.Println(arr[0])
	fmt.Println(arr[2])

	swapCode10(arr, 0, 0)

	fmt.Println(arr[0])
	fmt.Println(arr[2])
}
