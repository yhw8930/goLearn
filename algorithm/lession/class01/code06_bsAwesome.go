package main

import "fmt"

func main() {
	arr1 := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	arr2 := []int{19, 17}
	arr3 := []int{19, 17, 19, 15, 19}
	arr4 := []int{19, 17, 19, 15, 19, 23, 2, 45, 3, 4}
	fmt.Println(oneMinIndex(arr1))
	fmt.Println(oneMinIndex(arr2))
	fmt.Println(oneMinIndex(arr3))
	fmt.Println(oneMinIndex(arr4))
}

func oneMinIndex(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 || arr[0] < arr[1] {
		return 0
	}
	n := len(arr)
	if arr[n-1] < arr[n-2] {
		return n - 1
	}
	left, right := 1, n-2
	for left < right {
		mid := left + (right-left)>>1
		if arr[mid] > arr[mid-1] {
			right = mid - 1
		} else if arr[mid] > arr[mid+1] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
