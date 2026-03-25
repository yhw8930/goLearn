package main

import "fmt"

func GetMax(arr []int) int {
	return process(arr, 0, len(arr)-1)
}

func process(arr []int, L, R int) int {
	if L == R {
		return arr[L]
	}
	mid := L + ((R - L) >> 1)
	leftMax := process(arr, L, mid)
	rightMax := process(arr, mid+1, R)

	if leftMax > rightMax {
		return leftMax
	}
	return rightMax
}

func main() {
	arr := []int{3, 1, 5, 0, 2, 9, 7}
	fmt.Println(GetMax(arr)) // 输出 9
}
